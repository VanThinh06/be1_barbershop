// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: barber_managers.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createBarberManagers = `-- name: CreateBarberManagers :one
INSERT INTO "BarberManagers" (
    barber_id,
    manager_id
  )
VALUES (
    $1,
    $2
  )
RETURNING barber_id, manager_id, create_at, update_at
`

type CreateBarberManagersParams struct {
	BarberID  uuid.UUID `json:"barber_id"`
	ManagerID uuid.UUID `json:"manager_id"`
}

func (q *Queries) CreateBarberManagers(ctx context.Context, arg CreateBarberManagersParams) (BarberManager, error) {
	row := q.db.QueryRowContext(ctx, createBarberManagers, arg.BarberID, arg.ManagerID)
	var i BarberManager
	err := row.Scan(
		&i.BarberID,
		&i.ManagerID,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteBarberManagers = `-- name: DeleteBarberManagers :exec
DELETE FROM "BarberManagers"
WHERE "barber_id" = $1 AND "manager_id" = $2
`

type DeleteBarberManagersParams struct {
	BarberID  uuid.UUID `json:"barber_id"`
	ManagerID uuid.UUID `json:"manager_id"`
}

func (q *Queries) DeleteBarberManagers(ctx context.Context, arg DeleteBarberManagersParams) error {
	_, err := q.db.ExecContext(ctx, deleteBarberManagers, arg.BarberID, arg.ManagerID)
	return err
}

const getBarberManagers = `-- name: GetBarberManagers :one
SELECT barber_id, manager_id, create_at, update_at
FROM "BarberManagers"
WHERE "barber_id" = $1
`

func (q *Queries) GetBarberManagers(ctx context.Context, barberID uuid.UUID) (BarberManager, error) {
	row := q.db.QueryRowContext(ctx, getBarberManagers, barberID)
	var i BarberManager
	err := row.Scan(
		&i.BarberID,
		&i.ManagerID,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}
