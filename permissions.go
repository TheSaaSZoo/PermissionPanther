package main

import (
	"strings"

	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/scylla"
	"github.com/scylladb/gocqlx/v2/qb"
)

// Finds at what recursion level a permission exists
//
// Returns -1 if permission is not found
func CheckPermissions(ns, obj, permission, entity string, currentRecursion, maxRecursion int) int {
	if currentRecursion > maxRecursion {
		logger.Debug("Aborting nested group checks, exceeded %d recursions!", maxRecursion)
		// Fail fast
		return -1
	}
	logger.Debug("Running permission check, recursion: %d/%d", currentRecursion, maxRecursion)

	// First check for direct access
	directChan := make(chan bool)
	go CheckPermissionDirect(directChan, ns, obj, permission, entity)

	// Then check for groups with this permission
	logger.Debug("Getting groups with ", permission, " on ", obj)
	groupsChan := make(chan []scylla.Edge)
	go GetPermissionGroups(groupsChan, ns, obj, permission)

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
	HandleError(err)

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
	HandleError(err)

	if len(edges) == 0 {
		logger.Debug("Did not find any group lookups")
	} else {
		logger.Debug("Found group lookup!")
	}
	c <- edges
}