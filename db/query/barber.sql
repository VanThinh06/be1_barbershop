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
  WHERE barber_id = sqlc.arg('barber_id')
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