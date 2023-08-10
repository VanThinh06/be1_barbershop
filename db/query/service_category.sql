
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

-- name: GetListServiceCategorywithStore :many
SELECT *
FROM service_category
WHERE store_id = $1;

-- name: DeleteServiceCategory :one
DELETE FROM service_category
WHERE id = $1
RETURNING *;

-- name: UpdateServiceCategory :one
UPDATE service_category
set store_id = $2,
  work= $3,
  description =$4
WHERE id = $1
RETURNING *;