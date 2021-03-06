---
sidebar_position: 4
---

# Install the Server

:::tip
If you don't want to manage the server and database yourself, [we are working on an amazing managed offering with a generous free tier](https://permissionpanther.com?ref=docs).
:::
If you want to setup Permission Panther to run on your own infrastructure, the steps are simple. [Binaries](https://github.com/TheSaaSZoo/PermissionPanther/releases) and [container images](https://github.com/TheSaaSZoo/PermissionPanther/pkgs/container/permissionpanther) are readily available for both amd64 (x86) and arm64 platforms.

:::caution Warning
You should always pin a version for production deployments, as major changes could be breaking.
:::

## Setup CockroachDB

Both the [Docker](#docker) and [Docker Compose](#docker-compose) methods run a local CockroachDB instance. For production deployments, you will want to set up something more long-term. You can check out their their Serverless offering, or run a self-managed cluster. [Refer to their guide on how to setup a self-hosted CRDB instance](https://www.cockroachlabs.com/docs/stable/deploy-cockroachdb-on-premises.html).

## Apply Database Migration (if upgrading)

**If this is the first time you are running Permission Panther, you can skip this step.**

When running for the first time, Permission Panther will attempt to apply the latest table schemas for you. This will only succeed if new tables are being created. Between major versions, table modifications may be made, which will not be backwards compatible. The included migration files will need to be run in order to upgrade the database schemas.

### Clone the Git Repo

```
git clone --depth 1 https://github.com/TheSaaSZoo/PermissionPanther
```

### Setup [`sql-migrate`](https://github.com/rubenv/sql-migrate)
Install [`sql-migrate`](https://github.com/rubenv/sql-migrate) and set the `CRDB_DSN` environment variable as your database DSN.

### Apply Migrations

If you are jumping a major version, there could be an adjustment to the table schema. You will need to run the sql migration files on your database.

From the root directory of the cloned repo, run:

```
sql-migrate up
```

## Container Image

A container image is hosted at `ghcr.io/thesaaszoo/permissionpanther`. See [GitHub](https://github.com/TheSaaSZoo/PermissionPanther/pkgs/container/permissionpanther) for tagged versions and available releases.

## Docker Compose

Docker Compose works great for simple deployments. You can run both CockroachDB and Permission Panther with the following command:

```sh
curl \
  "https://raw.githubusercontent.com/TheSaaSZoo/PermissionPanther/main/docker-compose.yml" \
  -o docker-compose.yml

docker compose up -d
```

Permission Panther will be available at `localhost:8080`.

## Kubernetes

Permission Panther runs stateless, so you can create a Kubernetes deployment that can easily scale horizontally to as many pods as needed.

### Clone the repo

To get started, pull the source code and change to the directory with the manifests:

```
git clone https://github.com/TheSaaSZoo/PermissionPanther.git

cd k8s
```

### Create secret manifests

Next, create Kubneretes secret manifests for the CRDB DSN and admin key:

```
echo -n "postgresql://..." | kubectl create secret \
  generic crdb-dsn --dry-run=client \
  --from-file=dsn=/dev/stdin -o yaml > crdb-dsn.yml
```

```
echo -n <YOUR ADMIN KEY> | kubectl create secret \
  generic admin-key --dry-run=client \
  --from-file=key=/dev/stdin -o yaml > admin-key.yml
```

Then, you can deploy the manifests:

```
kubectl apply -f .
```

With this configuration, Permission Panther is only accessible from within the cluster at:
```
permission-panther.default.svc.cluster.local:80
```

:::note Note
This is a very basic Kubernetes deployment. A production-ready deployment should include a [horizontal pod autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/) and should have the resources tuned for your specific needs.

If you wish to expose the service externally, that will need to be configured as well depending on where your cluster is running.
:::

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

### Auth

To use the following admin endpoints, the header `ak: <ADMIN_KEY>` needs to be used, where `<ADMIN_KEY>` is the `ADMIN_KEY` environment variable set in your Permission Panther deployment.

### `POST /key`

Create a new api key for a namespace.

Query params:
  - `ns` (string): The namespace for the api key
  - `mr` (int): The max recursions for the api key

Returns:
```js
{
  "keyID": string,
  "keySecret": string
}
```

Example Usage:
```
curl -X POST -H 'ak: <ADMIN_KEY>' \
  http://localhost:8080/key?ns=<NAMESPACE>?mr=<MAX_RECURSIONS>
```

Replacing `<ADMIN_KEY>`, `<MAX_RECURSIONS>` and `<NAMESPACE>` with the appropriate values.

### `DELETE /key`

Delete an existing key

Query params:
  - `key` (string): The api `keyID`

Example Usage:

```
curl -X DELETE -H 'ak: <ADMIN_KEY>' \
  http://localhost:8080/key?key=<KEY_ID>
```

Replacing `<ADMIN_KEY>` and `<KEY_ID>` with the appropriate values.
