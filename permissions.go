package main

import (
	"context"
	"strings"
	"time"

	"github.com/danthegoodman1/PermissionPanther/crdb"
	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/pb"
	"github.com/danthegoodman1/PermissionPanther/query"
	"github.com/jackc/pgx/v4"
)

const (
	QueryTimeout = 10 * time.Second
)

// Finds at what recursion level a permission exists
//
// Returns -1 if permission is not found or invalid
func CheckPermissions(ns, object, permission, entity string, currentRecursion, maxRecursion int) int {
	if currentRecursion > maxRecursion {
		logger.Debug("Aborting nested group checks, exceeded %d recursions!", maxRecursion)
		// Fail fast
		return -1
	}
	logger.Debug("Running permission check, recursion: %d/%d", currentRecursion, maxRecursion)

	// First check for direct access
	directChan := make(chan bool)
	go CheckPermissionDirect(directChan, ns, object, permission, entity)

	// Then check for groups with this permission
	logger.Debug("Getting groups with '%+v' on '%+v'", permission, object)
	groupsChan := make(chan []query.Relation)
	go GetPermissionGroups(groupsChan, ns, object, permission)

	// Might be able to further optimize this with
	// select or with processing inside the goroutine
	// But that would only squeeze maybe a ms or so at 5+ recursions
	// And would require significant rework of this functionality
	directPerms := <-directChan
	groups := <-groupsChan

	// Check direct permission check results
	if directPerms {
		logger.Debug("Found access with recursion %d/%d", currentRecursion, maxRecursion)
		return currentRecursion
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
	return -1
}

// Checks whether there is the direct permission mapping from an entity to an object
func CheckPermissionDirect(c chan bool, ns, obj, permission, entity string) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		logger.Error(err.Error())
		c <- false
		return
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
			c <- false
		default:
			logger.Error("Error getting direct relation %+v", params)
			logger.Error(err.Error())
			c <- false
		}
		return
	}

	c <- true
}

// Returns array of groups that have this permission
func GetPermissionGroups(c chan []query.Relation, ns, obj, permission string) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		logger.Error(err.Error())
		c <- []query.Relation{}
		return
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
		logger.Error(err.Error())
		c <- []query.Relation{}
		return
	}
	c <- r
}

func ListEntityPermissions(ns, entity string, permission *string) (relations []*pb.Relation, err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		logger.Error(err.Error())
		return []*pb.Relation{}, err
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()

	var r []query.Relation

	if permission == nil {
		r, err = query.New(conn).ListEntityRelations(ctx, query.ListEntityRelationsParams{
			Ns:     ns,
			Entity: entity,
		})
	} else {
		r, err = query.New(conn).ListEntityRelationsWithPermission(ctx, query.ListEntityRelationsWithPermissionParams{
			Ns:         ns,
			Entity:     entity,
			Permission: *permission,
		})
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

func ListObjectPermissions(ns, object string, permission *string) (relations []*pb.Relation, err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		logger.Error(err.Error())
		return []*pb.Relation{}, err
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()

	var r []query.Relation

	if permission == nil {
		r, err = query.New(conn).ListObjectRelations(ctx, query.ListObjectRelationsParams{
			Ns:     ns,
			Object: object,
		})
	} else {
		r, err = query.New(conn).ListObjectRelationsWithPermission(ctx, query.ListObjectRelationsWithPermissionParams{
			Ns:         ns,
			Object:     object,
			Permission: *permission,
		})
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

func UpsertRelation(ns, obj, permission, entity string) (err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		logger.Error(err.Error())
		return err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()
	err = query.New(conn).InsertRelation(ctx, query.InsertRelationParams{
		Ns:         ns,
		Permission: permission,
		Object:     obj,
		Entity:     entity,
	})
	return
}

func DeleteRelation(ns, obj, permission, entity string) (err error) {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		logger.Error("Error acquiring pool connection")
		logger.Error(err.Error())
		return err
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), QueryTimeout)
	defer cancelFunc()
	err = query.New(conn).DeleteRelation(ctx, query.DeleteRelationParams{
		Ns:         ns,
		Permission: permission,
		Object:     obj,
		Entity:     entity,
	})
	return
}
