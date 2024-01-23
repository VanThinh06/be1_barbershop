-- name: CreateBarberShopChains :one
INSERT INTO "BarberShopChains" ("name")
VALUES ($1)
RETURNING *;

-- name: UpdateBarberShopChain :one
UPDATE "BarberShopChains"
SET "name" = $1
WHERE "id" = $2
RETURNING *;

-- name: DeleteBarberShopChain :exec
DELETE FROM "BarberShopChains"
WHERE "id" = $1
RETURNING *;