// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: barber_roles.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createBarberRole = `-- name: CreateBarberRole :one
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

type CreateBarberRoleParams struct {
	BarberID     uuid.UUID `json:"barber_id"`
	BarberShopID uuid.UUID `json:"barber_shop_id"`
	RoleID       int16     `json:"role_id"`
}

func (q *Queries) CreateBarberRole(ctx context.Context, arg CreateBarberRoleParams) (BarberRole, error) {
	row := q.db.QueryRow(ctx, createBarberRole, arg.BarberID, arg.BarberShopID, arg.RoleID)
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

const getBarberRole = `-- name: GetBarberRole :one
SELECT id, barber_id, barber_shop_id, role_id
FROM "BarberRoles"
WHERE "BarberRoles"."barber_id" = $1 AND "BarberRoles"."barber_shop_id" = $2
`

type GetBarberRoleParams struct {
	BarberID     uuid.UUID `json:"barber_id"`
	BarberShopID uuid.UUID `json:"barber_shop_id"`
}

func (q *Queries) GetBarberRole(ctx context.Context, arg GetBarberRoleParams) (BarberRole, error) {
	row := q.db.QueryRow(ctx, getBarberRole, arg.BarberID, arg.BarberShopID)
	var i BarberRole
	err := row.Scan(
		&i.ID,
		&i.BarberID,
		&i.BarberShopID,
		&i.RoleID,
	)
	return i, err
}
