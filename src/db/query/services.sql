-- name: CreateService :one
INSERT INTO "Services" (
    category_id,
    "chain_id",
    "name",
    timer,
    price,
    "description",
    "image"
  )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: CreateServicePrivate :one
INSERT INTO "Services" (
    category_id,
    "shop_id",
    "name",
    timer,
    price,
    "description",
    "image"
  )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetListServices :many
SELECT *
FROM "Services"
WHERE ("chain_id" = $1 OR "shop_id" = $2)
  AND "hidden" = false;

-- -- name: GetListServicewithCategory :many
-- SELECT *
-- FROM "service"
-- WHERE service_category_id = $1
-- LIMIT $2 OFFSET $3;

-- -- name: GetService :one
-- SELECT *
-- FROM "service"
-- WHERE id = $1
-- LIMIT 1;

-- -- name: UpdateService :one
-- UPDATE "service"
-- set service_category_id = $2,
--   work = $3,
--   timer = $4,
--   price = $5,
--   description = $6,
--   image = $7
-- WHERE id = $1
-- RETURNING *;

-- -- name: DeleteServicewithStoreCategory :exec
-- DELETE FROM "service"
-- WHERE service_category_id = $1
-- ;

-- -- name: DeleteService :one
-- DELETE FROM "service"
-- WHERE id = $1
-- RETURNING *
-- ;