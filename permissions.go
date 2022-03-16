package main

import (
	"context"
	"strings"
	"time"

	"github.com/TheSaaSZoo/PermissionPanther/crdb"
	"github.com/TheSaaSZoo/PermissionPanther/errs"
	"github.com/TheSaaSZoo/PermissionPanther/logger"
	"github.com/TheSaaSZoo/PermissionPanther/pb"
	"github.com/TheSaaSZoo/PermissionPanther/query"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

const (
	QueryTimeout = 10 * time.Second
)

var (
	ErrNotFound = errs.Error("not found")
)

// Finds at what recursion level a permission exists
//
// Returns -1 if permission is not found or invalid
func CheckPermissions(ns, object, permission, entity string, currentRecursion, maxRecursion int) (int, error) {
	if currentRecursion > maxRecursion {
		logger.Debug("Aborting nested group checks, exceeded %d recursions!", maxRecursion)
		logger.Logger.WithFields(logrus.Fields{
			"ns":        ns,
			"action":    "exceed_recursion",
			"recursion": maxRecursion,
		}).Info()
		// Fail fast
		return -1, nil
	}
	logger.Debug("Running permission check, recursion: %d/%d", currentRecursion, maxRecursion)
	logger.Logger.WithFields(logrus.Fields{
		"ns":        ns,
		"action":    "check_recursion",
		"recursion": currentRecursion,
	}).Info()

	// First check for direct access
	found, err := CheckPermissionDirect(ns, object, permission, entity)
	if err != nil {
		logger.Error("Error checking permission direct for check permission")
		return -1, err
	}

	// Then check for groups with this permission
	logger.Debug("Getting groups with '%+v' on '%+v'", permission, object)
	groups, err := GetPermissionGroups(ns, object, permission)
	if err != nil {
		logger.Error("Error getting permission groups for check permission")
		return -1, err
	}

	// Might be able to further optimize this with
	// select or with processing inside the goroutine
	// But that would only squeeze maybe a ms or so at 5+ recursions
	// And would require significant rework of this functionality

	// Check direct permission check results
	if found {
		logger.Debug("Found access with recursion %d/%d", currentRecursion, maxRecursion)
		return currentRecursion, nil
	}

	// Check group results
	for _, group := range groups {
		logger.Debug("Got group %s", group.Entity)
		// Get new object and permission to search
		break1 := strings.Split(group.Entity, "#")
		newPermission := break1[1]
		break2 := strings.Split(break1[0], "~")
		newObject := break2[1]
		return CheckPermissions(ns, newObject, newPermission, entity, currentRecursion+1, maxRecursion)
	}
	return -1, nil
}

// Checks whether there is the direct permission mapping from an entity to an object
func CheckPermissionDirect(ns, obj, permission, entity string) (bool, error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		return false, err
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()

	params := query.CheckRelationDirectParams{
		Ns:         ns,
		Object:     obj,
		Permission: permission,
		Entity:     entity,
	}

	_, err = query.New(conn).CheckRelationDirect(ctx, params)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			logger.Logger.WithFields(logrus.Fields{
				"ns":     ns,
				"action": "check_direct",
			}).Info()
			return false, nil
		default:
			logger.Error("Error getting direct relation %+v", params)
			return false, err
		}
	} else {
		logger.Logger.WithFields(logrus.Fields{
			"ns":     ns,
			"action": "check_direct",
		}).Info()
	}

	return true, nil
}

// Returns array of groups that have this permission
func GetPermissionGroups(ns, obj, permission string) ([]query.Relation, error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		return []query.Relation{}, err
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()

	params := query.GetGroupRelationsParams{
		Ns:         ns,
		Object:     obj,
		Permission: permission,
	}

	r, err := query.New(conn).GetGroupRelations(ctx, params)
	if err != nil {
		logger.Error("Error getting group relations %+v", params)
		return []query.Relation{}, err
	} else {
		logger.Logger.WithFields(logrus.Fields{
			"ns":     ns,
			"action": "check_groups",
			"length": len(r),
		}).Info()
	}

	return r, nil
}

func ListEntityPermissions(ns, entity, permission string, offset int32) (relations []*pb.Relation, err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		return []*pb.Relation{}, err
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()

	var r []query.Relation

	if permission == "" {
		r, err = query.New(conn).ListEntityRelations(ctx, query.ListEntityRelationsParams{
			Ns:     ns,
			Entity: entity,
			Offset: offset,
		})
	} else {
		r, err = query.New(conn).ListEntityRelationsWithPermission(ctx, query.ListEntityRelationsWithPermissionParams{
			Ns:         ns,
			Entity:     entity,
			Permission: permission,
			Offset:     offset,
		})
	}

	if err == nil {
		logger.Logger.WithFields(logrus.Fields{
			"ns":             ns,
			"action":         "list_entity",
			"length":         len(r),
			"has_permission": permission != "",
		}).Info()
	} else {
		return
	}

	for _, e := range r {
		relations = append(relations, &pb.Relation{
			Entity:     e.Entity,
			Permission: e.Permission,
			Object:     e.Object,
		})
	}

	return
}

func ListObjectPermissions(ns, object, permission string, offset int32) (relations []*pb.Relation, err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		return []*pb.Relation{}, err
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()

	var r []query.Relation

	if permission == "" {
		r, err = query.New(conn).ListObjectRelations(ctx, query.ListObjectRelationsParams{
			Ns:     ns,
			Object: object,
			Offset: offset,
		})
	} else {
		r, err = query.New(conn).ListObjectRelationsWithPermission(ctx, query.ListObjectRelationsWithPermissionParams{
			Ns:         ns,
			Object:     object,
			Permission: permission,
			Offset:     offset,
		})
	}

	if err == nil {
		logger.Logger.WithFields(logrus.Fields{
			"ns":             ns,
			"action":         "list_object",
			"length":         len(r),
			"has_permission": permission != "",
		}).Info()
	} else {
		return
	}

	for _, e := range r {
		relations = append(relations, &pb.Relation{
			Entity:     e.Entity,
			Permission: e.Permission,
			Object:     e.Object,
		})
	}

	return
}

func UpsertRelation(ns, obj, permission, entity string) (applied bool, err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		return false, err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()
	rows, err := query.New(conn).InsertRelation(ctx, query.InsertRelationParams{
		Ns:         ns,
		Permission: permission,
		Object:     obj,
		Entity:     entity,
	})

	if err != nil {
		logger.Error("Error inserting relation")
	} else {
		logger.Logger.WithFields(logrus.Fields{
			"ns":      ns,
			"action":  "upsert_relation",
			"applied": rows != 0,
		}).Info()
	}

	return rows != 0, nil
}

func DeleteRelation(ns, obj, permission, entity string) (applied bool, err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		return false, err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()
	var rows int64
	rows, err = query.New(conn).DeleteRelation(ctx, query.DeleteRelationParams{
		Ns:         ns,
		Permission: permission,
		Object:     obj,
		Entity:     entity,
	})

	if err == nil {
		logger.Logger.WithFields(logrus.Fields{
			"ns":      ns,
			"action":  "delete_relation",
			"applied": rows != 0,
		}).Info()
	}
	if rows == 0 {
		return false, nil
	}

	return true, nil
}
