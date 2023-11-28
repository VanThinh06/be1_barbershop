// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: appointments.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createAppointment = `-- name: CreateAppointment :one
INSERT INTO "Appointments" (
    customer_id,
    barber_id,    
    service_id,    
    appointment_datetime,
    "status"  
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING appointment_id, customer_id, barber_id, service_id, appointment_datetime, status, created_at, update_at
`

type CreateAppointmentParams struct {
	CustomerID          uuid.UUID `json:"customer_id"`
	BarberID            uuid.UUID `json:"barber_id"`
	ServiceID           uuid.UUID `json:"service_id"`
	AppointmentDatetime time.Time `json:"appointment_datetime"`
	Status              int32     `json:"status"`
}

func (q *Queries) CreateAppointment(ctx context.Context, arg CreateAppointmentParams) (Appointment, error) {
	row := q.db.QueryRowContext(ctx, createAppointment,
		arg.CustomerID,
		arg.BarberID,
		arg.ServiceID,
		arg.AppointmentDatetime,
		arg.Status,
	)
	var i Appointment
	err := row.Scan(
		&i.AppointmentID,
		&i.CustomerID,
		&i.BarberID,
		&i.ServiceID,
		&i.AppointmentDatetime,
		&i.Status,
		&i.CreatedAt,
		&i.UpdateAt,
	)
	return i, err
}