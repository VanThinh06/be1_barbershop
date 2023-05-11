// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package db

import (
	"context"

	null "gopkg.in/guregu/null.v4"
)

const createUsers = `-- name: CreateUsers :one

INSERT INTO users (username, full_name, email, hashed_password)
VALUES ($1,
        $2,
        $3,
        $4) RETURNING username, full_name, email, hashed_password, password_changed_at, created_at, image, role
`

type CreateUsersParams struct {
	Username       string      `json:"username"`
	FullName       null.String `json:"full_name"`
	Email          string      `json:"email"`
	HashedPassword string      `json:"hashed_password"`
}

func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUsers,
		arg.Username,
		arg.FullName,
		arg.Email,
		arg.HashedPassword,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.Image,
		&i.Role,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :one

SELECT username, full_name, email, hashed_password, password_changed_at, created_at, image, role
FROM users
WHERE username = $1
LIMIT 1
`

func (q *Queries) GetUsers(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUsers, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.Image,
		&i.Role,
	)
	return i, err
}
