package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/danthegoodman1/PermissionPanther/crdb"
	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/utils"
	"github.com/sirupsen/logrus"
)

var (
	MAX_RECURSIONS = 5
)

func main() {
	if os.Getenv("TEST_MODE") == "true" {
		logger.Logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.ConfigureLogger()
	}

	if err := crdb.ConnectToDB(); err != nil {
		logger.Error("Error connecting to crdb on start:")
		logger.Error(err.Error())
		os.Exit(1)
	}

	go StartGRPCServer(utils.GetEnvOrDefault("PORT", "8080"))
	go StartHTTPServer(utils.GetEnvOrDefault("HTTP_PORT", "9090"))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	logger.Warn("Received shutdown signal")
	GRPCServer.GracefulStop()
	logger.Info("Stopped gRPC server")
}
