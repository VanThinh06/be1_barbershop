// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: appointments.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const insertAppointmentAndGetInfo = `-- name: InsertAppointmentAndGetInfo :one
WITH inserted_appointment AS (
    INSERT INTO "Appointments" (
        barbershop_id,
        customer_id,
        barber_id,    
        appointment_datetime,
        "status"  
    )
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id, barbershop_id, customer_id, barber_id, service_id, appointment_datetime, status, create_at, update_at
)
SELECT 
    inserted_appointment.id, inserted_appointment.barbershop_id, inserted_appointment.customer_id, inserted_appointment.barber_id, inserted_appointment.service_id, inserted_appointment.appointment_datetime, inserted_appointment.status, inserted_appointment.create_at, inserted_appointment.update_at,
    "Barbers".nick_name AS barber_nick_name,
    "SessionsBarber".fcm_device,
    "BarberShops"."name" AS name_barber_shop
FROM 
    "inserted_appointment"
JOIN 
    "Barbers" ON "inserted_appointment".barber_id = "Barbers".barber_id
JOIN 
    "BarberShops" ON "inserted_appointment".barbershops_id = "BarberShops".shop_id
LEFT JOIN 
    "SessionsBarber" ON "Barbers".id = "SessionsBarber".barber_id
ORDER BY 
    "SessionsBarber".create_at DESC
LIMIT 1
`

type InsertAppointmentAndGetInfoParams struct {
	BarbershopID        uuid.UUID `json:"barbershop_id"`
	CustomerID          uuid.UUID `json:"customer_id"`
	BarberID            uuid.UUID `json:"barber_id"`
	AppointmentDatetime time.Time `json:"appointment_datetime"`
	Status              int32     `json:"status"`
}

type InsertAppointmentAndGetInfoRow struct {
	ID                  uuid.UUID      `json:"id"`
	BarbershopID        uuid.UUID      `json:"barbershop_id"`
	CustomerID          uuid.UUID      `json:"customer_id"`
	BarberID            uuid.UUID      `json:"barber_id"`
	ServiceID           uuid.UUID      `json:"service_id"`
	AppointmentDatetime time.Time      `json:"appointment_datetime"`
	Status              int32          `json:"status"`
	CreateAt            time.Time      `json:"create_at"`
	UpdateAt            time.Time      `json:"update_at"`
	BarberNickName      string         `json:"barber_nick_name"`
	FcmDevice           sql.NullString `json:"fcm_device"`
	NameBarberShop      string         `json:"name_barber_shop"`
}

func (q *Queries) InsertAppointmentAndGetInfo(ctx context.Context, arg InsertAppointmentAndGetInfoParams) (InsertAppointmentAndGetInfoRow, error) {
	row := q.db.QueryRowContext(ctx, insertAppointmentAndGetInfo,
		arg.BarbershopID,
		arg.CustomerID,
		arg.BarberID,
		arg.AppointmentDatetime,
		arg.Status,
	)
	var i InsertAppointmentAndGetInfoRow
	err := row.Scan(
		&i.ID,
		&i.BarbershopID,
		&i.CustomerID,
		&i.BarberID,
		&i.ServiceID,
		&i.AppointmentDatetime,
		&i.Status,
		&i.CreateAt,
		&i.UpdateAt,
		&i.BarberNickName,
		&i.FcmDevice,
		&i.NameBarberShop,
	)
	return i, err
}

const insertServicesForAppointment = `-- name: InsertServicesForAppointment :many
INSERT INTO "BarberShopServices_Appointments" ("BarberShopServices_id", "Appointments_service_id")
SELECT unnest($2::uuid[]), $1
RETURNING "BarberShopServices_id", "Appointments_service_id"
`

type InsertServicesForAppointmentParams struct {
	AppointmentsServiceID uuid.UUID   `json:"Appointments_service_id"`
	ServicesID            []uuid.UUID `json:"services_id"`
}

func (q *Queries) InsertServicesForAppointment(ctx context.Context, arg InsertServicesForAppointmentParams) ([]BarberShopServicesAppointment, error) {
	rows, err := q.db.QueryContext(ctx, insertServicesForAppointment, arg.AppointmentsServiceID, pq.Array(arg.ServicesID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []BarberShopServicesAppointment{}
	for rows.Next() {
		var i BarberShopServicesAppointment
		if err := rows.Scan(&i.BarberShopServicesID, &i.AppointmentsServiceID); err != nil {
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

const listAppointmentByDateWithService = `-- name: ListAppointmentByDateWithService :many
SELECT 
    "Appointments".id, "Appointments".barbershop_id, "Appointments".customer_id, "Appointments".barber_id, "Appointments".service_id, "Appointments".appointment_datetime, "Appointments".status, "Appointments".create_at, "Appointments".update_at,
    SUM("Services"."timer") AS "service_timer"
FROM "Appointments"
LEFT JOIN "BarberShopServices_Appointments" ON "Appointments"."id" = "BarberShopServices_Appointments"."Appointments_service_id"
LEFT JOIN "BarberShopServices" ON "BarberShopServices_Appointments"."Services_id" = "BarberShopServices"."id"
WHERE DATE("Appointments"."appointment_datetime") = $1
    AND "Appointments"."barber_id" = $2
GROUP BY "Appointments"."id", "Appointments"."appointment_datetime" 
ORDER BY "Appointments"."appointment_datetime"
`

type ListAppointmentByDateWithServiceParams struct {
	AppointmentDatetime time.Time `json:"appointment_datetime"`
	BarberID            uuid.UUID `json:"barber_id"`
}

type ListAppointmentByDateWithServiceRow struct {
	ID                  uuid.UUID `json:"id"`
	BarbershopID        uuid.UUID `json:"barbershop_id"`
	CustomerID          uuid.UUID `json:"customer_id"`
	BarberID            uuid.UUID `json:"barber_id"`
	ServiceID           uuid.UUID `json:"service_id"`
	AppointmentDatetime time.Time `json:"appointment_datetime"`
	Status              int32     `json:"status"`
	CreateAt            time.Time `json:"create_at"`
	UpdateAt            time.Time `json:"update_at"`
	ServiceTimer        int64     `json:"service_timer"`
}

func (q *Queries) ListAppointmentByDateWithService(ctx context.Context, arg ListAppointmentByDateWithServiceParams) ([]ListAppointmentByDateWithServiceRow, error) {
	rows, err := q.db.QueryContext(ctx, listAppointmentByDateWithService, arg.AppointmentDatetime, arg.BarberID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAppointmentByDateWithServiceRow{}
	for rows.Next() {
		var i ListAppointmentByDateWithServiceRow
		if err := rows.Scan(
			&i.ID,
			&i.BarbershopID,
			&i.CustomerID,
			&i.BarberID,
			&i.ServiceID,
			&i.AppointmentDatetime,
			&i.Status,
			&i.CreateAt,
			&i.UpdateAt,
			&i.ServiceTimer,
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
