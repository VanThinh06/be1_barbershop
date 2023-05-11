-- name: CreateUsers :one

INSERT INTO users (username, full_name, email, hashed_password)
VALUES ($1,
        $2,
        $3,
        $4) RETURNING *;

-- name: GetUsers :one

SELECT *
FROM users
WHERE username = $1
LIMIT 1;
