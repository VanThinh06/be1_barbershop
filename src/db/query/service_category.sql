
-- name: CreateServicesCategoryPublic :one
INSERT INTO "ServiceCategory" (
"name",
"chain_id"
) VALUES (
  $1, $2
)
RETURNING *;

-- name: CreateServicesCategoryPrivate :one
INSERT INTO "ServiceCategory" (
shop_id,
"name"
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateServicesCategoryPublicSeparate :one
INSERT INTO "ServiceCategory" (
    "shop_id",
    "chain_id",
    "name"
  )
VALUES ($1, $2, $3)
RETURNING *;

-- GetServicesCategoryBarberShops :many
SELECT * FROM "ServiceCategory" AS sc1
WHERE sc1."shop_id" = $1 AND sc1."chain_id" = $2

UNION

SELECT * FROM "ServiceCategory" AS sc2
WHERE sc2."shop_id" = $1 AND sc2."chain_id" = $2 AND sc2."shop_id" IS NOT NULL AND sc2."chain_id" IS NOT NULL;

-- -- name: GetListServiceCategorywithStore :many
-- SELECT *
-- FROM service_category
-- WHERE store_id = $1;

-- -- name: DeleteServiceCategory :one
-- DELETE FROM service_category
-- WHERE id = $1
-- RETURNING *;

-- -- name: UpdateServiceCategory :one
-- UPDATE service_category
-- set store_id = $2,
--   work= $3,
--   description =$4
-- WHERE id = $1
-- RETURNING *;