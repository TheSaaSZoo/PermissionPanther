package main

import (
	"log"
	"os"
	"testing"

	"github.com/danthegoodman1/PermissionPanther/crdb"
	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	logger.Logger.SetLevel(logrus.DebugLevel)

	if err := crdb.ConnectToDB(); err != nil {
		logger.Error("Error connecting to crdb on start:")
		logger.Error(err.Error())
		os.Exit(1)
	}
	// TODO: Setup test items

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
