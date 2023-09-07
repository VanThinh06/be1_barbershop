-- name: GetStore :one
SELECT *
FROM store
WHERE id = $1
LIMIT 1;

-- name: GetListStore :many
SELECT *
FROM store
WHERE 
	 $1 = ANY (manager_id)
LIMIT $2 OFFSET $3
;

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
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: UpdateStore :one
UPDATE store
set name_id = $2,
  name_store = $3,
  location = $4,
  address =$5,
  image =$6,
  list_image =$7,
  manager_id = $8,
  employee_id =$9,
  status =$10,
  update_at = $11
WHERE id = $1
RETURNING *;

-- name: DeleteStore :one
DELETE FROM store
WHERE id = $1
RETURNING *;