-- name: CreateBarber :one

INSERT INTO barber(username, full_name, email, hashed_password, avatar, role, status, store_id, store_manager)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        $9) RETURNING *;

-- name: GetBarber :one

SELECT username, full_name, email, avatar, role, status, store_id, store_manager
FROM barber
WHERE username = $1
LIMIT 1
;
