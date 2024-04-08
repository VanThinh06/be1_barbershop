// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: barbershop_service_categories.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createBarberShopServiceCategories = `-- name: CreateBarberShopServiceCategories :one
INSERT INTO "BarberShopServiceCategories" (
  "barber_shop_id",
  "service_category_id"
)
VALUES (
  $1,
  $2
)
RETURNING id, barber_shop_id, service_category_id
`

type CreateBarberShopServiceCategoriesParams struct {
	BarberShopID      uuid.UUID `json:"barber_shop_id"`
	ServiceCategoryID int16     `json:"service_category_id"`
}

func (q *Queries) CreateBarberShopServiceCategories(ctx context.Context, arg CreateBarberShopServiceCategoriesParams) (BarberShopServiceCategory, error) {
	row := q.db.QueryRow(ctx, createBarberShopServiceCategories, arg.BarberShopID, arg.ServiceCategoryID)
	var i BarberShopServiceCategory
	err := row.Scan(&i.ID, &i.BarberShopID, &i.ServiceCategoryID)
	return i, err
}

const deleteBarberShopServiceCategories = `-- name: DeleteBarberShopServiceCategories :exec
DELETE FROM "BarberShopServiceCategories"
WHERE
  "id" = $1
`

func (q *Queries) DeleteBarberShopServiceCategories(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteBarberShopServiceCategories, id)
	return err
}

const listBarberShopServiceCategories = `-- name: ListBarberShopServiceCategories :many
SELECT
  bsc.id, bsc.barber_shop_id, bsc.service_category_id,
  sc."name" as "service_category_name",
  sc."is_global" as "service_category_is_global"
FROM
  "BarberShopServiceCategories" bsc
JOIN
  "ServiceCategories" sc ON bsc."service_category_id" = sc."id"
WHERE
  bsc."barber_shop_id" = $1
`

type ListBarberShopServiceCategoriesRow struct {
	ID                      uuid.UUID `json:"id"`
	BarberShopID            uuid.UUID `json:"barber_shop_id"`
	ServiceCategoryID       int16     `json:"service_category_id"`
	ServiceCategoryName     string    `json:"service_category_name"`
	ServiceCategoryIsGlobal bool      `json:"service_category_is_global"`
}

func (q *Queries) ListBarberShopServiceCategories(ctx context.Context, barberShopID uuid.UUID) ([]ListBarberShopServiceCategoriesRow, error) {
	rows, err := q.db.Query(ctx, listBarberShopServiceCategories, barberShopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListBarberShopServiceCategoriesRow{}
	for rows.Next() {
		var i ListBarberShopServiceCategoriesRow
		if err := rows.Scan(
			&i.ID,
			&i.BarberShopID,
			&i.ServiceCategoryID,
			&i.ServiceCategoryName,
			&i.ServiceCategoryIsGlobal,
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

const updateBarberShopServiceCategories = `-- name: UpdateBarberShopServiceCategories :one
UPDATE "BarberShopServiceCategories"
SET
  "service_category_id" = $1
WHERE
  "id" = $2
RETURNING id, barber_shop_id, service_category_id
`

type UpdateBarberShopServiceCategoriesParams struct {
	ServiceCategoryID int16     `json:"service_category_id"`
	ID                uuid.UUID `json:"id"`
}

func (q *Queries) UpdateBarberShopServiceCategories(ctx context.Context, arg UpdateBarberShopServiceCategoriesParams) (BarberShopServiceCategory, error) {
	row := q.db.QueryRow(ctx, updateBarberShopServiceCategories, arg.ServiceCategoryID, arg.ID)
	var i BarberShopServiceCategory
	err := row.Scan(&i.ID, &i.BarberShopID, &i.ServiceCategoryID)
	return i, err
}
