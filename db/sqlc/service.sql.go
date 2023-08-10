// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: service.sql

package db

import (
	"context"

	"github.com/google/uuid"
	null "gopkg.in/guregu/null.v4"
)

const createService = `-- name: CreateService :one
INSERT INTO service (
    service_category_id,
    work,
    timer,
    price,
    description,
    image
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, service_category_id, work, timer, price, description, image, created_at, update_at
`

type CreateServiceParams struct {
	ServiceCategoryID uuid.UUID   `json:"service_category_id"`
	Work              string      `json:"work"`
	Timer             null.Int    `json:"timer"`
	Price             float32     `json:"price"`
	Description       null.String `json:"description"`
	Image             null.String `json:"image"`
}

func (q *Queries) CreateService(ctx context.Context, arg CreateServiceParams) (Service, error) {
	row := q.db.QueryRowContext(ctx, createService,
		arg.ServiceCategoryID,
		arg.Work,
		arg.Timer,
		arg.Price,
		arg.Description,
		arg.Image,
	)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.ServiceCategoryID,
		&i.Work,
		&i.Timer,
		&i.Price,
		&i.Description,
		&i.Image,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteService = `-- name: DeleteService :one
DELETE FROM "service"
WHERE id = $1
RETURNING id, service_category_id, work, timer, price, description, image, created_at, update_at
`

func (q *Queries) DeleteService(ctx context.Context, id uuid.UUID) (Service, error) {
	row := q.db.QueryRowContext(ctx, deleteService, id)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.ServiceCategoryID,
		&i.Work,
		&i.Timer,
		&i.Price,
		&i.Description,
		&i.Image,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteServicewithStoreCategory = `-- name: DeleteServicewithStoreCategory :exec
DELETE FROM "service"
WHERE service_category_id = $1
`

func (q *Queries) DeleteServicewithStoreCategory(ctx context.Context, serviceCategoryID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteServicewithStoreCategory, serviceCategoryID)
	return err
}

const getListServicewithCategory = `-- name: GetListServicewithCategory :many
SELECT id, service_category_id, work, timer, price, description, image, created_at, update_at
FROM "service"
WHERE service_category_id = $1
LIMIT $2 OFFSET $3
`

type GetListServicewithCategoryParams struct {
	ServiceCategoryID uuid.UUID `json:"service_category_id"`
	Limit             int32     `json:"limit"`
	Offset            int32     `json:"offset"`
}

func (q *Queries) GetListServicewithCategory(ctx context.Context, arg GetListServicewithCategoryParams) ([]Service, error) {
	rows, err := q.db.QueryContext(ctx, getListServicewithCategory, arg.ServiceCategoryID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Service{}
	for rows.Next() {
		var i Service
		if err := rows.Scan(
			&i.ID,
			&i.ServiceCategoryID,
			&i.Work,
			&i.Timer,
			&i.Price,
			&i.Description,
			&i.Image,
			&i.CreatedAt,
			&i.UpdateAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getService = `-- name: GetService :one
SELECT id, service_category_id, work, timer, price, description, image, created_at, update_at
FROM "service"
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetService(ctx context.Context, id uuid.UUID) (Service, error) {
	row := q.db.QueryRowContext(ctx, getService, id)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.ServiceCategoryID,
		&i.Work,
		&i.Timer,
		&i.Price,
		&i.Description,
		&i.Image,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const updateService = `-- name: UpdateService :one
UPDATE "service"
set service_category_id = $2,
  work = $3,
  timer = $4,
  price = $5,
  description = $6,
  image = $7
WHERE id = $1
RETURNING id, service_category_id, work, timer, price, description, image, created_at, update_at
`

type UpdateServiceParams struct {
	ID                uuid.UUID   `json:"id"`
	ServiceCategoryID uuid.UUID   `json:"service_category_id"`
	Work              string      `json:"work"`
	Timer             null.Int    `json:"timer"`
	Price             float32     `json:"price"`
	Description       null.String `json:"description"`
	Image             null.String `json:"image"`
}

func (q *Queries) UpdateService(ctx context.Context, arg UpdateServiceParams) (Service, error) {
	row := q.db.QueryRowContext(ctx, updateService,
		arg.ID,
		arg.ServiceCategoryID,
		arg.Work,
		arg.Timer,
		arg.Price,
		arg.Description,
		arg.Image,
	)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.ServiceCategoryID,
		&i.Work,
		&i.Timer,
		&i.Price,
		&i.Description,
		&i.Image,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}
