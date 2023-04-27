// Code generated by sqlc. DO NOT EDIT.
// source: employee.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createEmployee = `-- name: CreateEmployee :one
INSERT INTO employee ( username, role, image, store_id)
VALUES ( $1,
         $2,
         $3,
         $4) RETURNING id, username, role, image, store_id, created_at, update_at
`

type CreateEmployeeParams struct {
	Username string         `json:"username"`
	Role     string         `json:"role"`
	Image    sql.NullString `json:"image"`
	StoreID  uuid.UUID      `json:"store_id"`
}

func (q *Queries) CreateEmployee(ctx context.Context, arg CreateEmployeeParams) (Employee, error) {
	row := q.db.QueryRowContext(ctx, createEmployee,
		arg.Username,
		arg.Role,
		arg.Image,
		arg.StoreID,
	)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Role,
		&i.Image,
		&i.StoreID,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const deleteEmployee = `-- name: DeleteEmployee :exec
DELETE
FROM employee
WHERE id = $1
`

func (q *Queries) DeleteEmployee(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteEmployee, id)
	return err
}

const getEmployee = `-- name: GetEmployee :one
SELECT id, username, role, image, store_id, created_at, update_at
FROM employee
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetEmployee(ctx context.Context, id uuid.UUID) (Employee, error) {
	row := q.db.QueryRowContext(ctx, getEmployee, id)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Role,
		&i.Image,
		&i.StoreID,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}

const getListEmployeewithStore = `-- name: GetListEmployeewithStore :many
SELECT id, username, role, image, store_id, created_at, update_at
FROM employee
WHERE store_id = $1
LIMIT $2
`

type GetListEmployeewithStoreParams struct {
	StoreID uuid.UUID `json:"store_id"`
	Limit   int32     `json:"limit"`
}

func (q *Queries) GetListEmployeewithStore(ctx context.Context, arg GetListEmployeewithStoreParams) ([]Employee, error) {
	rows, err := q.db.QueryContext(ctx, getListEmployeewithStore, arg.StoreID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Employee{}
	for rows.Next() {
		var i Employee
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Role,
			&i.Image,
			&i.StoreID,
			&i.CreatedAt,
			&i.UpdateAt,
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

const updateAuthor = `-- name: UpdateAuthor :one
UPDATE employee
set username = $2,
    image = $3,
    store_id = $4
WHERE id = $1 RETURNING id, username, role, image, store_id, created_at, update_at
`

type UpdateAuthorParams struct {
	ID       uuid.UUID      `json:"id"`
	Username string         `json:"username"`
	Image    sql.NullString `json:"image"`
	StoreID  uuid.UUID      `json:"store_id"`
}

func (q *Queries) UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) (Employee, error) {
	row := q.db.QueryRowContext(ctx, updateAuthor,
		arg.ID,
		arg.Username,
		arg.Image,
		arg.StoreID,
	)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Role,
		&i.Image,
		&i.StoreID,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}
