package main

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/danthegoodman1/PermissionPanther/crdb"
	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/query"
	"github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	logger.Logger.SetLevel(logrus.DebugLevel)

	if err := crdb.ConnectToDB(); err != nil {
		logger.Error("Error connecting to crdb on start:")
		logger.Error(err.Error())
		os.Exit(1)
	}
	SetupCRDB()

	exitVal := m.Run()

	if exitVal == 0 && os.Getenv("PERSIST_DB") != "true" {
		log.Println("Tests complete, cleaning DB")
	} else if exitVal == 1 {
		log.Println("Test failed, keeping DB contents")
	} else if os.Getenv("PERSIST_DB") == "true" {
		log.Println("PERSIST_DB=true, keeping DB contents")
	} else {
		log.Println("Unknown exit, keeping db contents")
		// TODO: Clear DB
	}
}

func SetupCRDB() {
	conn, err := crdb.PGPool.Acquire(context.Background())
	defer conn.Release()
	if err != nil {
		panic(err)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()

	_, err = conn.Exec(ctx, "TRUNCATE relations")
	if err != nil {
		panic(err)
	}

	queries := query.New(conn)

	err = queries.InsertRelation(ctx, query.InsertRelationParams{
		Object:     "obj1",
		Ns:         "nspc",
		Entity:     "user1",
		Permission: "access",
	})
	HandleError(err)
	// Group
	err = queries.InsertRelation(ctx, query.InsertRelationParams{
		Object:     "obj2",
		Ns:         "nspc",
		Entity:     "~obj1#access",
		Permission: "access",
	})
	HandleError(err)
	// Group
	err = queries.InsertRelation(ctx, query.InsertRelationParams{
		Object:     "obj3",
		Ns:         "nspc",
		Entity:     "~obj2#access",
		Permission: "access",
	})
	HandleError(err)
	// Group
	err = queries.InsertRelation(ctx, query.InsertRelationParams{
		Object:     "obj4",
		Ns:         "nspc",
		Entity:     "~obj3#access",
		Permission: "access",
	})
	HandleError(err)
	// Group
	err = queries.InsertRelation(ctx, query.InsertRelationParams{
		Object:     "obj5",
		Ns:         "nspc",
		Entity:     "~obj4#access",
		Permission: "access",
	})
	HandleError(err)
	// Group
	err = queries.InsertRelation(ctx, query.InsertRelationParams{
		Object:     "obj6",
		Ns:         "nspc",
		Entity:     "~obj5#access",
		Permission: "access",
	})
	HandleError(err)

	// User 2 relations
	// Direct
	err = queries.InsertRelation(ctx, query.InsertRelationParams{
		Object:     "obj1",
		Ns:         "nspc",
		Entity:     "user2",
		Permission: "access",
	})
	HandleError(err)
	// Direct
	err = queries.InsertRelation(ctx, query.InsertRelationParams{
		Object:     "obj2",
		Ns:         "nspc",
		Entity:     "user2",
		Permission: "access",
	})
	HandleError(err)
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
