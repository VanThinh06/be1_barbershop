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

-- name: GetBarberManagers :one
SELECT *
FROM "BarberManagers"
WHERE "barber_id" = $1;

-- name: DeleteBarberManagers :exec
DELETE FROM "BarberManagers"
WHERE "barber_id" = $1 AND "manager_id" = $2;