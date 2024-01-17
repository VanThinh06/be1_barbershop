-- name: InsertBarberShopChains
INSERT INTO "BarberShopChains" ("name")
VALUES ($1)
RETURNING *;

-- name: UpdateBarberShopChain
UPDATE "BarberShopChains"
SET "name" = $1
WHERE "id" = $2
RETURNING *;

-- name: DeleteBarberShopChain
DELETE FROM "BarberShopChains"
WHERE "id" = $1
RETURNING *;