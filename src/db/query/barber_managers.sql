-- name: CreateBarberManagers :one
INSERT INTO "BarberManagers" (
    barber_id,
    manager_id
  )
VALUES (
    $1,
    $2
  )
RETURNING *;

-- name: ListBarberManagers :many
SELECT "manager_id"
FROM "BarberManagers"
WHERE "barber_id" = $1;

-- name: DeleteBarberManagers :exec
DELETE FROM "BarberManagers"
WHERE "barber_id" = $1 AND "manager_id" = $2;