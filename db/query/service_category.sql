
-- name: CreateServiceCategory :one
INSERT INTO service_category (
store_id,
work,
description
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetServiceCategory :one
SELECT *
FROM service_category
WHERE id = $1
LIMIT 1;

-- name: GetListServicewithStore :many
SELECT *
FROM service_category
WHERE store_id = $1;