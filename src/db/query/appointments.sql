-- name: InsertAppointment :one
INSERT INTO "Appointments" (
    barbershops_id,
    customer_id,
    barber_id,    
    appointment_datetime,
    "status"  
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: InsertServicesForAppointment :many
WITH inserted_services AS (
  INSERT INTO "Services_Appointments" ("Services_id", "Appointments_service_id")
  SELECT unnest(sqlc.arg(Services_id)::uuid[]), $1
  RETURNING "Services_id", "Appointments_service_id"
)
SELECT "Services_id", "Appointments_service_id", s.*
FROM inserted_services
JOIN "Services" s ON "Services_id" = s."id";

-- name: GetAppointmentByDate :many
SELECT * FROM "Appointments"
WHERE DATE(appointment_datetime) = $1
AND barber_id = $2
ORDER BY appointment_datetime;