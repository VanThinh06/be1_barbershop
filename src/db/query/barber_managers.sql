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
WHERE "id" = $1;

-- name: UpdateBarberManagers :one
UPDATE "BarberManagers"
SET "manager_id" = $1,
    "update_at" = NOW()
WHERE "id" = $2
RETURNING *;

-- name: DeleteBarberManagers :exec
DELETE FROM "BarberManagers"
WHERE "id" = $1;