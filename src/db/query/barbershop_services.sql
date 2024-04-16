-- name: CreateBarberShopService :one
INSERT INTO "BarberShopServices" (
barber_shop_id,
category_id,
gender_id,
name,
timer,
price,
description,
image_url
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- -- name: DeleteBarberShopServices :exec
-- DELETE FROM "BarberShopServices"
-- WHERE
--   "id" = $1;