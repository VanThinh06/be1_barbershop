-- name: CreateUsers :one

INSERT INTO users (username, full_name, email, hashed_password, image, role)
VALUES ($1,
        $2,
        $3,
        $4,
        $5, $6
        ) RETURNING *;

-- name: GetUsers :one

SELECT *
FROM users
WHERE username = $1
LIMIT 1;
