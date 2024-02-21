// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: customers.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const changePasswordCustomer = `-- name: ChangePasswordCustomer :one
UPDATE "Customers" 
set   
  "hashed_password" = $1::varchar(100),
  "password_changed_at" = NOW()
WHERE "id" = $2
RETURNING id, name, email, phone, gender_id, hashed_password, avatar, is_social_auth, password_changed_at, create_at, update_at
`

type ChangePasswordCustomerParams struct {
	HashedPassword string    `json:"hashed_password"`
	ID             uuid.UUID `json:"id"`
}

func (q *Queries) ChangePasswordCustomer(ctx context.Context, arg ChangePasswordCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, changePasswordCustomer, arg.HashedPassword, arg.ID)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.GenderID,
		&i.HashedPassword,
		&i.Avatar,
		&i.IsSocialAuth,
		&i.PasswordChangedAt,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO "Customers" (
   "name",
    email,
    phone,
    hashed_password,
    is_social_auth
  )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5::bool
  )
RETURNING id, name, email, phone, gender_id, hashed_password, avatar, is_social_auth, password_changed_at, create_at, update_at
`

type CreateCustomerParams struct {
	Name           string         `json:"name"`
	Email          string         `json:"email"`
	Phone          sql.NullString `json:"phone"`
	HashedPassword sql.NullString `json:"hashed_password"`
	IsSocialAuth   bool           `json:"is_social_auth"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, createCustomer,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.HashedPassword,
		arg.IsSocialAuth,
	)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.GenderID,
		&i.HashedPassword,
		&i.Avatar,
		&i.IsSocialAuth,
		&i.PasswordChangedAt,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const getCustomer = `-- name: GetCustomer :one
SELECT id, name, email, phone, gender_id, hashed_password, avatar, is_social_auth, password_changed_at, create_at, update_at
  FROM "Customers"
  WHERE "id" = $1
`

func (q *Queries) GetCustomer(ctx context.Context, id uuid.UUID) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getCustomer, id)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.GenderID,
		&i.HashedPassword,
		&i.Avatar,
		&i.IsSocialAuth,
		&i.PasswordChangedAt,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const getUserCustomer = `-- name: GetUserCustomer :one
SELECT id, name, email, phone, gender_id, hashed_password, avatar, is_social_auth, password_changed_at, create_at, update_at
FROM "Customers"
WHERE
    (
        ($2::varchar = 'email' AND email = $1)
        OR
        ($2::varchar = 'phone' AND phone = $1)
    )
LIMIT 1
`

type GetUserCustomerParams struct {
	Email        string `json:"email"`
	TypeUsername string `json:"type_username"`
}

func (q *Queries) GetUserCustomer(ctx context.Context, arg GetUserCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, getUserCustomer, arg.Email, arg.TypeUsername)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.GenderID,
		&i.HashedPassword,
		&i.Avatar,
		&i.IsSocialAuth,
		&i.PasswordChangedAt,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}

const updateCustomer = `-- name: UpdateCustomer :one
UPDATE "Customers"
set name = coalesce($1, name),
  email = coalesce($2, email),
  phone = coalesce($3, phone),
  gender_id = coalesce($4, gender_id),
  avatar = coalesce($5, avatar),
  "update_at" = NOW()
  WHERE "id" = $6
RETURNING id, name, email, phone, gender_id, hashed_password, avatar, is_social_auth, password_changed_at, create_at, update_at
`

type UpdateCustomerParams struct {
	Name     sql.NullString `json:"name"`
	Email    sql.NullString `json:"email"`
	Phone    sql.NullString `json:"phone"`
	GenderID sql.NullInt32  `json:"gender_id"`
	Avatar   sql.NullString `json:"avatar"`
	ID       uuid.UUID      `json:"id"`
}

func (q *Queries) UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) (Customer, error) {
	row := q.db.QueryRowContext(ctx, updateCustomer,
		arg.Name,
		arg.Email,
		arg.Phone,
		arg.GenderID,
		arg.Avatar,
		arg.ID,
	)
	var i Customer
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Phone,
		&i.GenderID,
		&i.HashedPassword,
		&i.Avatar,
		&i.IsSocialAuth,
		&i.PasswordChangedAt,
		&i.CreateAt,
		&i.UpdateAt,
	)
	return i, err
}
