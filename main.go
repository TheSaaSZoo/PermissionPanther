package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/TheSaaSZoo/PermissionPanther/crdb"
	"github.com/TheSaaSZoo/PermissionPanther/logger"
	"github.com/TheSaaSZoo/PermissionPanther/utils"
	"github.com/sirupsen/logrus"
	"github.com/soheilhy/cmux"
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

	utils.CheckFlags()
	if utils.CACHE_TTL != 0 {
		logger.Debug("CACHE_TTL found, setting up API Key cache")
		err := InitCache()
		if err != nil {
			logger.Error("Error initializing cache:")
			logger.Error(err.Error())
		}
	}

	logger.Info("Starting cmux listener on port %s", utils.GetEnvOrDefault("PORT", "8080"))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", utils.GetEnvOrDefault("PORT", "8080")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	m := cmux.New(lis)
	grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP2()) // m.Match(cmux.HTTP1Fast())
	go StartGRPCServer(grpcL)
	go StartHTTPServer(httpL)

	m.Serve()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	logger.Warn("Received shutdown signal")
	GRPCServer.GracefulStop()
	logger.Info("Stopped gRPC server")
	Server.Echo.Close()
	logger.Info("Stopped HTTP server")
}
