package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/scylla"
)

func main() {
	logger.ConfigureLogger()

	scylla.DBConfig()
	scylla.DBConnectWithKeyspace()

	go StartGRPCServer("8080")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	logger.Warn("Received shutdown signal")
	GRPCServer.GracefulStop()
	logger.Info("Stopped gRPC server")
}
