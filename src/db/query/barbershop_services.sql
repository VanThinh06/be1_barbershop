-- name: CreateBarberShopService :one
INSERT INTO "BarberShopServices" (
barber_shop_id,
category_id,
gender_id,
name,
timer,
price,
description,
image_url,
combo_services
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetTimerBarberShopServices :one
SELECT SUM("timer") AS total_timer
FROM "BarberShopServices"
WHERE "id" = ANY($1::uuid[]);

-- name: ListBarberShopServices :many
SELECT "id", "name", "category_id"
FROM "BarberShopServices"
WHERE "barber_shop_id" = $1;

-- -- name: DeleteBarberShopServices :exec
-- DELETE FROM "BarberShopServices"
-- WHERE
--   "id" = $1;