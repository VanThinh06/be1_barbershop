-- name: CreateBarberShopServiceCategoryWithChain :one
INSERT INTO "BarberShopServiceCategories" (
  "barbershop_chain_id",
  "service_category_id"
)
VALUES (
  $1,
  $2
)
RETURNING *
;

-- name: CreateBarberShopServiceCategoryWithShop :one
INSERT INTO "BarberShopServiceCategories" (
  "barbershop_id",
  "service_category_id"
)
VALUES (
  $1,
  $2
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
  bsc."barbershop_id" = $1
  AND bsc."barbershop_chain_id" = $2;

-- name: UpdateBarberShopServiceCategory :exec
UPDATE "BarberShopServiceCategories"
SET
  "service_category_id" = $1
WHERE
  "id" = $2;;

-- name: DeleteBarberShopServiceCategory :exec
DELETE FROM "BarberShopServiceCategories"
WHERE
  "id" = $1;;