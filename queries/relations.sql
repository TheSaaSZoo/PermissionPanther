-- name: CheckRelationDirect :one
SELECT 1
FROM relations
WHERE ns = $1
AND entity IN ($2, '*')
AND permission = $3
AND object IN ($4, '*');

-- name: GetGroupRelations :many
SELECT *
FROM relations
WHERE ns = $1
AND entity > '~'
AND permission = $2
AND object = $3;

-- name: ListEntityRelations :many
SELECT *
FROM relations
WHERE ns = $1
AND entity = $2;

-- name: ListEntityRelationsWithPermission :many
SELECT *
FROM relations
WHERE ns = $1
AND entity = $2
AND permission = $3;

-- name: ListObjectRelations :many
SELECT *
FROM relations
WHERE ns = $1
AND object = $2;

-- name: ListObjectRelationsWithPermission :many
SELECT *
FROM relations
WHERE ns = $1
AND object = $2
AND permission = $3;

-- name: InsertRelation :exec
INSERT INTO relations (ns, entity, permission, object)
VALUES ($1, $2, $3, $4)
ON CONFLICT DO NOTHING;

-- name: DeleteRelation :exec
DELETE FROM relations
WHERE ns = $1
AND entity = $2
AND permission = $3
AND object = $4;
