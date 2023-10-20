-- name: CreateBarber :one
INSERT INTO "Barbers" (
    shop_id,
    nick_name,
    full_name,
    phone,
    email,
    gender,
    "role",
    hashed_password,
    avatar
  )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
  )
RETURNING *;
-- name: GetEmailBarber :one
SELECT *
FROM "Barbers"
WHERE email = $1
LIMIT 1;
-- name: UpdateBarber :one
UPDATE "Barbers"
set shop_id = $1,
  nick_name = $2,
  full_name = $3,
  phone = $4,
  email = $5,
  gender = $6,
  "role" = $7,
  avatar = $8,
  "status" = $9,
  "update_at" = $10
WHERE barber_id = $11
RETURNING *;
-- name: UpdateIDShopManager :one
UPDATE "Barbers"
set shop_id = $1
WHERE barber_id = $2
RETURNING *;
-- name: UpdateStatusBarber :one
UPDATE "Barbers"
set "status" = $1
WHERE barber_id = $2
RETURNING *;