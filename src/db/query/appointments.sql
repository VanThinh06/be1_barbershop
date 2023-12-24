-- name: InsertAppointmentAndGetInfo :one
WITH inserted_appointment AS (
    INSERT INTO "Appointments" (
        barbershops_id,
        customer_id,
        barber_id,    
        appointment_datetime,
        "status"  
    )
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT 
    "inserted_appointment".*,
    "Barbers".nick_name AS barber_nick_name,
    "SessionsBarbers".fcm_device,
    "BarberShops"."name" AS name_barber_shop
FROM 
    "inserted_appointment"
JOIN 
    "Barbers" ON "inserted_appointment".barber_id = "Barbers".barber_id
JOIN 
    "BarberShops" ON "inserted_appointment".barbershops_id = "BarberShops".shop_id
LEFT JOIN 
    "SessionsBarbers" ON "Barbers".barber_id = "SessionsBarbers".barber_id
ORDER BY 
    "SessionsBarbers".created_at DESC
LIMIT 1;

-- name: InsertServicesForAppointment :many
WITH inserted_services AS (
  INSERT INTO "Services_Appointments" ("Services_id", "Appointments_service_id")
  SELECT unnest(sqlc.arg(Services_id)::uuid[]), $1
  RETURNING "Services_id", "Appointments_service_id"
)
SELECT "Services_id", "Appointments_service_id", s.*
FROM inserted_services
JOIN "Services" s ON "Services_id" = s."id";

-- name: GetAppointmentByDateWithService :many
SELECT 
    "Appointments".*,
    SUM("Services"."timer") AS "service_timer"
FROM "Appointments"
LEFT JOIN "Services_Appointments" ON "Appointments"."appointment_id" = "Services_Appointments"."Appointments_service_id"
LEFT JOIN "Services" ON "Services_Appointments"."Services_id" = "Services"."id"
WHERE DATE("Appointments"."appointment_datetime") = $1
    AND "Appointments"."barber_id" = $2
GROUP BY "Appointments"."appointment_id", "Appointments"."appointment_datetime" 
ORDER BY "Appointments"."appointment_datetime";
