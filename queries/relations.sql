-- name: CheckRelationDirect :one
SELECT *
FROM relations
WHERE ns = $1
AND entity = $2
AND permission = $3
AND object = $4;

-- name: GetGroupRelations :many
SELECT *
FROM relations
WHERE ns = $1
AND entity > '~'
AND permission = $2
AND object = $3;
