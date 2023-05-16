-- name: CreateStore :one

INSERT INTO store (
name_id,
name_store,
location,
image,
manager_id,
employee_id,
status
)
VALUES ($1,
        $2,
        $3,
        $4,
        $5, $6, $7
        ) RETURNING *;

-- name: GetStore :one

SELECT *
FROM store
WHERE name_id = $1
LIMIT 1;
