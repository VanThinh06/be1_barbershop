-- name: CreateAppointment :one
INSERT INTO "Appointments" (
    customer_id,
    barber_id,    
    service_id,    
    appointment_datetime,
    "status"  
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;