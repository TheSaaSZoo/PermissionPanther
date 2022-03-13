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

## Features

- **Inheritance** - “Who ever has the `editor` permission of this folder, also has the `editor` permission for all files inside that folder”, or "Who ever is an `editor` can `read`, `write`, etc."
- **Fine-grained scoping and future-proofing** - Since an `object` can be anything, we can reduce permissions down to what ever access level we want, or anything we want, without changing the way our code works.
- **Always check for the same permission** - If you are checking if someone can view something, always check for the `view` permission. No more checking lots of potential roles and conditions that could change over time.
- **Everything RBAC can do, and more** - With Permission Groups, you can define roles that inherit a set of permissions. Now you have all the features of RBAC, with so much more, without having to check multiple roles for a certain action.

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
