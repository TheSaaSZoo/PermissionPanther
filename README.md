# PermissionPanther
The permissions platform for developers who want to spend less time on permissions.

Permissions for killer apps.

[Docs](https://docs.permissionpather.com?ref=gh)

[Check out our managed offering with generous free usage!](https://permissionpanther.com?ref=gh)

## What is Permission Panther?

Permission Panther is a simple and highly scalable Relationship Based Access Control (ReBAC) platform that allows to you spend less time on permissions and more time on your code.

## Why another ReBAC solution? How does it compare to others?

Check out [our awesome blog post for more](https://docs.permissionpanther.com/blog/rbac-vs-rebac?ref=gh), but in short, existing solution are no easier to understand and maintain than custom code. We wanted to build a platform that could:

- Set and check access in one line of code
- Is accessible to every developer
- Required no schemas or entity definitions
- Was the most simple part of any codebase, and as easy to use as any other package
- Scale with the apps it protects

We wanted to build a solution that prevents this:

```js
// First check if they are invited explicitly
let role = await getUserInvited(user.Org, resource.ID) // database call

if (role && [
      "viewer",
      "writer",
      "editor",
      "editor",
      "admin",
      "viewer"
    ].includes(role)) {
    // They can view
}

if (resource.Org == user.Org) {
  // Check if they are part of the owning organization
  role = await getUserOrgRole(user.ID) // database call
  if (role && [
        "viewer",
        "writer",
        "editor",
        "editor",
        "admin",
        "viewer"
      ].includes(role)) {
    // They can view
  }
}

// They cannot view
```

And enables this:

```js
if (await client.CheckPermission(user.ID, "VIEW", resource.ID).valid) {
  // They can view
} else {
  // They cannot view
}
```

## Features

- **Inheritance** - “Who ever has the `editor` permission of this folder, also has the `editor` permission for all files inside that folder”, or "Who ever is an `editor` can `read`, `write`, etc."
- **Fine-grained scoping and future-proofing** - Since an `object` can be anything, we can reduce permissions down to what ever access level we want, or anything we want, without changing the way our code works.
- **Always check for the same permission** - If you are checking if someone can view something, always check for the `view` permission. No more checking lots of potential roles and conditions that could change over time.
- **Everything RBAC can do, and more** - With Permission Groups, you can define roles that inherit a set of permissions. Now you have all the features of RBAC, with so much more, without having to check multiple roles for a certain action.

See the [Concepts](https://docs.permissionpanther.com/docs/getting-started/concepts?ref=gh) docs for more.

## Install and Host

[See the self-hosting guide](https://docs.permissionpanther.com/docs/getting-started/install-server?ref=gh) to deploy.

You can also check out our [managed offering](https://permissionpanther.com?ref=gh) if you don't want to worry about the database, scaling, or updating.
