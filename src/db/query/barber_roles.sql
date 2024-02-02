-- name: CreateBarberRoles :one
INSERT INTO "BarberRoles" (
    barber_id,
    barber_shop_id,
    role_id
  )
VALUES (
    $1,
    $2,
    $3
  )
RETURNING *;

-- name: GetBarberRoles :one
SELECT *
FROM "BarberRoles"
WHERE "BarberRoles"."barber_id" = $1 AND "BarberRoles"."barber_shop_id" = $2;

-- name: ListBarbersRoles :many
SELECT *
FROM "BarberRoles"
JOIN "Roles" ON "BarberRoles"."role_id" = "Roles"."id"
WHERE "BarberRoles"."barber_shop_id" = $1
ORDER BY "Roles"."id";


-- name: UpdateBarberRoles :one
UPDATE "BarberRoles"
SET "role_id" = $1,
    "update_at" = NOW()
WHERE "id" = $2
RETURNING *;

-- name: DeleteBarberRoles :exec
DELETE FROM "BarberRoles"
WHERE "id" = $1;