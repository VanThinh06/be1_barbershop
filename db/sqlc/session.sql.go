// Code generated by sqlc. DO NOT EDIT.
// source: session.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createSessionBarber = `-- name: CreateSessionBarber :one

INSERT INTO "sessions_barber" (id,
                               username,
                               refresh_token,
                               user_agent,
                               is_manager,
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
        $8,
        $9 ) RETURNING id, username, refresh_token, is_manager, user_agent, client_ip, fcm_device, is_blocked, expires_at, create_at
`

type CreateSessionBarberParams struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	IsManager    bool      `json:"is_manager"`
	ClientIp     string    `json:"client_ip"`
	FcmDevice    string    `json:"fcm_device"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (q *Queries) CreateSessionBarber(ctx context.Context, arg CreateSessionBarberParams) (SessionsBarber, error) {
	row := q.db.QueryRowContext(ctx, createSessionBarber,
		arg.ID,
		arg.Username,
		arg.RefreshToken,
		arg.UserAgent,
		arg.IsManager,
		arg.ClientIp,
		arg.FcmDevice,
		arg.IsBlocked,
		arg.ExpiresAt,
	)
	var i SessionsBarber
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.IsManager,
		&i.UserAgent,
		&i.ClientIp,
		&i.FcmDevice,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreateAt,
	)
	return i, err
}

const getSession = `-- name: GetSession :one

SELECT id, username, refresh_token, is_manager, user_agent, client_ip, fcm_device, is_blocked, expires_at, create_at
FROM "sessions_barber"
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetSession(ctx context.Context, id uuid.UUID) (SessionsBarber, error) {
	row := q.db.QueryRowContext(ctx, getSession, id)
	var i SessionsBarber
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.IsManager,
		&i.UserAgent,
		&i.ClientIp,
		&i.FcmDevice,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreateAt,
	)
	return i, err
}
