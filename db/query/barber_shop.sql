-- name: CreateBarberShops :one

INSERT INTO "BarberShops" (code_barber_shop,
                           owner_id,
                           name,
                           location,
                           address,
                           image)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6) RETURNING *;

-- name: GetCodeBarberShop :one

SELECT "shop_id"
FROM "BarberShops"
WHERE code_barber_shop = $1
LIMIT 1;

-- -- name: UpdateStore :one
-- UPDATE store
-- set name_id = $2,
--   name_store = $3,
--   location = $4,
--   address = $5,
--   image = $6,
--   list_image = $7,
--   manager_id = $8,
--   employee_id = $9,
--   status = $10,
--   update_at = $11
-- WHERE id = $1
-- RETURNING *;
-- -- name: DeleteStore :one
-- DELETE FROM store
-- WHERE id = $1
-- RETURNING *;