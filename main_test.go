package main

import (
	"log"
	"os"
	"testing"

	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/scylla"
	"github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	logger.Logger.SetLevel(logrus.DebugLevel)

	scylla.DBConfig()
	scylla.DBConnect()
	scylla.DBTestKeyspaceSetup()
	scylla.DBConnectWithKeyspace()
	scylla.DBTestSetup()

	exitVal := m.Run()

	if exitVal == 0 && os.Getenv("PERSIST_DB") != "true" {
		log.Println("Tests complete, cleaning DB")
	} else if exitVal == 1 {
		log.Println("Test failed, keeping DB contents")
	} else if os.Getenv("PERSIST_DB") == "true" {
		log.Println("PERSIST_DB=true, keeping DB contents")
	} else {
		log.Println("Unknown exit, keeping db contents")
	}
}
