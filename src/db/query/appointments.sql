-- name: CreateAppointments :one
WITH inserted_appointment AS (
    INSERT INTO "Appointments" (
        barber_shop_id,
        customer_id,
        barber_id,    
        appointment_date_time,
        "status"  
    )
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT 
    "inserted_appointment".*,
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
LIMIT 1;

-- name: CreateServicesForAppointments :many
INSERT INTO "BarberShopServices_Appointments" ("BarberShopServices_id", "AppointmentsService_id")
SELECT unnest(sqlc.arg(services_id)::uuid[]), $1
RETURNING "BarberShopServices_id", "AppointmentsService_id";

-- name: ListAppointmentsByDate :many
SELECT 
    "Appointments".*,
    SUM("Services"."timer") AS "service_timer"
FROM "Appointments"
LEFT JOIN "BarberShopServices_Appointments" ON "Appointments"."id" = "BarberShopServices_Appointments"."Appointments_service_id"
LEFT JOIN "BarberShopServices" ON "BarberShopServices_Appointments"."Services_id" = "BarberShopServices"."id"
WHERE DATE("Appointments"."appointment_date_time") = $1
    AND "Appointments"."barber_id" = $2
GROUP BY "Appointments"."id", "Appointments"."appointment_date_time" 
ORDER BY "Appointments"."appointment_date_time";


-- TRIGGER 
-- CREATE OR REPLACE FUNCTION check_appointment_conflict()
-- RETURNS TRIGGER AS $$
-- BEGIN
--   IF EXISTS (
--     SELECT 1
--     FROM "Appointments"
--     WHERE "barbershops_id" = NEW."barbershops_id"
--       AND "barber_id" = NEW."barber_id"
--       AND (
--         (NEW."appointment_datetime" + NEW."timer" * interval '1 minute') BETWEEN "appointment_datetime" AND "appointment_datetime" + NEW."timer" * interval '1 minute'
--         OR NEW."appointment_datetime" BETWEEN "appointment_datetime" AND "appointment_datetime" + NEW."timer" * interval '1 minute'
--       )
--   ) THEN
--     RAISE EXCEPTION 'Appointment conflict: Another appointment exists in this time slot.';
--   END IF;
--   RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- CREATE TRIGGER check_appointment_conflict_trigger
-- BEFORE INSERT ON "Appointments"
-- FOR EACH ROW EXECUTE FUNCTION check_appointment_conflict();