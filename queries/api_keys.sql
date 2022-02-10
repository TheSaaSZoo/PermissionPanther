-- name: InsertAPIKey :exec
INSERT INTO keys (id, secret_hash, ns)
VALUES ($1, $2, $3);

-- name: DeleteAPIKey :execrows
DELETE FROM keys
WHERE id = $1;

-- name: ListAPIKeysForNS :many
SELECT *
FROM keys
WHERE ns = $1;

-- name: SelectAPIKeyHash :one
SELECT secret_hash
FROM keys
WHERE id = $1;
