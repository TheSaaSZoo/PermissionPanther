package main

import (
	"context"
	"net"
	"net/http"
	"strconv"
	"time"

	_ "net/http/pprof"

	"github.com/TheSaaSZoo/PermissionPanther/crdb"
	"github.com/TheSaaSZoo/PermissionPanther/logger"
	"github.com/TheSaaSZoo/PermissionPanther/query"
	"github.com/TheSaaSZoo/PermissionPanther/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	nanoid "github.com/matoous/go-nanoid/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/net/http2"
)

type HTTPServer struct {
	Echo *echo.Echo
}

var (
	Server *HTTPServer
)

func StartHTTPServer(lis net.Listener) {
	echoInstance := echo.New()
	Server = &HTTPServer{
		Echo: echoInstance,
	}
	Server.Echo.HideBanner = true
	// Server.Echo.Use(middleware.Logger())
	config := middleware.LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out},"proto":"${protocol}"}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}
	Server.Echo.Use(middleware.LoggerWithConfig(config))

	// Setup admin routes
	Server.Echo.POST("/key", CreateAPIKey, ValidateAdminKey)
	Server.Echo.DELETE("/key", DeleteAPIKey, ValidateAdminKey)

	// Count requests
	Server.Echo.GET("/metrics", wrapPromHandler)
	SetupMetrics()

	logger.Info("Starting Permission Panther HTTP API")
	Server.Echo.Listener = lis
	server := &http2.Server{}
	Server.Echo.StartH2CServer("", server) // HTTP/2
	// Server.Echo.Start("") // HTTP/1.1
}

func ValidateRequest(c echo.Context, s interface{}) error {
	if err := c.Bind(s); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(s); err != nil {
		return err
	}
	return nil
}

func wrapPromHandler(c echo.Context) error {
	h := promhttp.Handler()
	h.ServeHTTP(c.Response(), c.Request())
	return nil
}

func ValidateAdminKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		adminKeyHeader := c.Request().Header.Get("ak")
		if adminKeyHeader != utils.ADMIN_KEY {
			return c.String(http.StatusForbidden, "Invalid admin key")
		} else {
			return next(c)
		}
	}
}

func CreateAPIKey(c echo.Context) error {
	ns := c.QueryParam("ns")
	mr := c.QueryParam("mr")
	if ns == "" {
		return c.String(400, "Missing `ns` query param")
	}
	if mr == "" {
		return c.String(400, "Missing `mr` query param")
	}

	max_recursions, err := strconv.Atoi(mr)
	if err != nil {
		logger.Warn("Failed to cast `mr` query param to int")
		logger.Warn(err.Error())
		return c.String(http.StatusInternalServerError, err.Error())
	}

	keyID := nanoid.Must()
	keySecret := nanoid.Must()

	keySecretHash, err := HashPassword(keySecret)
	if err != nil {
		logger.Error("Error generating argon2 hash:")
		logger.Error(err.Error())
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Store the key hash
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	conn, err := crdb.PGPool.Acquire(ctx)
	if err != nil {
		logger.Error("Error acquiring pgpool connection")
		logger.Error(err.Error())
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer conn.Release()

	err = query.New(conn).InsertAPIKey(ctx, query.InsertAPIKeyParams{
		ID:            keyID,
		SecretHash:    string(keySecretHash),
		Ns:            ns,
		MaxRecursions: int64(max_recursions),
	})
	if err != nil {
		logger.Error("Error inserting api key")
		logger.Error(err.Error())
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(200, map[string]interface{}{
		"keyID":     keyID,
		"keySecret": keySecret,
	})
}

func DeleteAPIKey(c echo.Context) error {
	key := c.QueryParam("key")
	if key == "" {
		return c.String(400, "Missing `key` query param")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	conn, err := crdb.PGPool.Acquire(ctx)
	if err != nil {
		logger.Error("Error acquiring pgpool connection")
		logger.Error(err.Error())
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer conn.Release()

	rows, err := query.New(conn).DeleteAPIKey(ctx, key)
	if err != nil {
		logger.Error("Error deleting api key")
		logger.Error(err.Error())
		return c.String(http.StatusInternalServerError, err.Error())
	} else if rows == 0 {
		logger.Debug("Key %s not found", key)
		return c.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
	return c.NoContent(http.StatusNoContent)
}
