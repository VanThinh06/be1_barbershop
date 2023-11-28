// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: services.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createService = `-- name: CreateService :one
INSERT INTO "Services" (
    category_id,
    "name",
    timer,
    price,
    "description",
    "image"
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, category_id, name, timer, price, description, image, created_at, update_at
`

type CreateServiceParams struct {
	CategoryID  uuid.UUID       `json:"category_id"`
	Name        string          `json:"name"`
	Timer       sql.NullInt32   `json:"timer"`
	Price       sql.NullFloat64 `json:"price"`
	Description sql.NullString  `json:"description"`
	Image       sql.NullString  `json:"image"`
}

func (q *Queries) CreateService(ctx context.Context, arg CreateServiceParams) (Service, error) {
	row := q.db.QueryRowContext(ctx, createService,
		arg.CategoryID,
		arg.Name,
		arg.Timer,
		arg.Price,
		arg.Description,
		arg.Image,
	)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.CategoryID,
		&i.Name,
		&i.Timer,
		&i.Price,
		&i.Description,
		&i.Image,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}