-- name: CheckRelationDirect :one
SELECT 1
FROM relations
WHERE ns = $1
AND entity = $2
AND permission = $3
AND object = $4;

-- name: CheckRelationWildcard :one
SELECT 1
WHERE EXISTS (
    SELECT 1 FROM relations
    WHERE relations.ns = $1
    AND relations.entity = $2
    AND relations.permission = $3
    AND relations.object = $4
) OR EXISTS (
    SELECT 1 FROM relations
    WHERE relations.ns = $1
    AND relations.entity = '*'
    AND relations.permission = $3
    AND relations.object = $4
) OR EXISTS (
    SELECT 1 FROM relations
    WHERE relations.ns = $1
    AND relations.entity = $2
    AND relations.permission = $3
    AND relations.object = '*'
) OR EXISTS (
    SELECT 1 FROM relations
    WHERE relations.ns = $1
    AND relations.entity = $2
    AND relations.permission = '*'
    AND relations.object = $4
);

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

-- name: InsertRelation :execrows
INSERT INTO relations (ns, entity, permission, object)
VALUES ($1, $2, $3, $4)
ON CONFLICT DO NOTHING;

-- name: DeleteRelation :execrows
DELETE FROM relations
WHERE ns = $1
AND entity = $2
AND permission = $3
AND object = $4;

-- name: InsertPermissionGroup :exec
INSERT INTO permission_groups (ns, perms, name)
VALUES ($1, $2, $3);

-- name: DeletePermissionGroup :one
DELETE FROM permission_groups
WHERE ns = $1
AND name = $2
RETURNING perms;

-- name: SelectPermissionGroup :one
SELECT * FROM permission_groups
WHERE ns = $1
AND name = $2;

-- name: AddPermissionToGroup :execrows
-- Adds a permission if it does not exist
UPDATE permission_groups
SET perms = array_append(perms, @permission)
WHERE array_position(perms, @permission) IS NULL
AND ns = $2
AND name = $1;

-- name: RemovePermissionFromGroup :execrows
-- Removes a permission if it exists
UPDATE permission_groups
SET perms = array_remove(perms, @permission)
WHERE array_position(perms, @permission) IS NOT NULL
AND ns = $2
AND name = $1;

-- name: AddMemberToPermissionGroup :exec
INSERT INTO permission_group_membership (group_name, entity, ns, object)
VALUES ($1, $2, $3, $4);

-- name: RemoveMemberFromPermissionGroup :execrows
DELETE FROM permission_group_membership
WHERE ns = $1
AND group_name = $2
AND entity = $3;

-- name: ListEntitiesInPermissionGroup :many
SELECT *
FROM permission_group_membership
WHERE entity > $1
AND ns = $2
AND group_name = $3
LIMIT 50;
