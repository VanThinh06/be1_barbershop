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
    avatar,
    manager_id,
    "haircut"
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
    $9,
    $10,
    $11
  )
RETURNING *;

-- name: GetEmailBarber :one
SELECT *
FROM "Barbers"
WHERE email = $1
LIMIT 1;

-- name: UpdateBarber :one
UPDATE "Barbers"
set shop_id = coalesce(sqlc.narg('shop_id'), shop_id),
  nick_name = coalesce(sqlc.narg('nick_name'), nick_name),
  full_name = coalesce(sqlc.narg('full_name'), full_name),
  phone = coalesce(sqlc.narg('phone'), phone),
  email = coalesce(sqlc.narg('email'), email),
  gender = coalesce(sqlc.narg('gender'), gender),
  avatar = coalesce(sqlc.narg('avatar'), avatar),
  "status" = coalesce(sqlc.narg('status'), status),
  "hashed_password" = coalesce(sqlc.narg('hashed_password'), hashed_password),
  "password_changed_at" = coalesce(
    sqlc.narg('password_changed_at'),
    password_changed_at
  ),
  "update_at" = sqlc.arg('update_at'),
  "haircut" = coalesce(sqlc.narg('haircut'), haircut)
  WHERE "barber_id" = sqlc.arg('barber_id')
RETURNING *;

-- name: BarberGetIdShop :one
SELECT shop_id::varchar, "role"
FROM "Barbers"
WHERE barber_id = $1
LIMIT 1;

-- name: GetBarberInBarberShop :many
SELECT *
FROM "Barbers"
WHERE shop_id =sqlc.arg('shop_id') AND "haircut" = true;
