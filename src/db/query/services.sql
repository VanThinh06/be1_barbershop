-- name: CreateService :one
INSERT INTO "Services" (
    category_id,
    "name",
    timer,
    price,
    "description",
    "image"
  )
VALUES ($1, $2, $3, $4, $5, $6)
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
WHERE (("category_id" = $1 AND "shop_id" IS NULL) OR ("category_id" = $1 AND "shop_id" = $2))
  AND "hidden" = false;

-- name: GetListServiceDetails :many
SELECT
    sc.id AS category_id,
    sc.name AS category_name,
    s.id AS service_id,
    s.name AS service_name,
    s.timer,
    s.price,
    s.description,
    s.image
FROM
    "ServiceCategories" sc
JOIN
    "Services" s ON sc.id = s.category_id
WHERE
    (sc.chain_id = $1 OR
    sc.shop_id = $2) AND
    sc.hidden = false AND
    s.hidden = false;

-- name: GetTimerService :one
SELECT SUM("timer") AS total_timer
FROM "Services"
WHERE "id" IN ( sqlc.arg(services_id)::uuid[] )

;

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