// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: service_packages.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createServicePackage = `-- name: CreateServicePackage :one
INSERT INTO "ServicePackages" (
barber_shop_id,
name,
gender_id,
timer,
price,
description,
image_url
  )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, barber_shop_id, name, gender_id, timer, price, discount_price, discount_start_time, discount_end_time, description, image_url, is_active
`

type CreateServicePackageParams struct {
	BarberShopID uuid.UUID      `json:"barber_shop_id"`
	Name         string         `json:"name"`
	GenderID     int16          `json:"gender_id"`
	Timer        int16          `json:"timer"`
	Price        float32        `json:"price"`
	Description  sql.NullString `json:"description"`
	ImageUrl     sql.NullString `json:"image_url"`
}

func (q *Queries) CreateServicePackage(ctx context.Context, arg CreateServicePackageParams) (ServicePackage, error) {
	row := q.db.QueryRow(ctx, createServicePackage,
		arg.BarberShopID,
		arg.Name,
		arg.GenderID,
		arg.Timer,
		arg.Price,
		arg.Description,
		arg.ImageUrl,
	)
	var i ServicePackage
	err := row.Scan(
		&i.ID,
		&i.BarberShopID,
		&i.Name,
		&i.GenderID,
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

const createServicePackageItem = `-- name: CreateServicePackageItem :one
INSERT INTO "ServicePackageItems" (
service_package_id,
service_item_id
  )
VALUES ($1, $2)
RETURNING id, service_package_id, service_item_id
`

type CreateServicePackageItemParams struct {
	ServicePackageID uuid.UUID `json:"service_package_id"`
	ServiceItemID    uuid.UUID `json:"service_item_id"`
}

func (q *Queries) CreateServicePackageItem(ctx context.Context, arg CreateServicePackageItemParams) (ServicePackageItem, error) {
	row := q.db.QueryRow(ctx, createServicePackageItem, arg.ServicePackageID, arg.ServiceItemID)
	var i ServicePackageItem
	err := row.Scan(&i.ID, &i.ServicePackageID, &i.ServiceItemID)
	return i, err
}

const getServicePackage = `-- name: GetServicePackage :many
SELECT 
  cs.id,
  cs.name AS combo_service_name,
  cs.description AS combo_service_description,
  cs.price AS combo_service_price,
  cs.image_url AS combo_service_image_url,
  cs.timer AS combo_service_timer,
  cs.gender_id AS combo_service_gender,
  cs.is_active AS combo_service_is_active,
  bss.id AS barber_shop_service_id,
  bss.name AS barber_shop_service_name,
  bss.price AS barber_shop_service_price
FROM 
  "ServicePackageItems" csi
JOIN 
  "ServicePackages" cs ON csi.combo_service_id = cs.id
JOIN 
  "ServiceItems" bss ON csi.barber_shop_service_id = bss.id
WHERE 
  cs.id = $1
