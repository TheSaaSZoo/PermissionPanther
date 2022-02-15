package main

import (
	"context"

	"github.com/danthegoodman1/PermissionPanther/crdb"
	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/query"
	"github.com/jackc/pgconn"
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

	rows, err := query.New(conn).DeletePermissionGroup(ctx, query.DeletePermissionGroupParams{
		Ns:   ns,
		Name: groupName,
	})
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
			// TODO: For all entities in the group, remove their permissions from the group if they exists
			// Do both of these in transaction
		}
	}

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

	rows, err := query.New(conn).AddPermissionToGroup(ctx, query.AddPermissionToGroupParams{
		Ns:         ns,
		Name:       groupName,
		Permission: perm,
	})
	if err == nil && rows != 0 {
		logger.Logger.WithFields(logrus.Fields{
			"ns":     ns,
			"action": "add_perm_to_group",
		}).Info()
	}
	if rows != 0 {
		applied = true
		if propagate {
			// TODO: For all entities in the group
			// Do both of these in transaction
		}
	}

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

	rows, err := query.New(conn).RemovePermissionFromGroup(ctx, query.RemovePermissionFromGroupParams{
		Ns:         ns,
		Name:       groupName,
		Permission: perm,
	})
	if err == nil && rows != 0 {
		logger.Logger.WithFields(logrus.Fields{
			"ns":     ns,
			"action": "remove_perm_from_group",
		}).Info()
	}
	if rows != 0 {
		applied = true
		if propagate {
			// TODO: For all entities in the group
			// Do both of these in transaction
		}
	}

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

func AddMemberToPermissionGroup(ns, groupName, entity string) (applied bool, err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		return false, err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()

	// TODO: Give them all the permissions if they do not exist in transaction
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

	// TODO: Remove all of the permissions from the group if they exist in transaction
}
