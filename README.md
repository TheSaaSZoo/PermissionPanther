<p align="center">
<img src="https://github.com/TheSaaSZoo/PermissionPanther/raw/main/docs_web/static/img/g1.png" border="0" alt="logo">
</p>

# Permission Panther
The permissions platform for developers who want to spend less time on permissions.
---

[Docs](https://docs.permissionpanther.com?ref=gh)

[Concepts](https://docs.permissionpanther.com/docs/getting-started/concepts?ref=gh)

[Quick Start](https://docs.permissionpanther.com/docs/getting-started/quick-start?ref=gh)

[Problems Permission Panther Solves](https://docs.permissionpanther.com/docs/problems-solved)

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

[See the self-hosting guide](https://docs.permissionpanther.com/docs/getting-started/install-server?ref=gh) to deploy. [Binaries](https://github.com/TheSaaSZoo/PermissionPanther/releases) and [container images](https://github.com/TheSaaSZoo/PermissionPanther/pkgs/container/permissionpanther) are readily available for both amd64 (x86) and arm64 platforms.

You can also check out our [managed offering](https://permissionpanther.com?ref=gh) if you don't want to worry about the database, scaling, or updating.

## Problems Permission Panther Solves

### Multi-Tenancy Management

Whether you are building managed kubernetes observability, a revolutionary cloud platform, or analytics tools, you need to have you users be able to share access to your product.

Permission Panther allows project owners to invite other users to their projects, without complex permission management. You can give them varying permissions to allow/prevent them using certain features, and what they can see.

Allow your tenants to share without interferring with other tenants, ensure that they can only access the actions and data they are explicily permitted to, and nothing else.

### Content Sharing

Whether it is social media, git repos, or you are building the next google drive, you need to control what your users can do.

Permission Panther can provide you with the access control to determine who can view posts, comment on files, or edit all documents in a folder. What ever actions you may need, Permission Panther ensures users have them.

### Role Based Access Control (RBAC)

ReBAC can do everything that RBAC can do, with far less code, and far more functionality.

### Physical Access Control

Just because it’s software, doesn’t mean it can’t work in the real world!

Permission Panther can be your go-to solution for managing access to physical infrastructure. Whether it be an office building, a single door, or giving support technicians temporary access to on-site locations with Relation TTLs.

Permission Panther can control access to physical objects just as well as it can control access to software!

### Support Agent Management

For some services, support agents shouldn’t have the ability to access data or infrastructure without explicit approval of customers.

Permission Panther can allow for temporary access to customer data, along with limited functionality, so support agents only have the access they need. You can even have various levels of support agent escallation with Permission Groups, so that escallation to increased access and duration only occurs as situations require.

### Premium Feature Access

Instead of checking what subscription a user has in every request, give users with higher subscription tiers specific permissions to access certain features. Protect your premium features with a single line of code. With Permission Groups you can define what access a subscription tier has.

### Feature Flagging

When accessing a new feature, check if they have permission to a new or old version.
