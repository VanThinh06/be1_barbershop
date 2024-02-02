// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: service_categories.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createServiceCategories = `-- name: CreateServiceCategories :one
INSERT INTO "ServiceCategories" ("name", "is_global")
VALUES ($1, $2)
RETURNING id, name, is_global, create_at, update_at
`

type CreateServiceCategoriesParams struct {
	Name     string `json:"name"`
	IsGlobal bool   `json:"is_global"`
}

func (q *Queries) CreateServiceCategories(ctx context.Context, arg CreateServiceCategoriesParams) (ServiceCategory, error) {
	row := q.db.QueryRowContext(ctx, createServiceCategories, arg.Name, arg.IsGlobal)
	var i ServiceCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsGlobal,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteServiceCategories = `-- name: DeleteServiceCategories :exec
DELETE FROM "ServiceCategories"
WHERE "id" = $1
`

func (q *Queries) DeleteServiceCategories(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteServiceCategories, id)
	return err
}

const getServiceCategories = `-- name: GetServiceCategories :one
SELECT id, name, is_global, create_at, update_at
FROM "ServiceCategories"
WHERE "id" = $1
`

func (q *Queries) GetServiceCategories(ctx context.Context, id uuid.UUID) (ServiceCategory, error) {
	row := q.db.QueryRowContext(ctx, getServiceCategories, id)
	var i ServiceCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsGlobal,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const updateServiceCategories = `-- name: UpdateServiceCategories :one
UPDATE "ServiceCategories"
SET "name" = coalesce($1, name),
    "is_global" = coalesce($2, is_global),
    "update_at" = NOW()
WHERE "id" = $3
RETURNING id, name, is_global, create_at, update_at
`

type UpdateServiceCategoriesParams struct {
	Name     sql.NullString `json:"name"`
	IsGlobal sql.NullBool   `json:"is_global"`
	ID       uuid.UUID      `json:"id"`
}

func (q *Queries) UpdateServiceCategories(ctx context.Context, arg UpdateServiceCategoriesParams) (ServiceCategory, error) {
	row := q.db.QueryRowContext(ctx, updateServiceCategories, arg.Name, arg.IsGlobal, arg.ID)
	var i ServiceCategory
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.IsGlobal,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}
