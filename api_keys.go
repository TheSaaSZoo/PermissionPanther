package main

import (
	"context"
	"time"

	"github.com/danthegoodman1/PermissionPanther/crdb"
	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/query"
	"github.com/jackc/pgx/v4"
)

func CheckAPIKey(key string) (namespace string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	conn, err := crdb.PGPool.Acquire(ctx)
	if err != nil {
		logger.Error("Error acquiring pgpool connection")
		return "", err
	}
	defer conn.Release()

	ns, err := query.New(conn).SelectAPIKeyNS(ctx, keyID)
	if err != nil {
		if err != pgx.ErrNoRows {
			logger.Error("Error selecting api key")
		}
		return "", err
	}

	return ns, nil
}
