package utils

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/danthegoodman1/PermissionPanther/logger"
)

var (
	HTTP_PORT  = os.Getenv("HTTP_PORT")
	REDIS_HOST = os.Getenv("REDIS_HOST")
)

func GetEnvOrDefault(env, defaultVal string) string {
	e := os.Getenv(env)
	if e == "" {
		logger.Debug("Using default value for env var ", env)
		return defaultVal
	} else {
		return e
	}
}

func GetEnvOrFail(env string) string {
	e := os.Getenv(env)
	if e == "" {
		logger.Error(fmt.Sprintf("Failed to find env var '%s'", env))
		os.Exit(1)
		return ""
	} else {
		return e
	}
}

func CheckFlags() {
	flag.StringVar(&HTTP_PORT, "http-port", HTTP_PORT, "Specify the http port to listen on")
	flag.StringVar(&REDIS_HOST, "redis-host", REDIS_HOST, "The redis host to use for service discovery")

	flag.Parse()

	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}
}

func HandleTestError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
