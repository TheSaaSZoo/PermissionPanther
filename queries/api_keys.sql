-- name: InsertAPIKey :exec
INSERT INTO keys (id, secret_hash, ns, max_recursions)
VALUES ($1, $2, $3, $4);

-- name: DeleteAPIKey :execrows
DELETE FROM keys
WHERE id = $1;

-- name: ListAPIKeysForNS :many
SELECT *
FROM keys
WHERE ns = $1;

-- name: SelectAPIKey :one
SELECT secret_hash, ns, max_recursions
FROM keys
WHERE id = $1;
