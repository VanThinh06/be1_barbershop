// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: sessions_barber.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createSessionBarber = `-- name: CreateSessionBarber :one
INSERT INTO "SessionsBarber" (
  barber_id,
  refresh_token,
  user_agent,
  client_ip,
  fcm_device,
  is_blocked,
  expires_at
)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  now() + interval '30 minutes' 
)
RETURNING id, barber_id, refresh_token, user_agent, client_ip, fcm_device, is_blocked, expires_at, create_at
`

type CreateSessionBarberParams struct {
	BarberID     uuid.UUID `json:"barber_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	FcmDevice    string    `json:"fcm_device"`
	IsBlocked    bool      `json:"is_blocked"`
}

func (q *Queries) CreateSessionBarber(ctx context.Context, arg CreateSessionBarberParams) (SessionsBarber, error) {
	row := q.db.QueryRowContext(ctx, createSessionBarber,
		arg.BarberID,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.FcmDevice,
		arg.IsBlocked,
	)
	var i SessionsBarber
	err := row.Scan(
		&i.ID,
		&i.BarberID,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.FcmDevice,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreateAt,
	)
	return i, err
}

const getSessionBarber = `-- name: GetSessionBarber :one
SELECT id, barber_id, refresh_token, user_agent, client_ip, fcm_device, is_blocked, expires_at, create_at
FROM "SessionsBarber"
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetSessionBarber(ctx context.Context, id uuid.UUID) (SessionsBarber, error) {
	row := q.db.QueryRowContext(ctx, getSessionBarber, id)
	var i SessionsBarber
	err := row.Scan(
		&i.ID,
		&i.BarberID,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.FcmDevice,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreateAt,
	)
	return i, err
}

const updateSessionRefreshToken = `-- name: UpdateSessionRefreshToken :exec
UPDATE "SessionsBarber"
SET refresh_token = $2,
 "expires_at" = now() + interval '30 minutes'
WHERE id = $1
`

type UpdateSessionRefreshTokenParams struct {
	ID           uuid.UUID `json:"id"`
	RefreshToken string    `json:"refresh_token"`
}

func (q *Queries) UpdateSessionRefreshToken(ctx context.Context, arg UpdateSessionRefreshTokenParams) error {
	_, err := q.db.ExecContext(ctx, updateSessionRefreshToken, arg.ID, arg.RefreshToken)
	return err
}
