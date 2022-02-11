# PermissionPanther
Relationship Based Access Control (ReBAC) for killer software

## Environment Variables

### `CACHE_TTL`

The TTL in milliseconds that the instance will cache API key queries. If set to `0`, caching is disabled. Default `0`.

More cache hits result in lower latency and higher concurrency per instance.

### `ADMIN_KEY`

A secure string should be provided to serve as the admin key to the admin HTTP endpoints. This is a temporary measure.
