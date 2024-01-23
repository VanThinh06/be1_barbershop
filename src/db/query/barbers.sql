-- name: CreateBarber :one
INSERT INTO "Barbers" (
    gender_id,
    email,
    phone,
    hashed_password,
    nick_name,
    full_name,
    avatar_url
  )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
  )
RETURNING *;

-- name: GetBarber :one
SELECT
  b.*, 
  bs."name" as "shop_name",
  bs."address" as "shop_address",
  bs."coordinates" as "shop_coordinates",
  bs."start_time" as "shop_start_time",
  bs."end_time" as "shop_end_time",
  bs."break_time" as "shop_break_time",
  bs."break_minutes" as "shop_break_minutes",
  bs."interval_scheduler" as "shop_interval_scheduler",
  bs."reputation" as "shop_reputation",
  bs."rate" as "shop_rate",
  bs."branch_count" as "shop_branch_count",
  br."role_id" as "barber_role_id",
  br."barbershop_id" as "barber_role_barbershop_id"
FROM
  "Barbers" b
JOIN
  "BarberRoles" br ON b."id" = br."barber_id"
JOIN
  "BarberShops" bs ON br."barbershop_id" = bs."id"
WHERE
  b."id" = $1
  AND bs."id" = $2;

-- name: GetEmailBarber :one
SELECT *
FROM "Barbers"
WHERE email = $1
LIMIT 1;

-- name: ListBarbersInBarberShop :many
SELECT
  b.*,
  br."role_id" as "barber_role_id"
FROM
  "Barbers" b
JOIN
  "BarberRoles" br ON b."id" = br."barber_id"
WHERE
  br."barbershop_id" = $1;

-- name: UpdateBarber :one
UPDATE "Barbers"
SET 
  gender_id = coalesce(sqlc.narg('gender_id'), gender_id),
  email = coalesce(sqlc.narg('email'), email),
  phone = coalesce(sqlc.narg('phone'), phone),
  "hashed_password" = coalesce(sqlc.narg('hashed_password'), hashed_password),
  nick_name = coalesce(sqlc.narg('nick_name'), nick_name),
  full_name = coalesce(sqlc.narg('full_name'), full_name),
  haircut = coalesce(sqlc.narg('haircut'), haircut),
  avatar_url = coalesce(sqlc.narg('avatar_url'), avatar_url),
  "password_changed_at" = coalesce(
    sqlc.narg('password_changed_at'),
    password_changed_at
  ),
  "updated_at" = now()
  WHERE "id" = $1
RETURNING *;

-- name: DeleteBarber :exec
DELETE FROM "Barbers"
WHERE "id" = $1;