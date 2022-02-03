package scylla

import (
	"log"
	"os"
	"time"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"
)

var (
	// DB Connection
	CQLSession *gocqlx.Session

	// ---------------------------------------------------------------------------
	// DB Models
	// ---------------------------------------------------------------------------
	EdgeMetadata = &table.Metadata{
		Name:    "edge",
		Columns: []string{"obj", "ns", "entity", "permission"},
		PartKey: []string{"ns", "obj"},
		SortKey: []string{"permission", "entity"},
	}

	EntityMetadata = &table.Metadata{
		Name:    "entity_permission",
		Columns: []string{"obj", "ns", "entity", "permission"},
		PartKey: []string{"ns", "entity"},
		SortKey: []string{"permission", "obj"},
	}

	// ---------------------------------------------------------------------------
	// DB Tables
	// ---------------------------------------------------------------------------
	EdgeTable   *table.Table
	EntityIndex *table.Table
)

type Edge struct {
	Obj        string
	Ns         string
	Entity     string // node id or another relation (ns:obj#access)
	Permission string
}

func DBConnect() {
	// Create gocql cluster.
	cluster := gocql.NewCluster("localhost:9042")

	// Increase timeout if testing
	if os.Getenv("TEST_MODE") == "true" {
		cluster.Timeout = 1 * time.Second
	}

	// Wrap session on creation, gocqlx session embeds gocql.Session pointer.
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal(err)
	}
	CQLSession = &session
}

func DBConnectWithKeyspace() {
	// Create gocql cluster.
	cluster := gocql.NewCluster("localhost:9042")
	cluster.Keyspace = "access_kspc"

	// Increase timeout if testing
	if os.Getenv("TEST_MODE") == "true" {
		cluster.Timeout = 1 * time.Second
	}

	// Wrap session on creation, gocqlx session embeds gocql.Session pointer.
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal(err)
	}
	CQLSession = &session
}

func DBConfig() {
	EdgeTable = table.New(*EdgeMetadata)
	EntityIndex = table.New(*EntityMetadata)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func DBTestKeyspaceSetup() {
	// Drop NS
	err := CQLSession.ExecStmt("DROP KEYSPACE IF EXISTS access_kspc")
	HandleError(err)

	// Create NS
	err = CQLSession.ExecStmt("CREATE KEYSPACE IF NOT EXISTS access_kspc WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}")
	HandleError(err)
}

func DBTestSetup() {
	// Create Tables
	HandleError(CQLSession.ExecStmt(`
	CREATE TABLE IF NOT EXISTS edge (
		obj TEXT,
		entity TEXT,
		permission TEXT,
		ns TEXT,
		PRIMARY KEY ((ns, obj), permission, entity)
	)
	`))

	// Walk backward by entity
	// Could also call reverse_edge
	HandleError(CQLSession.ExecStmt(`
	CREATE MATERIALIZED VIEW IF NOT EXISTS entity_permission AS
		SELECT * FROM edge
		WHERE ns IS NOT NULL
			AND entity IS NOT NULL
			AND permission IS NOT NULL
			AND obj IS NOT NULL
		PRIMARY KEY((ns, entity), permission, obj)
	`))

	// User 1 relations
	// Direct
	q := CQLSession.Query(EdgeTable.Insert()).BindStruct(Edge{
		Obj:        "obj1",
		Ns:         "nspc",
		Entity:     "user1",
		Permission: "access",
	})
	HandleError(q.ExecRelease())
	// Group
	q = CQLSession.Query(EdgeTable.Insert()).BindStruct(Edge{
		Obj:        "obj2",
		Ns:         "nspc",
		Entity:     "~obj1#access",
		Permission: "access",
	})
	HandleError(q.ExecRelease())
	// Group
	q = CQLSession.Query(EdgeTable.Insert()).BindStruct(Edge{
		Obj:        "obj3",
		Ns:         "nspc",
		Entity:     "~obj2#access",
		Permission: "access",
	})
	HandleError(q.ExecRelease())
	// Group
	q = CQLSession.Query(EdgeTable.Insert()).BindStruct(Edge{
		Obj:        "obj4",
		Ns:         "nspc",
		Entity:     "~obj3#access",
		Permission: "access",
	})
	HandleError(q.ExecRelease())
	// Group
	q = CQLSession.Query(EdgeTable.Insert()).BindStruct(Edge{
		Obj:        "obj5",
		Ns:         "nspc",
		Entity:     "~obj4#access",
		Permission: "access",
	})
	HandleError(q.ExecRelease())
	// Group
	q = CQLSession.Query(EdgeTable.Insert()).BindStruct(Edge{
		Obj:        "obj6",
		Ns:         "nspc",
		Entity:     "~obj5#access",
		Permission: "access",
	})
	HandleError(q.ExecRelease())

	// User 2 relations
	// Direct
	q = CQLSession.Query(EdgeTable.Insert()).BindStruct(Edge{
		Obj:        "obj1",
		Ns:         "nspc",
		Entity:     "user2",
		Permission: "access",
	})
	HandleError(q.ExecRelease())
	// Direct
	q = CQLSession.Query(EdgeTable.Insert()).BindStruct(Edge{
		Obj:        "obj2",
		Ns:         "nspc",
		Entity:     "user2",
		Permission: "access",
	})
	HandleError(q.ExecRelease())
}
