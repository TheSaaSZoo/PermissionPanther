package main

import (
	"strings"

	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/pb"
	"github.com/danthegoodman1/PermissionPanther/scylla"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2/qb"
)

// Finds at what recursion level a permission exists
//
// Returns -1 if permission is not found
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
	logger.Debug("Getting groups with ", permission, " on ", object)
	groupsChan := make(chan []scylla.Edge)
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
	logger.Debug("Running permission direct check")
	var edges []scylla.Edge

	query, names := qb.Select(scylla.EdgeMetadata.Name).
		Columns("*").
		Where(qb.Eq("entity")).
		Where(qb.Eq("obj")).
		Where(qb.Eq("ns")).
		Where(qb.Eq("permission")).ToCql()

	q := scylla.CQLSession.Query(query, names).BindStruct(scylla.Edge{
		Obj:        obj,
		Ns:         ns,
		Entity:     entity,
		Permission: permission,
	})
	logger.Debug("Direct Query: %v", q.Query)
	err := q.SelectRelease(&edges)
	if err != nil {
		logger.Error("Error checking direct permissions")
		logger.Error(err.Error())
		c <- false
	}

	c <- len(edges) > 0
}

// Returns array of groups that have this permission
func GetPermissionGroups(c chan []scylla.Edge, ns, obj, permission string) {
	var edges []scylla.Edge

	query, names := qb.Select(scylla.EdgeMetadata.Name).
		Columns("*").
		Where(qb.Gt("entity")).
		Where(qb.Eq("obj")).
		Where(qb.Eq("ns")).
		Where(qb.Eq("permission")).ToCql()

	q := scylla.CQLSession.Query(query, names).BindStruct(scylla.Edge{
		Obj:        obj,
		Ns:         ns,
		Entity:     "~",
		Permission: permission,
	})
	logger.Debug("Group Query: %v", q.Query)
	err := q.SelectRelease(&edges)
	if err != nil {
		logger.Error("Error checking group permissions")
		logger.Error(err.Error())
		c <- []scylla.Edge{}
	}

	if len(edges) == 0 {
		logger.Debug("Did not find any group lookups")
	} else {
		logger.Debug("Found group lookup!")
	}
	c <- edges
}

func ListEntityPermissions(ns, entity string, permission *string) (relations []*pb.Relation, err error) {
	var edges []scylla.Edge

	queryBuilder := qb.Select(scylla.EntityMetadata.Name).
		Columns("*").
		Where(qb.Eq("entity")).
		Where(qb.Eq("ns"))

	edge := scylla.Edge{
		Ns:     ns,
		Entity: entity,
	}

	if permission != nil {
		queryBuilder = queryBuilder.Where(qb.Eq("permission"))
		edge.Permission = *permission
	}

	query, names := queryBuilder.ToCql()

	q := scylla.CQLSession.Query(query, names).BindStruct(edge)
	logger.Debug("Direct Query: %v", q.Query)
	err = q.SelectRelease(&edges)
	if err != nil {
		logger.Error("Error listing entity permissions")
		return
	}

	for _, e := range edges {
		relations = append(relations, &pb.Relation{
			Entity:     e.Entity,
			Permission: e.Permission,
			Object:     e.Obj,
		})
	}
	return
}

func ListObjectPermissions(ns, object string, permission *string) (relations []pb.Relation, err error) {
	var edges []scylla.Edge

	queryBuilder := qb.Select(scylla.EdgeMetadata.Name).
		Columns("*").
		Where(qb.Eq("obj")).
		Where(qb.Eq("ns"))

	edge := scylla.Edge{
		Ns:  ns,
		Obj: object,
	}

	if permission != nil {
		queryBuilder = queryBuilder.Where(qb.Eq("permission"))
		edge.Permission = *permission
	}

	query, names := queryBuilder.ToCql()

	q := scylla.CQLSession.Query(query, names).BindStruct(edge)
	logger.Debug("Direct Query: %v", q.Query)
	err = q.SelectRelease(&edges)
	if err != nil {
		logger.Error("Error listing object permissions")
		return
	}

	for _, e := range edges {
		relations = append(relations, pb.Relation{
			Entity:     e.Entity,
			Permission: e.Permission,
			Object:     e.Obj,
		})
	}
	return
}

func UpsertRelation(ns, obj, permission, entity string) (err error) {
	q := scylla.CQLSession.Query(scylla.EdgeTable.Insert()).BindStruct(scylla.Edge{
		Obj:        obj,
		Ns:         ns,
		Entity:     entity,
		Permission: permission,
	})
	// Enable consistent reads
	q = q.Consistency(gocql.All)
	err = q.ExecRelease()
	return
}

func DeleteRelation(ns, obj, permission, entity string) (err error) {
	q := scylla.CQLSession.Query(scylla.EdgeTable.Delete()).BindStruct(scylla.Edge{
		Obj:        obj,
		Ns:         ns,
		Entity:     entity,
		Permission: permission,
	})
	// Enable consistent reads
	q = q.Consistency(gocql.All)
	err = q.ExecRelease()
	return
}
