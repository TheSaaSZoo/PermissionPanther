package main

import (
	"context"

	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgx"
	"github.com/danthegoodman1/PermissionPanther/crdb"
	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/query"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func CreatePermissionGroup(ns, groupName string, perms []string) (applied bool, err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		return false, err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()

	err = query.New(conn).InsertPermissionGroup(ctx, query.InsertPermissionGroupParams{
		Ns:    ns,
		Name:  groupName,
		Perms: perms,
	})
	if err == nil {
		logger.Logger.WithFields(logrus.Fields{
			"ns":               ns,
			"action":           "create_perm_group",
			"gave_permissions": len(perms) > 0,
		}).Info()
	}
	if pgerr, ok := err.(*pgconn.PgError); ok && pgerr.Code == "23505" {
		// Unique violation, it exists
		return false, nil
	}
	applied = true
	return
}

func RemovePermissionGroup(ns, groupName string, propagate bool) (applied bool, err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		return false, err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()

	err = crdbpgx.ExecuteTx(ctx, conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		txQueries := query.New(tx)

		rows, err := txQueries.DeletePermissionGroup(ctx, query.DeletePermissionGroupParams{
			Ns:   ns,
			Name: groupName,
		})
		if err != nil {
			return err
		}
		if err == nil {
			logger.Logger.WithFields(logrus.Fields{
				"ns":        ns,
				"action":    "remove_perm_group",
				"propagate": propagate,
			}).Info()
		}
		if rows != 0 {
			applied = true
			if propagate {
				// TODO: For all entities in the group, remove their permissions from the group if they exists, log every change for billing
				// Do both of these in transaction
			}
		}
		return nil
	})

	return
}

func AddPermissionToGroup(ns, groupName, perm string, propagate bool) (applied bool, err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		return false, err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()

	err = crdbpgx.ExecuteTx(ctx, conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		txQueries := query.New(tx)
		rows, err := txQueries.AddPermissionToGroup(ctx, query.AddPermissionToGroupParams{
			Ns:         ns,
			Name:       groupName,
			Permission: perm,
		})
		if err != nil {
			return err
		}
		if err == nil && rows != 0 {
			logger.Logger.WithFields(logrus.Fields{
				"ns":     ns,
				"action": "add_perm_to_group",
			}).Info()
		}
		if rows != 0 {
			applied = true
			if propagate {
				// TODO: For all entities in the group, log every change for billing
				// Do both of these in transaction
			}
		}
		return nil
	})

	return
}

func RemovePermissionFromGroup(ns, groupName, perm string, propagate bool) (applied bool, err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		return false, err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()

	err = crdbpgx.ExecuteTx(ctx, conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		txQueries := query.New(tx)
		rows, err := txQueries.RemovePermissionFromGroup(ctx, query.RemovePermissionFromGroupParams{
			Ns:         ns,
			Name:       groupName,
			Permission: perm,
		})
		if err != nil {
			return err
		}
		if err == nil && rows != 0 {
			logger.Logger.WithFields(logrus.Fields{
				"ns":     ns,
				"action": "remove_perm_from_group",
			}).Info()
		}
		if rows != 0 {
			applied = true
			if propagate {
				// TODO: For all entities in the group, log every change for billing
				// Do both of these in transaction
			}
		}
		return nil
	})

	return
}

func ListEntitiesInPermissionGroup(ns, groupName, offset string) (entities []query.PermissionGroupMembership, err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		return []query.PermissionGroupMembership{}, err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()

	entities, err = query.New(conn).ListEntitiesInPermissionGroup(ctx, offset)

	if err != nil {
		logger.Error("Error getting permission group membership")
		return []query.PermissionGroupMembership{}, err
	} else {
		logger.Logger.WithFields(logrus.Fields{
			"ns":     ns,
			"action": "list_perm_group_membership",
			"length": len(entities),
		}).Info()
	}

	return
}

func AddMemberToPermissionGroup(ns, groupName, entity, object string) (applied bool, err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		return false, err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()

	err = crdbpgx.ExecuteTx(ctx, conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		txQueries := query.New(tx)

		err := txQueries.AddMemberToPermissionGroup(ctx, query.AddMemberToPermissionGroupParams{
			GroupName: groupName,
			Entity:    entity,
			Ns:        ns,
			Object:    object,
		})
		if err != nil {
			if pgerr, ok := err.(*pgconn.PgError); ok && pgerr.Code == "23505" {
				// Unique violation, it exists
				applied = false
				return nil
			} else {
				logger.Error("Error adding member to permission group")
				return err
			}
		}

		// TODO: propagate permissions from the group, log every change for billing

		applied = true
		return nil
	})
	return
}

func RemoveMemberFromPermissionGroup(ns, groupName, entity string) (applied bool, err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		return false, err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()

	err = crdbpgx.ExecuteTx(ctx, conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		txQueries := query.New(tx)

		rows, err := txQueries.RemoveMemberFromPermissionGroup(ctx, query.RemoveMemberFromPermissionGroupParams{
			Ns:        ns,
			GroupName: groupName,
			Entity:    entity,
		})
		if err != nil {
			logger.Error("Error removing member from permission group")
			return err
		}

		if rows != 0 {
			// TODO: propagate permissions from the group, log every change for billing
		}

		applied = true
		return nil
	})
	return
}
