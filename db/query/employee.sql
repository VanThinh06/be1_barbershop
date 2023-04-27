-- name: GetEmployee :one
SELECT *
FROM employee
WHERE id = $1
LIMIT 1;

-- name: GetListEmployeewithStore :many
SELECT *
FROM employee
WHERE store_id = $1
LIMIT $2;

-- name: CreateEmployee :one
INSERT INTO employee ( username, role, image, store_id)
VALUES ( $1,
         $2,
         $3,
         $4) RETURNING *;

-- name: UpdateEmployee :one
UPDATE employee
set username = $2,
    image = $3,
    store_id = $4
WHERE id = $1 RETURNING *;

-- name: DeleteEmployee :exec
DELETE
FROM employee
WHERE id = $1;

