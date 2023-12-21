
-- name: CreateServiceCategory :one
INSERT INTO "ServiceCategories" (
"name",
"chain_id",
"gender"
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: CreateServiceCategoryPrivate :one
INSERT INTO "ServiceCategories" (
shop_id,
"name",
"gender"
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetListServiceCategories :many
SELECT sc.*
FROM "ServiceCategories" sc
JOIN "BarberShops" bs ON sc."shop_id" = bs."shop_id"
WHERE (sc."shop_id" = $1 OR sc."chain_id" = bs."chain_id")
  AND sc."hidden" = false;

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