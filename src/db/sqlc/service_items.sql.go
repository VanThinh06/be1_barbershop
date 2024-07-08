// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: service_items.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createServiceItem = `-- name: CreateServiceItem :one
INSERT INTO "ServiceItems" (
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
RETURNING id, barber_shop_id, category_id, gender_id, name, timer, price, discount_price, discount_start_time, discount_end_time, description, image_url, is_active
`

type CreateServiceItemParams struct {
	BarberShopID uuid.UUID      `json:"barber_shop_id"`
	CategoryID   int16          `json:"category_id"`
	GenderID     int16          `json:"gender_id"`
	Name         string         `json:"name"`
	Timer        int16          `json:"timer"`
	Price        float32        `json:"price"`
	Description  sql.NullString `json:"description"`
	ImageUrl     sql.NullString `json:"image_url"`
}

func (q *Queries) CreateServiceItem(ctx context.Context, arg CreateServiceItemParams) (ServiceItem, error) {
	row := q.db.QueryRow(ctx, createServiceItem,
		arg.BarberShopID,
		arg.CategoryID,
		arg.GenderID,
		arg.Name,
		arg.Timer,
		arg.Price,
		arg.Description,
		arg.ImageUrl,
	)
	var i ServiceItem
	err := row.Scan(
		&i.ID,
		&i.BarberShopID,
		&i.CategoryID,
		&i.GenderID,
		&i.Name,
		&i.Timer,
		&i.Price,
		&i.DiscountPrice,
		&i.DiscountStartTime,
		&i.DiscountEndTime,
		&i.Description,
		&i.ImageUrl,
		&i.IsActive,
	)
	return i, err
}

const deleteServiceItem = `-- name: DeleteServiceItem :exec
DELETE FROM "ServiceItems"
WHERE
  "id" = $1
`

func (q *Queries) DeleteServiceItem(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteServiceItem, id)
	return err
}

const getServiceItem = `-- name: GetServiceItem :one
SELECT bs.id, bs.barber_shop_id, bs.category_id, bs.gender_id, bs.name, bs.timer, bs.price, bs.discount_price, bs.discount_start_time, bs.discount_end_time, bs.description, bs.image_url, bs.is_active, sc."name" AS "category_name"
FROM "ServiceItems" bs
LEFT JOIN 
    "ServiceCategories" sc ON sc."id" = bs."category_id"
WHERE bs."id" = $1
`

type GetServiceItemRow struct {
	ID                uuid.UUID        `json:"id"`
	BarberShopID      uuid.UUID        `json:"barber_shop_id"`
	CategoryID        int16            `json:"category_id"`
	GenderID          int16            `json:"gender_id"`
	Name              string           `json:"name"`
	Timer             int16            `json:"timer"`
	Price             float32          `json:"price"`
	DiscountPrice     pgtype.Float4    `json:"discount_price"`
	DiscountStartTime pgtype.Timestamp `json:"discount_start_time"`
	DiscountEndTime   pgtype.Timestamp `json:"discount_end_time"`
	Description       sql.NullString   `json:"description"`
	ImageUrl          sql.NullString   `json:"image_url"`
	IsActive          bool             `json:"is_active"`
	CategoryName      sql.NullString   `json:"category_name"`
}

func (q *Queries) GetServiceItem(ctx context.Context, id uuid.UUID) (GetServiceItemRow, error) {
	row := q.db.QueryRow(ctx, getServiceItem, id)
	var i GetServiceItemRow
	err := row.Scan(
		&i.ID,
		&i.BarberShopID,
		&i.CategoryID,
		&i.GenderID,
		&i.Name,
		&i.Timer,
		&i.Price,
		&i.DiscountPrice,
		&i.DiscountStartTime,
		&i.DiscountEndTime,
		&i.Description,
		&i.ImageUrl,
		&i.IsActive,
		&i.CategoryName,
	)
	return i, err
}

const getTimerServiceItem = `-- name: GetTimerServiceItem :one
SELECT SUM("timer") AS total_timer
FROM "ServiceItems"
WHERE "id" = ANY($1::uuid[])
`

func (q *Queries) GetTimerServiceItem(ctx context.Context, dollar_1 []uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, getTimerServiceItem, dollar_1)
	var total_timer int64
	err := row.Scan(&total_timer)
	return total_timer, err
}

const listServicesByCategory = `-- name: ListServicesByCategory :many
SELECT 
    sc."id" AS "category_id", 
    sc."name" AS "category_name", 
    bs."id" AS "service_id", 
    bs."name" AS "service_name",
    bs."timer",
    bs."price",
    bs."description",
    bs."image_url",
    bs."gender_id",
    bs."is_active",
    bs."discount_price",
    bs."discount_start_time",
    bs."discount_end_time"
FROM 
    "ServiceCategories" sc
LEFT JOIN 
    "ServiceItems" bs ON sc."id" = bs."category_id"
LEFT JOIN 
    "CategoryPositions" cp ON sc."id" = cp."category_id"
WHERE 
    bs."barber_shop_id" = $1
    AND (cp."visible" = true)  
ORDER BY
    cp."position",  -- Sắp xếp theo vị trí của category
    sc."id",
    bs."gender_id"
`

type ListServicesByCategoryRow struct {
	CategoryID        int16            `json:"category_id"`
	CategoryName      string           `json:"category_name"`
	ServiceID         uuid.NullUUID    `json:"service_id"`
	ServiceName       sql.NullString   `json:"service_name"`
	Timer             pgtype.Int2      `json:"timer"`
	Price             pgtype.Float4    `json:"price"`
	Description       sql.NullString   `json:"description"`
	ImageUrl          sql.NullString   `json:"image_url"`
	GenderID          pgtype.Int2      `json:"gender_id"`
	IsActive          pgtype.Bool      `json:"is_active"`
	DiscountPrice     pgtype.Float4    `json:"discount_price"`
	DiscountStartTime pgtype.Timestamp `json:"discount_start_time"`
	DiscountEndTime   pgtype.Timestamp `json:"discount_end_time"`
}

func (q *Queries) ListServicesByCategory(ctx context.Context, barberShopID uuid.UUID) ([]ListServicesByCategoryRow, error) {
	rows, err := q.db.Query(ctx, listServicesByCategory, barberShopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListServicesByCategoryRow{}
	for rows.Next() {
		var i ListServicesByCategoryRow
		if err := rows.Scan(
			&i.CategoryID,
			&i.CategoryName,
			&i.ServiceID,
			&i.ServiceName,
			&i.Timer,
			&i.Price,
			&i.Description,
			&i.ImageUrl,
			&i.GenderID,
			&i.IsActive,
			&i.DiscountPrice,
			&i.DiscountStartTime,
			&i.DiscountEndTime,
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

const updateServiceItem = `-- name: UpdateServiceItem :one


UPDATE "ServiceItems"
SET
  name = COALESCE($1, name),
  gender_id = COALESCE($2, gender_id),
  timer = COALESCE($3, timer),
  category_id = COALESCE($4, category_id),
  price = COALESCE($5, price),
  discount_price = COALESCE($6, 0),
  discount_start_time = $7,
  discount_end_time = $8,
  description = COALESCE($9, ''),
  image_url = COALESCE($10, ''),
  is_active = COALESCE($11, is_active)
WHERE id = $12
RETURNING id, barber_shop_id, category_id, gender_id, name, timer, price, discount_price, discount_start_time, discount_end_time, description, image_url, is_active
`

type UpdateServiceItemParams struct {
	Name              sql.NullString   `json:"name"`
	GenderID          pgtype.Int2      `json:"gender_id"`
	Timer             pgtype.Int2      `json:"timer"`
	CategoryID        pgtype.Int2      `json:"category_id"`
	Price             pgtype.Float4    `json:"price"`
	DiscountPrice     pgtype.Float4    `json:"discount_price"`
	DiscountStartTime pgtype.Timestamp `json:"discount_start_time"`
	DiscountEndTime   pgtype.Timestamp `json:"discount_end_time"`
	Description       sql.NullString   `json:"description"`
	ImageUrl          sql.NullString   `json:"image_url"`
	IsActive          pgtype.Bool      `json:"is_active"`
	ID                uuid.UUID        `json:"id"`
}

// Để phân loại theo gender_id nếu cần thiết
func (q *Queries) UpdateServiceItem(ctx context.Context, arg UpdateServiceItemParams) (ServiceItem, error) {
	row := q.db.QueryRow(ctx, updateServiceItem,
		arg.Name,
		arg.GenderID,
		arg.Timer,
		arg.CategoryID,
		arg.Price,
		arg.DiscountPrice,
		arg.DiscountStartTime,
		arg.DiscountEndTime,
		arg.Description,
		arg.ImageUrl,
		arg.IsActive,
		arg.ID,
	)
	var i ServiceItem
	err := row.Scan(
		&i.ID,
		&i.BarberShopID,
		&i.CategoryID,
		&i.GenderID,
		&i.Name,
		&i.Timer,
		&i.Price,
		&i.DiscountPrice,
		&i.DiscountStartTime,
		&i.DiscountEndTime,
		&i.Description,
		&i.ImageUrl,
		&i.IsActive,
	)
	return i, err
}
