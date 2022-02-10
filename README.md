# PermissionPanther
Relationship Based Access Control (ReBAC) for killer software

## Environment Variables

### `CACHE_TTL`

The TTL in milliseconds that the instance will cache API key queries. If set to `0`, caching is disabled. Default `0`.

More cache hits result in lower latency and higher concurrency per instance.

<!-- ### `ADMIN_KEY_HASH`

The SHA256 hash lower case hex string of a string used as the admin key for admin functionality (creating/destroying API keys). This must be defined.

The code used to compare the hash is:

```go
import (
  "crypto/sha256"
  "encoding/hex"
)

var (
  ADMIN_KEY_HASH := "2efa.."
)

func main() {
  keyBytes := []byte(givenKeyString)
  hashBytes := sha256.Sum256(keyBytes)
  hashString := hex.EncodeToString(hashBytes[:])
  validKey := ADMIN_KEY_HASH == hashString
}
```

To create a hash string from a intended key, you can use the following code:

```go
import (
  "crypto/sha256"
  "encoding/hex"
  "fmt"
)

var (
  MY_KEY_STRING := "thisisasupersecretkey"
)

func main() {
  keyBytes := []byte(MY_KEY_STRING)
  hashBytes := sha256.Sum256(keyBytes)
  hashString := hex.EncodeToString(hashBytes[:])
  fmt.Println(hashString)
}
``` -->
