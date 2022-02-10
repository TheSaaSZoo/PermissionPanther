-- name: InsertAPIKey :exec
INSERT INTO keys (secret_hash, ns)
VALUES ($1, $2);

-- name: DeleteAPIKey :execrows
DELETE FROM keys
WHERE secret_hash = $1;

-- name: ListAPIKeysForNS :many
SELECT *
FROM keys
WHERE ns = $1;

-- name: SelectAPIKeyNS :one
SELECT ns
FROM keys
WHERE secret_hash = $1;
