-- name: CreateServiceCategory :one
INSERT INTO "ServiceCategories" ("name", "barber_shop_id")
VALUES ($1, $2)
RETURNING *;

-- name: UpdateServiceCategory :exec
UPDATE "ServiceCategories"
SET "name" = coalesce(sqlc.narg('name'), name)
WHERE "id" = $1;

-- name: DeleteServiceCategories :exec
DELETE FROM "ServiceCategories"
WHERE "id" = $1;