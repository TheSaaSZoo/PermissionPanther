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
  -p 8080:8080 \
  -it --rm
  permissionpanther
```

## Docker Compose

Docker Compose works great for simple deployments:

```yml
version: '3.9'
services:
  permissionpanther:
    build: .
    environment:
      - ADMIN_KEY=CHANGE_ME!!!
      - CRDB_DSN=postgresql://root@crdb:26257/defaultdb?sslmode=disable
    ports:
      - 8080:8080
    restart: always
    depends_on:
      - crdb
  crdb:
    image: cockroachdb/cockroach:latest
    command: start-single-node --insecure
    ports:
      - 26257:26257
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
