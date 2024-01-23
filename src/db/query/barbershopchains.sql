-- name: CreateBarberShopChains :one
INSERT INTO "BarberShopChains" ("name", "description", "founder", "founding_date", website)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetBarberShopChains :one
SELECT *
FROM "BarberShopChains"
WHERE "id" = $1;

-- name: UpdateBarberShopChains :one
UPDATE "BarberShopChains"
SET "name" = coalesce(sqlc.narg('name'), name),
    "description" = coalesce(sqlc.narg('description'), description),
    founder = coalesce(sqlc.narg('founder'),founder),
    founding_date = coalesce(sqlc.narg('founding_date'), founding_date),
    website = coalesce(sqlc.narg('website'), website)
    "update_at" = NOW()
WHERE "id" = sqlc.arg('id')
RETURNING *;

-- name: DeleteBarberShopChains :exec
DELETE FROM "BarberShopChains"
WHERE "id" = $1;