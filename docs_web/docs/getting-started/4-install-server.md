---
sidebar_position: 4
---

# Install the Server

:::tip
If you don't want to manage the server and database yourself, [we are working on an amazing managed offering with a generous free tier](https://permissionpanther.com?ref=docs).
:::
If you want to setup Permission Panther to run on your own infrastructure, the steps are simple.

## Setup CockroachDB

You need to setup a CockroachDB instance, whether it be their Serverless offering or a self-managed cluster. [Refer to their guide on how to setup a self-hosted CRDB instance](https://www.cockroachlabs.com/docs/stable/deploy-cockroachdb-on-premises.html).

:::note Quick Tip
For testing you can run a container locally with the command

```
docker run -d -p 26257:26257 \
  cockroachdb/cockroach:latest \
  start-single-node --insecure
```

And use the DSN

```
postgresql://root@localhost:26257/defaultdb?sslmode=disable
```
:::

## Apply Database Migration

### Clone the Git Repo

```
git clone --depth 1 https://github.com/TheSaaSZoo/PermissionPanther
```

### Setup [`sql-migrate`](https://github.com/rubenv/sql-migrate)
Install [`sql-migrate`](https://github.com/rubenv/sql-migrate) and set the `CRDB_DSN` environment variable as your database DSN.

### Apply Migrations
From the root directory of the cloned repo, run:

```
sql-migrate up
```

## Docker

Docker is the quickest way to get Permission Panther running:

```
docker run --name permissionpanther \
  --env CRDB_DSN "postgres://..." \
  --env ADMIN_KEY="example_key" \
  -it --rm
  permissionpanther
```

## Docker Compose

Docker Compose works great for simple long-term deployments:

```yml

```

## Kubernetes

Permission Panther runs stateless, so you can create a Kubernetes deployment that can easily scale horizontally to as many pods as needed.

To get started, pull the example repo that contains all of the manifests:

```
git clone blah...

cd blah...
```

Then from these, you will need to create secrets for your `CRDB_DSN` and `ADMIN_KEY`

```
echo -n "postgresql://..." | kubectl create secret \
  generic crdb-dsn --dry-run=client \
  --from-file=dsn=/dev/stdin -o yaml > crdb-dsn-secret.yml
```

```
echo -n <YOUR ADMIN KEY> | kubectl create secret \
  generic admin-key --dry-run=client \
  --from-file=key=/dev/stdin -o yaml > admin-key-secret.yml
```

Then, you can deploy the manifests:

```
kubectl apply -f ./
```

## GCP Cloud Run

Cloud Run can be a great way to run small to medium sized deployments (at large scale k8s is suggested for cost effectiveness).

## Configuration

### Environment Variables

Configuration is done through a few environment variables:

### `CACHE_TTL`

The TTL in milliseconds that the instance will cache API key queries. If set to `0`, caching is disabled. Default `0`.

More cache hits result in lower latency and higher concurrency per instance.

### `ADMIN_KEY`

A secure string should be provided to serve as the admin key to the admin HTTP endpoints. This is a temporary measure.


## Admin Endpoints

### `POST /key`

Create a new api key for a namespace.

Query params:
  - `ns`: The namespace for the api key

Returns:
```js
{
  "keyID": string,
  "keySecret": string
}
```

### `DELETE /key`

Delete an existing key

Query params:
  - `key`: The api `keyID`
