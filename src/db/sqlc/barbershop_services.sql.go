// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: barbershop_services.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createBarberShopServices = `-- name: CreateBarberShopServices :one
INSERT INTO "BarberShopServices" (
    barbershop_category_id,
    barbershop_chain_id,
    barbershop_id,
    gender_id,
    "name",
    "timer",
    price,
    description,
    image
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, barbershop_category_id, barbershop_chain_id, barbershop_id, gender_id, name, timer, price, description, image, is_custom, is_removed, create_at, update_at
`

type CreateBarberShopServicesParams struct {
	BarbershopCategoryID uuid.UUID      `json:"barbershop_category_id"`
	BarbershopChainID    uuid.NullUUID  `json:"barbershop_chain_id"`
	BarbershopID         uuid.NullUUID  `json:"barbershop_id"`
	GenderID             sql.NullInt32  `json:"gender_id"`
	Name                 string         `json:"name"`
	Timer                int32          `json:"timer"`
	Price                float32        `json:"price"`
	Description          sql.NullString `json:"description"`
	Image                sql.NullString `json:"image"`
}

func (q *Queries) CreateBarberShopServices(ctx context.Context, arg CreateBarberShopServicesParams) (BarberShopService, error) {
	row := q.db.QueryRowContext(ctx, createBarberShopServices,
		arg.BarbershopCategoryID,
		arg.BarbershopChainID,
		arg.BarbershopID,
		arg.GenderID,
		arg.Name,
		arg.Timer,
		arg.Price,
		arg.Description,
		arg.Image,
	)
	var i BarberShopService
	err := row.Scan(
		&i.ID,
		&i.BarbershopCategoryID,
		&i.BarbershopChainID,
		&i.BarbershopID,
		&i.GenderID,
		&i.Name,
		&i.Timer,
		&i.Price,
		&i.Description,
		&i.Image,
		&i.IsCustom,
		&i.IsRemoved,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteBarberShopServices = `-- name: DeleteBarberShopServices :exec
DELETE FROM "BarberShopServices"
WHERE
  "id" = $1
`

func (q *Queries) DeleteBarberShopServices(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteBarberShopServices, id)
	return err
}
