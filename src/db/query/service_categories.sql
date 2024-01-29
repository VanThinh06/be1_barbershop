-- name: CreateServiceCategories :one
INSERT INTO "ServiceCategories" ("name", "is_global")
VALUES ($1, $2)
RETURNING *;

-- name: GetServiceCategories :one
SELECT *
FROM "ServiceCategories"
WHERE "id" = $1;

-- name: UpdateServiceCategories :one
UPDATE "ServiceCategories"
SET "name" = coalesce(sqlc.narg('name'), name),
    "is_global" = coalesce(sqlc.narg('is_global'), is_global),
    "update_at" = NOW()
WHERE "id" = sqlc.arg('id')
RETURNING *;

-- name: DeleteServiceCategories :exec
DELETE FROM "ServiceCategories"
WHERE "id" = $1;