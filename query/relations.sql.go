// Code generated by sqlc. DO NOT EDIT.
// source: relations.sql

package query

import (
	"context"
)

const checkRelationDirect = `-- name: CheckRelationDirect :one
SELECT 1
FROM relations
WHERE ns = $1
AND entity = $2
AND permission = $3
AND object = $4
`

type CheckRelationDirectParams struct {
	Ns         string
	Entity     string
	Permission string
	Object     string
}

func (q *Queries) CheckRelationDirect(ctx context.Context, arg CheckRelationDirectParams) (interface{}, error) {
	row := q.db.QueryRow(ctx, checkRelationDirect,
		arg.Ns,
		arg.Entity,
		arg.Permission,
		arg.Object,
	)
	var column_1 interface{}
	err := row.Scan(&column_1)
	return column_1, err
}

const deleteRelation = `-- name: DeleteRelation :exec
DELETE FROM relations
WHERE ns = $1
AND entity = $2
AND permission = $3
AND object = $4
`

type DeleteRelationParams struct {
	Ns         string
	Entity     string
	Permission string
	Object     string
}

func (q *Queries) DeleteRelation(ctx context.Context, arg DeleteRelationParams) error {
	_, err := q.db.Exec(ctx, deleteRelation,
		arg.Ns,
		arg.Entity,
		arg.Permission,
		arg.Object,
	)
	return err
}

const getGroupRelations = `-- name: GetGroupRelations :many
SELECT object, entity, permission, ns
FROM relations
WHERE ns = $1
AND entity > '~'
AND permission = $2
AND object = $3
`

type GetGroupRelationsParams struct {
	Ns         string
	Permission string
	Object     string
}

func (q *Queries) GetGroupRelations(ctx context.Context, arg GetGroupRelationsParams) ([]Relation, error) {
	rows, err := q.db.Query(ctx, getGroupRelations, arg.Ns, arg.Permission, arg.Object)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Relation
	for rows.Next() {
		var i Relation
		if err := rows.Scan(
			&i.Object,
			&i.Entity,
			&i.Permission,
			&i.Ns,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertRelation = `-- name: InsertRelation :exec
INSERT INTO relations (ns, entity, permission, object)
VALUES ($1, $2, $3, $4)
ON CONFLICT DO NOTHING
`

type InsertRelationParams struct {
	Ns         string
	Entity     string
	Permission string
	Object     string
}

func (q *Queries) InsertRelation(ctx context.Context, arg InsertRelationParams) error {
	_, err := q.db.Exec(ctx, insertRelation,
		arg.Ns,
		arg.Entity,
		arg.Permission,
		arg.Object,
	)
	return err
}

const listEntityRelations = `-- name: ListEntityRelations :many
SELECT object, entity, permission, ns
FROM relations
WHERE ns = $1
AND entity = $2
`

type ListEntityRelationsParams struct {
	Ns     string
	Entity string
}

func (q *Queries) ListEntityRelations(ctx context.Context, arg ListEntityRelationsParams) ([]Relation, error) {
	rows, err := q.db.Query(ctx, listEntityRelations, arg.Ns, arg.Entity)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Relation
	for rows.Next() {
		var i Relation
		if err := rows.Scan(
			&i.Object,
			&i.Entity,
			&i.Permission,
			&i.Ns,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listEntityRelationsWithPermission = `-- name: ListEntityRelationsWithPermission :many
SELECT object, entity, permission, ns
FROM relations
WHERE ns = $1
AND entity = $2
AND permission = $3
`

type ListEntityRelationsWithPermissionParams struct {
	Ns         string
	Entity     string
	Permission string
}

func (q *Queries) ListEntityRelationsWithPermission(ctx context.Context, arg ListEntityRelationsWithPermissionParams) ([]Relation, error) {
	rows, err := q.db.Query(ctx, listEntityRelationsWithPermission, arg.Ns, arg.Entity, arg.Permission)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Relation
	for rows.Next() {
		var i Relation
		if err := rows.Scan(
			&i.Object,
			&i.Entity,
			&i.Permission,
			&i.Ns,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listObjectRelations = `-- name: ListObjectRelations :many
SELECT object, entity, permission, ns
FROM relations
WHERE ns = $1
AND object = $2
`

type ListObjectRelationsParams struct {
	Ns     string
	Object string
}

func (q *Queries) ListObjectRelations(ctx context.Context, arg ListObjectRelationsParams) ([]Relation, error) {
	rows, err := q.db.Query(ctx, listObjectRelations, arg.Ns, arg.Object)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Relation
	for rows.Next() {
		var i Relation
		if err := rows.Scan(
			&i.Object,
			&i.Entity,
			&i.Permission,
			&i.Ns,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listObjectRelationsWithPermission = `-- name: ListObjectRelationsWithPermission :many
SELECT object, entity, permission, ns
FROM relations
WHERE ns = $1
AND object = $2
AND permission = $3
`

type ListObjectRelationsWithPermissionParams struct {
	Ns         string
	Object     string
	Permission string
}

func (q *Queries) ListObjectRelationsWithPermission(ctx context.Context, arg ListObjectRelationsWithPermissionParams) ([]Relation, error) {
	rows, err := q.db.Query(ctx, listObjectRelationsWithPermission, arg.Ns, arg.Object, arg.Permission)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Relation
	for rows.Next() {
		var i Relation
		if err := rows.Scan(
			&i.Object,
			&i.Entity,
			&i.Permission,
			&i.Ns,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
