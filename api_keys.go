package main

import (
	"context"
	"strings"
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
		MaxCost:     1 << 22, // 4.2MB
		NumCounters: 1e4,     // 10k keys
		BufferItems: 64,
	})
	return err
}

func CheckAPIKey(keyID, keySecret string) (namespace string, err error) {
	keyHash := ""
	found := false
	var val interface{}

	// Check cache
	if utils.CACHE_TTL != 0 {
		// First check the cache
		val, found = APIKeyCache.Get(keyID)
		if found {
			valString, ok := val.(string)
			if !ok {
				logger.Error("Error casting cached value for keyid %s to string", keyID)
			} else {
				keyHash = strings.Split(valString, "#")[0]
				namespace = strings.Split(valString, "#")[1]
			}
		}
	}

	// If we don't have results from the cache
	if !found {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
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
	if utils.CACHE_TTL != 0 && found {
		logger.Debug("Loading cache for keyid %s", keyID)
		// 72 for bcrypt, 27 left for namespace
		added := APIKeyCache.SetWithTTL(keyID, strings.Join([]string{keyHash, namespace}, "#"), 100, time.Millisecond*time.Duration(utils.CACHE_TTL))
		if !added {
			logger.Warn("Did not add keyID %s to cache!", keyID)
		}
	}

	return
}
