// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: barbershop_services.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createBarberShopService = `-- name: CreateBarberShopService :one
INSERT INTO "BarberShopServices" (
barber_shop_id,
category_id,
gender_id,
name,
timer,
price,
description,
image_url
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, barber_shop_id, category_id, gender_id, name, timer, price, description, image_url
`

type CreateBarberShopServiceParams struct {
	BarberShopID uuid.UUID      `json:"barber_shop_id"`
	CategoryID   int16          `json:"category_id"`
	GenderID     int16          `json:"gender_id"`
	Name         string         `json:"name"`
	Timer        int16          `json:"timer"`
	Price        float32        `json:"price"`
	Description  sql.NullString `json:"description"`
	ImageUrl     sql.NullString `json:"image_url"`
}

func (q *Queries) CreateBarberShopService(ctx context.Context, arg CreateBarberShopServiceParams) (BarberShopService, error) {
	row := q.db.QueryRow(ctx, createBarberShopService,
		arg.BarberShopID,
		arg.CategoryID,
		arg.GenderID,
		arg.Name,
		arg.Timer,
		arg.Price,
		arg.Description,
		arg.ImageUrl,
	)
	var i BarberShopService
	err := row.Scan(
		&i.ID,
		&i.BarberShopID,
		&i.CategoryID,
		&i.GenderID,
		&i.Name,
		&i.Timer,
		&i.Price,
		&i.Description,
		&i.ImageUrl,
	)
	return i, err
}
