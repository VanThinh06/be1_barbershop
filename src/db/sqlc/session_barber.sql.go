// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: session_barber.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSessionBarber = `-- name: CreateSessionBarber :one
INSERT INTO "SessionsBarber" (
                               id,
                               barber_id,
                               refresh_token,
                               user_agent,
                               client_ip,
                               fcm_device,
                               is_blocked,
                               expires_at)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8
        ) RETURNING id, barber_id, refresh_token, user_agent, client_ip, fcm_device, is_blocked, expires_at, create_at
`

type CreateSessionBarberParams struct {
	ID           uuid.UUID `json:"id"`
	BarberID     uuid.UUID `json:"barber_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	FcmDevice    string    `json:"fcm_device"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (q *Queries) CreateSessionBarber(ctx context.Context, arg CreateSessionBarberParams) (SessionsBarber, error) {
	row := q.db.QueryRowContext(ctx, createSessionBarber,
		arg.ID,
		arg.BarberID,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.FcmDevice,
		arg.IsBlocked,
		arg.ExpiresAt,
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

const getSessionsBarber = `-- name: GetSessionsBarber :one
SELECT id, barber_id, refresh_token, user_agent, client_ip, fcm_device, is_blocked, expires_at, create_at
FROM "SessionsBarber"
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetSessionsBarber(ctx context.Context, id uuid.UUID) (SessionsBarber, error) {
	row := q.db.QueryRowContext(ctx, getSessionsBarber, id)
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

const updateRefreshTokenSessionsBarber = `-- name: UpdateRefreshTokenSessionsBarber :one
UPDATE "SessionsBarber"
set refresh_token = $2,
    expires_at = $3
WHERE id = $1
RETURNING id, barber_id, refresh_token, user_agent, client_ip, fcm_device, is_blocked, expires_at, create_at
`

type UpdateRefreshTokenSessionsBarberParams struct {
	ID           uuid.UUID `json:"id"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (q *Queries) UpdateRefreshTokenSessionsBarber(ctx context.Context, arg UpdateRefreshTokenSessionsBarberParams) (SessionsBarber, error) {
	row := q.db.QueryRowContext(ctx, updateRefreshTokenSessionsBarber, arg.ID, arg.RefreshToken, arg.ExpiresAt)
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
