-- name: CreateStore :one
INSERT INTO store (
name_id,
name_store,
location,
address,
image,
list_image,
manager_id,
status
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;