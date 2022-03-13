# PermissionPanther
The permissions platform for developers who want to spend less time on permissions.

Permissions for killer apps.

https://permissionpanther.com

## What is Permission Panther?

Permission Panther is a simple and highly scalable Relationship Based Access Control (ReBAC) platform that allows to you spend less time on permissions and more time on your code.

## Why another ReBAC solution? How does it compare to others?

Check out [our awesome blog post for more](https://docs.permissionpanther.com/blog/rbac-vs-rebac), but in short, existing solution are no easier to understand and maintain than custom code. We wanted to build a platform that could:

- Set and check access in one line of code
- Is accessible to every developer
- Required no schemas or entity definitions
- Was the most simple part of any codebase, and as easy to use as any other package
- Scale with the apps it protects

## Environment Variables

### `CACHE_TTL`

The TTL in milliseconds that the instance will cache API key queries. If set to `0`, caching is disabled. Default `0`.

More cache hits result in lower latency and higher concurrency per instance.

### `ADMIN_KEY`

A secure string should be provided to serve as the admin key to the admin HTTP endpoints. This is a temporary measure.

## Admin Endpoints

### POST /key

Create a new key

Query params:
  - `ns`: The namespace for the api key

Returns:
```js
{
  "keyID": string,
  "keySecret": string
}
```

### DELETE /key

Delete an existing key

Query params:
  - `key`: The api `keyID`
