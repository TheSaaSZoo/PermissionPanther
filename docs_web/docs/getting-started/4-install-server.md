---
sidebar_position: 4
---

# Install the Server

:::tip
If you don't want to manage the server and database yourself, [we have an amazing managed offering with a generous free tier](https://permissionpanther.com).
:::
If you want to setup Permission Panther to run on your own infrastructure, the steps are simple.

## Setup CockroachDB

You need to setup a CockroachDB instance, whether it be their Serverless offering or a self-managed cluster. [Refer to their guide on how to setup a self-hosted CRDB instance](https://www.cockroachlabs.com/docs/stable/deploy-cockroachdb-on-premises.html)

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
