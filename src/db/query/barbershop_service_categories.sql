-- name: CreateBarberShopServiceCategories :one
INSERT INTO "BarberShopServiceCategories" (
  "barber_shop_chain_id",
  "barber_shop_id",
  "service_category_id"
)
VALUES (
  $1,
  $2,
  $3
)
RETURNING *
;

-- name: ListBarberShopServiceCategories :many
SELECT
  bsc.*,
  sc."name" as "service_category_name",
  sc."is_global" as "service_category_is_global"
FROM
  "BarberShopServiceCategories" bsc
JOIN
  "ServiceCategories" sc ON bsc."service_category_id" = sc."id"
WHERE
  bsc."barber_shop_id" = $1
  AND bsc."barber_shop_chain_id" = $2;

-- name: UpdateBarberShopServiceCategories :one
UPDATE "BarberShopServiceCategories"
SET
  "service_category_id" = $1
WHERE
  "id" = $2
RETURNING *
;

-- name: DeleteBarberShopServiceCategories :exec
DELETE FROM "BarberShopServiceCategories"
WHERE
  "id" = $1
;