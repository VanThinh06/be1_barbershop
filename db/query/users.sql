-- name: CreateUsers :one

INSERT INTO users (username, full_name, email, hashed_password, fcm_device, image, role, address )
VALUES ($1,
        $2,
        $3,
        $4,
        $5, $6, $7, $8
        ) RETURNING *;

-- name: GetUsers :one

SELECT *
FROM users
WHERE username = $1
LIMIT 1;

-- name: DeleteUsers :exec
DELETE FROM users WHERE username = $1;