`

type GetServicePackageRow struct {
	ID                      uuid.UUID      `json:"id"`
	ComboServiceName        string         `json:"combo_service_name"`
	ComboServiceDescription sql.NullString `json:"combo_service_description"`
	ComboServicePrice       float32        `json:"combo_service_price"`
	ComboServiceImageUrl    sql.NullString `json:"combo_service_image_url"`
	ComboServiceTimer       int16          `json:"combo_service_timer"`
	ComboServiceGender      int16          `json:"combo_service_gender"`
	ComboServiceIsActive    bool           `json:"combo_service_is_active"`
	BarberShopServiceID     uuid.UUID      `json:"barber_shop_service_id"`
	BarberShopServiceName   string         `json:"barber_shop_service_name"`
	BarberShopServicePrice  float32        `json:"barber_shop_service_price"`
}

func (q *Queries) GetServicePackage(ctx context.Context, id uuid.UUID) ([]GetServicePackageRow, error) {
	rows, err := q.db.Query(ctx, getServicePackage, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetServicePackageRow{}
	for rows.Next() {
		var i GetServicePackageRow
		if err := rows.Scan(
			&i.ID,
			&i.ComboServiceName,
			&i.ComboServiceDescription,
			&i.ComboServicePrice,
			&i.ComboServiceImageUrl,
			&i.ComboServiceTimer,
			&i.ComboServiceGender,
			&i.ComboServiceIsActive,
			&i.BarberShopServiceID,
			&i.BarberShopServiceName,
			&i.BarberShopServicePrice,
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

const listServicePackages = `-- name: ListServicePackages :many
SELECT id, barber_shop_id, combo_service_gender, combo_service_name, combo_service_timer, combo_service_price, combo_service_discount_price, combo_service_discount_start_time, combo_service_discount_end_time, combo_service_description, combo_service_image_url, combo_service_is_active, service_item_ids FROM "view_service_packages"
WHERE barber_shop_id = $1
`

func (q *Queries) ListServicePackages(ctx context.Context, barberShopID uuid.UUID) ([]ViewServicePackage, error) {
	rows, err := q.db.Query(ctx, listServicePackages, barberShopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ViewServicePackage{}
	for rows.Next() {
		var i ViewServicePackage
		if err := rows.Scan(
			&i.ID,
			&i.BarberShopID,
			&i.ComboServiceGender,
			&i.ComboServiceName,
			&i.ComboServiceTimer,
			&i.ComboServicePrice,
			&i.ComboServiceDiscountPrice,
			&i.ComboServiceDiscountStartTime,
			&i.ComboServiceDiscountEndTime,
			&i.ComboServiceDescription,
			&i.ComboServiceImageUrl,
			&i.ComboServiceIsActive,
			&i.ServiceItemIds,
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

const updateServicePackage = `-- name: UpdateServicePackage :one
UPDATE "ServicePackages"
SET
  name = COALESCE($1, name),
  gender_id = COALESCE($2, gender_id),
  timer = COALESCE($3, timer),
  price = COALESCE($4, price),
  discount_price = COALESCE($5, 0),
  discount_start_time = $6,
  discount_end_time = $7,
  description = COALESCE($8, ''),
  image_url = COALESCE($9, ''),
  is_active = COALESCE($10, is_active)
WHERE id = $11
RETURNING id, barber_shop_id, name, gender_id, timer, price, discount_price, discount_start_time, discount_end_time, description, image_url, is_active
`

type UpdateServicePackageParams struct {
	Name              sql.NullString   `json:"name"`
	GenderID          pgtype.Int2      `json:"gender_id"`
	Timer             pgtype.Int2      `json:"timer"`
	Price             pgtype.Float4    `json:"price"`
	DiscountPrice     pgtype.Float4    `json:"discount_price"`
	DiscountStartTime pgtype.Timestamp `json:"discount_start_time"`
	DiscountEndTime   pgtype.Timestamp `json:"discount_end_time"`
	Description       sql.NullString   `json:"description"`
	ImageUrl          sql.NullString   `json:"image_url"`
	IsActive          pgtype.Bool      `json:"is_active"`
	ID                uuid.UUID        `json:"id"`
}

func (q *Queries) UpdateServicePackage(ctx context.Context, arg UpdateServicePackageParams) (ServicePackage, error) {
	row := q.db.QueryRow(ctx, updateServicePackage,
		arg.Name,
		arg.GenderID,
		arg.Timer,
		arg.Price,
		arg.DiscountPrice,
		arg.DiscountStartTime,
		arg.DiscountEndTime,
		arg.Description,
		arg.ImageUrl,
		arg.IsActive,
		arg.ID,
	)
	var i ServicePackage
	err := row.Scan(
		&i.ID,
		&i.BarberShopID,
		&i.Name,
		&i.GenderID,
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

const upsertServicePackageItem = `-- name: UpsertServicePackageItem :exec
WITH upserted AS (
  INSERT INTO "ServicePackageItems" (service_package_id, service_item_id)
  SELECT $1, unnest($2::uuid[])
  ON CONFLICT (service_package_id, service_item_id) DO NOTHING
)
DELETE FROM "ServicePackageItems" AS csi
WHERE csi.service_package_id = $1
  AND csi.service_item_id NOT IN (
    SELECT unnest($2::uuid[])
)
`

type UpsertServicePackageItemParams struct {
	ServicePackageID uuid.UUID   `json:"service_package_id"`
	Column2          []uuid.UUID `json:"column_2"`
}

func (q *Queries) UpsertServicePackageItem(ctx context.Context, arg UpsertServicePackageItemParams) error {
	_, err := q.db.Exec(ctx, upsertServicePackageItem, arg.ServicePackageID, arg.Column2)
	return err
}
