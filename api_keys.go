package main

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/danthegoodman1/PermissionPanther/crdb"
	"github.com/danthegoodman1/PermissionPanther/errs"
	"github.com/danthegoodman1/PermissionPanther/logger"
	"github.com/danthegoodman1/PermissionPanther/query"
	"github.com/danthegoodman1/PermissionPanther/utils"
	"github.com/dgraph-io/ristretto"
	"github.com/jackc/pgx/v4"
	"github.com/minio/argon2"
)

var (
	APIKeyCache *ristretto.Cache
	AP          = &ArgonParams{
		memory:      64 * 64,
		iterations:  1,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	ErrInvalidHash    = errs.Error("invalid hash")
	ErrInvalidVersion = errs.Error("invalid version")
)

type ArgonParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

type CachedKey struct {
	SecretHash    string
	Namespace     string
	MaxRecursions int
}

func InitCache() error {
	var err error
	APIKeyCache, err = ristretto.NewCache(&ristretto.Config{
		MaxCost:     1 << 22, // 4.2MB
		NumCounters: 1e4,     // 10k keys
		BufferItems: 64,
	})
	return err
}

func CheckAPIKey(keyID, keySecret string) (apiKey *CachedKey, err error) {
	found := false
	var val interface{}

	// Check cache
	if utils.CACHE_TTL != 0 {
		// First check the cache
		val, found = APIKeyCache.Get(keyID)
		if found {
			cachedKey, ok := val.(CachedKey)
			if !ok {
				logger.Error("Error casting cached value for keyid %s to CachedKey", keyID)
			} else {
				apiKey = &cachedKey
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
			return nil, err
		}
		defer conn.Release()

		key, err := query.New(conn).SelectAPIKey(ctx, keyID)
		if err != nil {
			if err != pgx.ErrNoRows {
				logger.Error("Error selecting api key")
			}
			return nil, err
		}
		apiKey = &CachedKey{
			SecretHash:    key.SecretHash,
			Namespace:     key.Ns,
			MaxRecursions: int(key.MaxRecursions),
		}
	}

	// Validate secret against argon
	valid, err := ComparePasswordHash(keySecret, apiKey.SecretHash)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, ErrInvalidHash
	}

	// Load cache
	if utils.CACHE_TTL != 0 && !found {
		logger.Debug("Loading cache for keyid %s", keyID)
		added := APIKeyCache.SetWithTTL(keyID, *apiKey, 100, time.Millisecond*time.Duration(utils.CACHE_TTL))
		if !added {
			logger.Warn("Did not add keyID %s to cache!", keyID)
		}
	}

	return
}

func HashPassword(password string) (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), b, AP.iterations, AP.memory, AP.parallelism, AP.keyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(b)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, AP.memory, AP.iterations, AP.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

// Extracts the salt and hash from the string
func DecodePasswordHash(encodedHash string) (salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, ErrInvalidVersion
	}

	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &AP.memory, &AP.iterations, &AP.parallelism)
	if err != nil {
		return nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, err
	}
	AP.saltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, err
	}
	AP.keyLength = uint32(len(hash))

	return salt, hash, nil
}

func ComparePasswordHash(password, encodedHash string) (match bool, err error) {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	salt, hash, err := DecodePasswordHash(encodedHash)
	if err != nil {
		return false, err
	}

	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey([]byte(password), salt, AP.iterations, AP.memory, AP.parallelism, AP.keyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}
