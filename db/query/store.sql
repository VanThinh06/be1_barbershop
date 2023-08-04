-- name: GetStore :one

SELECT *
FROM store
WHERE id = $1
LIMIT 1
;


-- name: GetListStore :many
SELECT * FROM store
LIMIT $1
OFFSET $2;

-- name: CreateStore :one
INSERT INTO store (
name_id,
name_store,
location,
address,
image,
list_image,
manager_id,
employee_id,
status
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;
