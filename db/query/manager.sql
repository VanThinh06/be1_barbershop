-- name: CreateManager :one
INSERT INTO manager ( username, role, image, store_id)
VALUES ( $1,
         $2,
         $3,
         $4) RETURNING *;

-- name: UpdateManager :one
UPDATE manager
set username = $2,
    image = $3,
    role = $4
WHERE id = $1 RETURNING *;
