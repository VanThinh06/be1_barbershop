// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: barber_roles.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createBarberRoles = `-- name: CreateBarberRoles :one
INSERT INTO "BarberRoles" (
    barber_id,
    barber_shop_id,
    role_id
  )
VALUES (
    $1,
    $2,
    $3
  )
RETURNING id, barber_id, barber_shop_id, role_id
`

type CreateBarberRolesParams struct {
	BarberID     uuid.UUID `json:"barber_id"`
	BarberShopID uuid.UUID `json:"barber_shop_id"`
	RoleID       int16     `json:"role_id"`
}

func (q *Queries) CreateBarberRoles(ctx context.Context, arg CreateBarberRolesParams) (BarberRole, error) {
	row := q.db.QueryRow(ctx, createBarberRoles, arg.BarberID, arg.BarberShopID, arg.RoleID)
	var i BarberRole
	err := row.Scan(
		&i.ID,
		&i.BarberID,
		&i.BarberShopID,
		&i.RoleID,
	)
	return i, err
}

const deleteBarberRole = `-- name: DeleteBarberRole :exec
DELETE FROM "BarberRoles"
WHERE "id" = $1
`

func (q *Queries) DeleteBarberRole(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteBarberRole, id)
	return err
}

const getBarberRoles = `-- name: GetBarberRoles :one
SELECT id, barber_id, barber_shop_id, role_id
FROM "BarberRoles"
WHERE "BarberRoles"."barber_id" = $1 AND "BarberRoles"."barber_shop_id" = $2
`

type GetBarberRolesParams struct {
	BarberID     uuid.UUID `json:"barber_id"`
	BarberShopID uuid.UUID `json:"barber_shop_id"`
}

func (q *Queries) GetBarberRoles(ctx context.Context, arg GetBarberRolesParams) (BarberRole, error) {
	row := q.db.QueryRow(ctx, getBarberRoles, arg.BarberID, arg.BarberShopID)
	var i BarberRole
	err := row.Scan(
		&i.ID,
		&i.BarberID,
		&i.BarberShopID,
		&i.RoleID,
	)
	return i, err
}

const listBarbersRoles = `-- name: ListBarbersRoles :many
SELECT "BarberRoles".id, barber_id, barber_shop_id, role_id, "Roles".id, name, type
FROM "BarberRoles"
JOIN "Roles" ON "BarberRoles"."role_id" = "Roles"."id"
WHERE "BarberRoles"."barber_shop_id" = $1
ORDER BY "Roles"."id"
`

type ListBarbersRolesRow struct {
	ID           uuid.UUID      `json:"id"`
	BarberID     uuid.UUID      `json:"barber_id"`
	BarberShopID uuid.UUID      `json:"barber_shop_id"`
	RoleID       int16          `json:"role_id"`
	ID_2         int16          `json:"id_2"`
	Name         string         `json:"name"`
	Type         sql.NullString `json:"type"`
}

func (q *Queries) ListBarbersRoles(ctx context.Context, barberShopID uuid.UUID) ([]ListBarbersRolesRow, error) {
	rows, err := q.db.Query(ctx, listBarbersRoles, barberShopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListBarbersRolesRow{}
	for rows.Next() {
		var i ListBarbersRolesRow
		if err := rows.Scan(
			&i.ID,
			&i.BarberID,
			&i.BarberShopID,
			&i.RoleID,
			&i.ID_2,
			&i.Name,
			&i.Type,
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

const updateBarberRoles = `-- name: UpdateBarberRoles :one
UPDATE "BarberRoles"
SET "role_id" = $1,
    "update_at" = NOW()
WHERE "id" = $2
RETURNING id, barber_id, barber_shop_id, role_id
`

type UpdateBarberRolesParams struct {
	RoleID int16     `json:"role_id"`
	ID     uuid.UUID `json:"id"`
}

func (q *Queries) UpdateBarberRoles(ctx context.Context, arg UpdateBarberRolesParams) (BarberRole, error) {
	row := q.db.QueryRow(ctx, updateBarberRoles, arg.RoleID, arg.ID)
	var i BarberRole
	err := row.Scan(
		&i.ID,
		&i.BarberID,
		&i.BarberShopID,
		&i.RoleID,
	)
	return i, err
}
