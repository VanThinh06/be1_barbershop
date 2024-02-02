// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: barbershop_service_categories.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createBarberShopServiceCategories = `-- name: CreateBarberShopServiceCategories :one
INSERT INTO "BarberShopServiceCategories" (
  "barber_shop_chain_id",
  "barber_shop_id",
  "service_category_id"
)
VALUES (
  $1,
  $2,
  $3
)
RETURNING id, barber_shop_chain_id, barber_shop_id, service_category_id, create_at, update_at
`

type CreateBarberShopServiceCategoriesParams struct {
	BarberShopChainID uuid.NullUUID `json:"barber_shop_chain_id"`
	BarberShopID      uuid.NullUUID `json:"barber_shop_id"`
	ServiceCategoryID uuid.UUID     `json:"service_category_id"`
}

func (q *Queries) CreateBarberShopServiceCategories(ctx context.Context, arg CreateBarberShopServiceCategoriesParams) (BarberShopServiceCategory, error) {
	row := q.db.QueryRowContext(ctx, createBarberShopServiceCategories, arg.BarberShopChainID, arg.BarberShopID, arg.ServiceCategoryID)
	var i BarberShopServiceCategory
	err := row.Scan(
		&i.ID,
		&i.BarberShopChainID,
		&i.BarberShopID,
		&i.ServiceCategoryID,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteBarberShopServiceCategories = `-- name: DeleteBarberShopServiceCategories :exec
DELETE FROM "BarberShopServiceCategories"
WHERE
  "id" = $1
`

func (q *Queries) DeleteBarberShopServiceCategories(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteBarberShopServiceCategories, id)
	return err
}

const listBarberShopServiceCategories = `-- name: ListBarberShopServiceCategories :many
SELECT
  bsc.id, bsc.barber_shop_chain_id, bsc.barber_shop_id, bsc.service_category_id, bsc.create_at, bsc.update_at,
  sc."name" as "service_category_name",
  sc."is_global" as "service_category_is_global"
FROM
  "BarberShopServiceCategories" bsc
JOIN
  "ServiceCategories" sc ON bsc."service_category_id" = sc."id"
WHERE
  bsc."barber_shop_id" = $1
  AND bsc."barber_shop_chain_id" = $2
`

type ListBarberShopServiceCategoriesParams struct {
	BarberShopID      uuid.NullUUID `json:"barber_shop_id"`
	BarberShopChainID uuid.NullUUID `json:"barber_shop_chain_id"`
}

type ListBarberShopServiceCategoriesRow struct {
	ID                      uuid.UUID     `json:"id"`
	BarberShopChainID       uuid.NullUUID `json:"barber_shop_chain_id"`
	BarberShopID            uuid.NullUUID `json:"barber_shop_id"`
	ServiceCategoryID       uuid.UUID     `json:"service_category_id"`
	CreateAt                time.Time     `json:"create_at"`
	UpdateAt                time.Time     `json:"update_at"`
	ServiceCategoryName     string        `json:"service_category_name"`
	ServiceCategoryIsGlobal bool          `json:"service_category_is_global"`
}

func (q *Queries) ListBarberShopServiceCategories(ctx context.Context, arg ListBarberShopServiceCategoriesParams) ([]ListBarberShopServiceCategoriesRow, error) {
	rows, err := q.db.QueryContext(ctx, listBarberShopServiceCategories, arg.BarberShopID, arg.BarberShopChainID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListBarberShopServiceCategoriesRow{}
	for rows.Next() {
		var i ListBarberShopServiceCategoriesRow
		if err := rows.Scan(
			&i.ID,
			&i.BarberShopChainID,
			&i.BarberShopID,
			&i.ServiceCategoryID,
			&i.CreateAt,
			&i.UpdateAt,
			&i.ServiceCategoryName,
			&i.ServiceCategoryIsGlobal,
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

const updateBarberShopServiceCategories = `-- name: UpdateBarberShopServiceCategories :one
UPDATE "BarberShopServiceCategories"
SET
  "service_category_id" = $1
WHERE
  "id" = $2
RETURNING id, barber_shop_chain_id, barber_shop_id, service_category_id, create_at, update_at
`

type UpdateBarberShopServiceCategoriesParams struct {
	ServiceCategoryID uuid.UUID `json:"service_category_id"`
	ID                uuid.UUID `json:"id"`
}

func (q *Queries) UpdateBarberShopServiceCategories(ctx context.Context, arg UpdateBarberShopServiceCategoriesParams) (BarberShopServiceCategory, error) {
	row := q.db.QueryRowContext(ctx, updateBarberShopServiceCategories, arg.ServiceCategoryID, arg.ID)
	var i BarberShopServiceCategory
	err := row.Scan(
		&i.ID,
		&i.BarberShopChainID,
		&i.BarberShopID,
		&i.ServiceCategoryID,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}
