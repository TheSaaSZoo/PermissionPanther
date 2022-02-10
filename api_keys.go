package main

import (
	"context"
	"time"

	"github.com/danthegoodman1/PermissionPanther/crdb"
	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/query"
	"github.com/danthegoodman1/PermissionPanther/utils"
	"github.com/dgraph-io/ristretto"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

var (
	APIKeyCache *ristretto.Cache
)

func InitCache() error {
	var err error
	APIKeyCache, err = ristretto.NewCache(&ristretto.Config{
		MaxCost: 1 << 22, // 4.2MB
	})
	return err
}

func CheckAPIKey(keyID, keySecret string) (namespace string, err error) {
	// Check cache
	keyHash := ""
	if utils.CACHE_TTL != 0 {
		// First check the cache
		val, found := APIKeyCache.Get(keyID)
		if found {
			valString, ok := val.(string)
			if !ok {
				logger.Error("Error casting cached value for keyid %s to string", keyID)
			} else {
				keyHash = valString
			}
		}
	}

	// If we don't have results from the cache
	if keyHash == "" {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		conn, err := crdb.PGPool.Acquire(ctx)
		if err != nil {
			logger.Error("Error acquiring pgpool connection")
			return "", err
		}
		defer conn.Release()

		key, err := query.New(conn).SelectAPIKey(ctx, keyID)
		if err != nil {
			if err != pgx.ErrNoRows {
				logger.Error("Error selecting api key")
			}
			return "", err
		}
		keyHash = key.SecretHash
		namespace = key.Ns
	}

	// Validate secret against bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(keyHash), []byte(keySecret))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		namespace = ""
		return
	}

	// Load cache
	if utils.CACHE_TTL != 0 {
		logger.Debug("Loading cache for keyid %s", keyID)
		added := APIKeyCache.SetWithTTL(keyID, keyHash, 72, time.Millisecond*time.Duration(utils.CACHE_TTL))
		if !added {
			logger.Warn("Did not add keyID %s to cache!", keyID)
		}
	}

	return
}
