// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: barbershop_service_categories.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createBarberShopServiceCategoryWithChain = `-- name: CreateBarberShopServiceCategoryWithChain :one
INSERT INTO "BarberShopServiceCategories" (
  "barbershop_chain_id",
  "service_category_id"
)
VALUES (
  $1,
  $2
)
RETURNING id, barbershop_chain_id, barbershop_id, service_category_id, create_at, update_at
`

type CreateBarberShopServiceCategoryWithChainParams struct {
	BarbershopChainID uuid.NullUUID `json:"barbershop_chain_id"`
	ServiceCategoryID uuid.NullUUID `json:"service_category_id"`
}

func (q *Queries) CreateBarberShopServiceCategoryWithChain(ctx context.Context, arg CreateBarberShopServiceCategoryWithChainParams) (BarberShopServiceCategory, error) {
	row := q.db.QueryRowContext(ctx, createBarberShopServiceCategoryWithChain, arg.BarbershopChainID, arg.ServiceCategoryID)
	var i BarberShopServiceCategory
	err := row.Scan(
		&i.ID,
		&i.BarbershopChainID,
		&i.BarbershopID,
		&i.ServiceCategoryID,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const createBarberShopServiceCategoryWithShop = `-- name: CreateBarberShopServiceCategoryWithShop :one
INSERT INTO "BarberShopServiceCategories" (
  "barbershop_id",
  "service_category_id"
)
VALUES (
  $1,
  $2
)
RETURNING id, barbershop_chain_id, barbershop_id, service_category_id, create_at, update_at
`

type CreateBarberShopServiceCategoryWithShopParams struct {
	BarbershopID      uuid.NullUUID `json:"barbershop_id"`
	ServiceCategoryID uuid.NullUUID `json:"service_category_id"`
}

func (q *Queries) CreateBarberShopServiceCategoryWithShop(ctx context.Context, arg CreateBarberShopServiceCategoryWithShopParams) (BarberShopServiceCategory, error) {
	row := q.db.QueryRowContext(ctx, createBarberShopServiceCategoryWithShop, arg.BarbershopID, arg.ServiceCategoryID)
	var i BarberShopServiceCategory
	err := row.Scan(
		&i.ID,
		&i.BarbershopChainID,
		&i.BarbershopID,
		&i.ServiceCategoryID,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteBarberShopServiceCategory = `-- name: DeleteBarberShopServiceCategory :exec
DELETE FROM "BarberShopServiceCategories"
WHERE
  "id" = $1
`

func (q *Queries) DeleteBarberShopServiceCategory(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteBarberShopServiceCategory, id)
	return err
}

const listBarberShopServiceCategories = `-- name: ListBarberShopServiceCategories :many
SELECT
  bsc.id, bsc.barbershop_chain_id, bsc.barbershop_id, bsc.service_category_id, bsc.create_at, bsc.update_at,
  sc."name" as "service_category_name",
  sc."is_global" as "service_category_is_global"
FROM
  "BarberShopServiceCategories" bsc
JOIN
  "ServiceCategories" sc ON bsc."service_category_id" = sc."id"
WHERE
  bsc."barbershop_id" = $1
  AND bsc."barbershop_chain_id" = $2
`

type ListBarberShopServiceCategoriesParams struct {
	BarbershopID      uuid.NullUUID `json:"barbershop_id"`
	BarbershopChainID uuid.NullUUID `json:"barbershop_chain_id"`
}

type ListBarberShopServiceCategoriesRow struct {
	ID                      uuid.UUID     `json:"id"`
	BarbershopChainID       uuid.NullUUID `json:"barbershop_chain_id"`
	BarbershopID            uuid.NullUUID `json:"barbershop_id"`
	ServiceCategoryID       uuid.NullUUID `json:"service_category_id"`
	CreateAt                time.Time     `json:"create_at"`
	UpdateAt                time.Time     `json:"update_at"`
	ServiceCategoryName     string        `json:"service_category_name"`
	ServiceCategoryIsGlobal bool          `json:"service_category_is_global"`
}

func (q *Queries) ListBarberShopServiceCategories(ctx context.Context, arg ListBarberShopServiceCategoriesParams) ([]ListBarberShopServiceCategoriesRow, error) {
	rows, err := q.db.QueryContext(ctx, listBarberShopServiceCategories, arg.BarbershopID, arg.BarbershopChainID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListBarberShopServiceCategoriesRow{}
	for rows.Next() {
		var i ListBarberShopServiceCategoriesRow
		if err := rows.Scan(
			&i.ID,
			&i.BarbershopChainID,
			&i.BarbershopID,
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

const updateBarberShopServiceCategory = `-- name: UpdateBarberShopServiceCategory :exec
UPDATE "BarberShopServiceCategories"
SET
  "service_category_id" = $1
WHERE
  "id" = $2
`

type UpdateBarberShopServiceCategoryParams struct {
	ServiceCategoryID uuid.NullUUID `json:"service_category_id"`
	ID                uuid.UUID     `json:"id"`
}

func (q *Queries) UpdateBarberShopServiceCategory(ctx context.Context, arg UpdateBarberShopServiceCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateBarberShopServiceCategory, arg.ServiceCategoryID, arg.ID)
	return err
}
