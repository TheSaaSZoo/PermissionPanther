---
sidebar_position: 1
---

# Concepts

Permission Panther keeps permissions extremely simple, but there are a few thing you might want to familiarize yourself with before jumping in:

## ReBAC vs. RBAC

**RBAC**, or Role-Based Access Control is a method of authorizing users by their role.

For example someone might be a 'user', 'admin', or 'owner'.

As you can see, these simple roles are overly permissive and don't allow for substantial control over what users have access to.

**ReBAC**, or Relationship-Based Access Control is a method of authorizing users by what permissions they have on an individual object. This is done with [Relations](#relations) (see below).

## Relations

Relations link an `entity` to an `object` by a `permission`. This is the basic building block of ReBAC, as it allow us to be extremely granular and expressive with our permissions. They are usually in the form of a tuple like:

```
(entity, permission, object)
```

For example, if we wanted to give the user `example_user` permission to `READ` the object `my_awesome_file`, that might look like:

```
("example_user", "READ", "my_awesome_file")
```

A lack of a relation is an implicit denial, so you can be assured that users only have access if we have declared it so.

## Inheritance and Recursion

**Inheritance** allows you to simplify the graph of relations by inheriting permissions from another object.

The best want to explain this is with an example: Think of a Git Repository management product where we want to have organizations that permit all members to be viewers of repositories owned by the organization.

This is accomplished by creating an **Inheritance** relation.

First, we want to ensure that users are members of a given organization with the relation:

```js
await client.SetPermission("user_a", "MEMBER", "org_a")
```

Next, we want to make an an **Inheritance** relation that permits anyone who is a `MEMBER` of the owning organization can `VIEW` a given repository. We can do this by using a special entity that will define this relation as an **Inheritance** relation.

```js
await client.SetPermission(client.Inherit("MEMBER", "org_a"), "VIEW", "repo_a")
```

This relation simply says:

> Whoever is a `MEMBER` of `org_a`, can `VIEW` `repo_a`

Permission Panther knows by this special entity scheme that it needs to look at who is a `MEMBER` on `org_a` to determine if if the user has permission.

Now, we can check if a given user has `VIEW` on a repository:

```js
await client.CheckPermission("user_a", "VIEW", "repo_a")
```

:::caution
**Inheritance** incurs **recursion**, which means that for every nested inheritance relation, Permission Panther checks for whether the user has permission through that inheritance relation, or checks whether there are more nested inheritance relations to look through, until it hits `max_recursion` or finds the relation.
:::

Permission Panther uses a `max_recursion` per API Key that allows you to granularly control which keys are allowed to recurse how many times.

The impact of recursion is a factor of how many nested inheritance relations you have, and the latency of your database queries.

Each recursion of an inheritance check should take (in the common case), 1-3ms. Permission Panther does nested inheritance relation checks concurrently, so ensuring that you have enough query bandwidth to perform highly nested checks will keep your latency low.

## Permission Groups

Sometimes individual permissions are too specific, and you might want to give a user a set of permissions instead.

For example, what if we wanted to give a `MAINTAINER` permission to all repositories that allowed them to `VIEW`, `WRITE`, `CREATE_PR`, `APPROVE_PR`, and more?

Rather than making many inheritance relations for each permission and repository, we can define a [Permission Group](#permission-groups) that includes multiple permissions under one.

First we define the Permission Group:

```js
await client.CreatePermissionGroup("MAINTAINER", ["VIEW", "WRITE", "CREATE_PR", ...])
```

This simply creates a shortcut that says:

> When I give someone `MAINTAINER`, give them `VIEW`, `WRITE`, `CREATE_PR`, etc.

Now, when we create the inheritance relation we can reference the Permission Group as the permission:

```js
await client.SetPermission(client.Inherit("MEMBER", "org_a"), client.PermissionGroup("MAINTAINER"), "repo_a")
```

This one line of code says:

> Whoever is a `MEMBER` of `org_a`, can `VIEW`, `WRITE`, `CREATE_PR`, etc. on `repo_a`

If a user wants to create a pull request on the repository, we can now check the following:

```js
await client.CheckPermission("user_a", "CREATE_PR", "repo_a")
```

We can also check whether the user belongs to this permission group:

```js
await client.CheckPermission("user_a", client.PermissionGroup("MAINTAINER"), "repo_a")
```
