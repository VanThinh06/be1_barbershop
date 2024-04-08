// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: barbershop_services.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createBarberShopServices = `-- name: CreateBarberShopServices :one
INSERT INTO "BarberShopServices" (
    barbershop_category_id
  )
VALUES ($1)
RETURNING id, barbershop_category_id, gender_id, name, timer, price, description, image_url
`

func (q *Queries) CreateBarberShopServices(ctx context.Context, barbershopCategoryID uuid.UUID) (BarberShopService, error) {
	row := q.db.QueryRow(ctx, createBarberShopServices, barbershopCategoryID)
	var i BarberShopService
	err := row.Scan(
		&i.ID,
		&i.BarbershopCategoryID,
		&i.GenderID,
		&i.Name,
		&i.Timer,
		&i.Price,
		&i.Description,
		&i.ImageUrl,
	)
	return i, err
}

const deleteBarberShopServices = `-- name: DeleteBarberShopServices :exec
DELETE FROM "BarberShopServices"
WHERE
  "id" = $1
`

func (q *Queries) DeleteBarberShopServices(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteBarberShopServices, id)
	return err
}
