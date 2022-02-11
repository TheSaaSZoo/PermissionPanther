package utils

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/danthegoodman1/PermissionPanther/logger"
)

var (
	HTTP_PORT     = os.Getenv("HTTP_PORT")
	ADMIN_KEY     = GetEnvOrFail("ADMIN_KEY")
	CACHE_TTL int = 0
)

func GetEnvOrDefault(env, defaultVal string) string {
	e := os.Getenv(env)
	if e == "" {
		logger.Debug("Using default value of '%s' for env var '%s'", defaultVal, env)
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

	flag.Parse()

	if HTTP_PORT == "" {
		HTTP_PORT = "8080"
	}
	cacheEnvVar := os.Getenv("CACHE_TTL")
	if cacheEnvVar != "" {
		intVar, err := strconv.Atoi(cacheEnvVar)
		if err != nil {
			logger.Error("Error converting CACHE_TTL to int:")
			logger.Error(err.Error())
		} else {
			CACHE_TTL = intVar
		}
	}
}

func HandleTestError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}
