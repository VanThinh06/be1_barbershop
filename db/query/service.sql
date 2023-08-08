-- name: CreateService :one
INSERT INTO service (
    service_category_id,
    work,
    timer,
    price,
    description,
    image
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetListServicewithCategory :many
SELECT *
FROM "service"
WHERE service_category_id = $1
LIMIT $2 OFFSET $3;

-- name: GetService :one
SELECT *
FROM "service"
WHERE id = $1
LIMIT 1;

-- name: UpdateService :one
UPDATE "service"
set work = $2,
  timer = $3,
  price = $4,
  description = $5,
  image = $6
WHERE id = $1
RETURNING *;

-- name: DeleteServicewithStoreCategory :exec
DELETE FROM service
WHERE service_category_id = $1;