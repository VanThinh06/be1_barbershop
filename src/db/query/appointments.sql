-- name: CreateAppointment :one
INSERT INTO "Appointments" (
    barbershops_id,
    customer_id,
    barber_id,    
    service_id,    
    appointment_datetime,
    "status"  
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAppointmentBarber :one
SELECT * FROM "Appointments"
WHERE DATE(appointment_datetime) = $1
AND employee_id = $1
ORDER BY appointment_datetime;