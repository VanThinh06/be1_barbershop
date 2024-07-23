// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: otp_request.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const getCountOTPsForBarberToday = `-- name: GetCountOTPsForBarberToday :one
SELECT COUNT(*)
FROM "OTPRequests"
WHERE barber_id = $1
AND date_trunc('day', requested_at) = date_trunc('day', now())
`

func (q *Queries) GetCountOTPsForBarberToday(ctx context.Context, barberID uuid.UUID) (int64, error) {
	row := q.db.QueryRow(ctx, getCountOTPsForBarberToday, barberID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getOTPRequestDetails = `-- name: GetOTPRequestDetails :one
SELECT id, otp, requested_at, is_confirm
FROM "OTPRequests" 
WHERE barber_id = $1 AND otp = $2 
ORDER BY requested_at DESC 
LIMIT 1
`

type GetOTPRequestDetailsParams struct {
	BarberID uuid.UUID `json:"barber_id"`
	Otp      string    `json:"otp"`
}

type GetOTPRequestDetailsRow struct {
	ID          uuid.UUID `json:"id"`
	Otp         string    `json:"otp"`
	RequestedAt time.Time `json:"requested_at"`
	IsConfirm   bool      `json:"is_confirm"`
}

func (q *Queries) GetOTPRequestDetails(ctx context.Context, arg GetOTPRequestDetailsParams) (GetOTPRequestDetailsRow, error) {
	row := q.db.QueryRow(ctx, getOTPRequestDetails, arg.BarberID, arg.Otp)
	var i GetOTPRequestDetailsRow
	err := row.Scan(
		&i.ID,
		&i.Otp,
		&i.RequestedAt,
		&i.IsConfirm,
	)
	return i, err
}

const insertNewOTPRequest = `-- name: InsertNewOTPRequest :exec
INSERT INTO "OTPRequests" (
    barber_id,
    otp) 
VALUES ($1, $2)
`

type InsertNewOTPRequestParams struct {
	BarberID uuid.UUID `json:"barber_id"`
	Otp      string    `json:"otp"`
}

func (q *Queries) InsertNewOTPRequest(ctx context.Context, arg InsertNewOTPRequestParams) error {
	_, err := q.db.Exec(ctx, insertNewOTPRequest, arg.BarberID, arg.Otp)
	return err
}

const updateOTPRequestConfirmation = `-- name: UpdateOTPRequestConfirmation :exec
UPDATE "OTPRequests"
SET is_confirm = $1
WHERE id = $2
`

type UpdateOTPRequestConfirmationParams struct {
	IsConfirm bool      `json:"is_confirm"`
	ID        uuid.UUID `json:"id"`
}

func (q *Queries) UpdateOTPRequestConfirmation(ctx context.Context, arg UpdateOTPRequestConfirmationParams) error {
	_, err := q.db.Exec(ctx, updateOTPRequestConfirmation, arg.IsConfirm, arg.ID)
	return err
}
