-- -- name: CreateService :one
-- INSERT INTO service (
--     service_category_id,
--     work,
--     timer,
--     price,
--     description,
--     image
--   )
-- VALUES ($1, $2, $3, $4, $5, $6)
-- RETURNING *;

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