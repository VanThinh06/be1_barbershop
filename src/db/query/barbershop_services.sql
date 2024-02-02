-- name: CreateBarberShopServices :one
INSERT INTO "BarberShopServices" (
    barbershop_category_id,
    barber_shop_chain_id,
    barber_shop_id,
    gender_id,
    "name",
    "timer",
    price,
    description,
    image
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: DeleteBarberShopServices :exec
DELETE FROM "BarberShopServices"
WHERE
  "id" = $1;