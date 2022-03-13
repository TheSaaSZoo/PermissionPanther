package crdb

import (
	"context"
	"time"

	"github.com/TheSaaSZoo/PermissionPanther/logger"
	"github.com/TheSaaSZoo/PermissionPanther/utils"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	PGPool *pgxpool.Pool
)

func ConnectToDB() error {
	logger.Debug("Connecting to CRDB...")
	var err error
	config, err := pgxpool.ParseConfig(utils.GetEnvOrFail("CRDB_DSN"))
	if err != nil {
		return err
	}

	config.MaxConns = 10
	config.MinConns = 1
	config.HealthCheckPeriod = time.Second * 5
	config.MaxConnLifetime = time.Minute * 30
	config.MaxConnIdleTime = time.Minute * 30

	PGPool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return err
	}
	logger.Debug("Connected to CRDB")
	return nil
}
