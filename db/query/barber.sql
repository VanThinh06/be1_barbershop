-- name: CreateBarber :one

INSERT INTO barber(name_id, store_id, store_manager, status)
VALUES ($1,
        $2,
        $3,
        $4
        ) RETURNING *;

-- name: GetBarber :one

SELECT *
FROM barber
WHERE name_id = $1
LIMIT 1;

