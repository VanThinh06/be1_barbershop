-- name: CreateBarberShop :one

INSERT INTO "BarberShops" (
                           owner_id,
                           chain_id,
                           name,
                           coordinates,
                           address,
                           image,
                           facility
                           )
VALUES (
        sqlc.arg(owner_id),
        sqlc.arg(chain_id),
        sqlc.arg(name),
        ST_GeographyFromText('POINT(' || sqlc.arg(longitude)::float8 || ' ' || sqlc.arg(latitude)::float8 || ')'),
        sqlc.arg(address),
        sqlc.narg(image),
        sqlc.arg(facility)
        ) RETURNING *; 
-- name: GetBarberShop :one

SELECT "shop_id"
FROM "BarberShops"
WHERE shop_id = $1
LIMIT 1;

-- name: FindBarberShopsNearbyLocations :many

SELECT
    owner_id,
    shop_id,
    status,
    name,
    coordinates,
    address,
    image,
    list_image,
    created_at,
    CAST(ST_X(ST_GeomFromWKB(coordinates::geometry)) AS float8) AS longitude,
    CAST(ST_Y(ST_GeomFromWKB(coordinates::geometry)) AS float8) AS latitude,
    CAST(ST_Distance(
        ST_SetSRID(ST_MakePoint(sqlc.arg(current_longitude)::float, sqlc.arg(current_latitude)::float), 4326),
        coordinates::geography
    ) AS float) AS distance
FROM "BarberShops"
WHERE  ST_Distance(coordinates, ST_SetSRID(ST_MakePoint(sqlc.arg(current_longitude)::float, sqlc.arg(current_latitude)::float), 4326)) <= sqlc.arg(distance_in_meters)::int
ORDER BY ST_Distance(coordinates, ST_SetSRID(ST_MakePoint(sqlc.arg(current_longitude)::float, sqlc.arg(current_latitude)::float), 4326));

-- name: BarberShopInChain :one
SELECT EXISTS (
  SELECT 1
  FROM "BarberShops"
  WHERE "owner_id" = $1
    AND "chain_id" IS NULL
) AS "barbershop_not_in_chain";

-- name: UpdateChainForBarberShops :exec  
UPDATE "BarberShops"
SET "chain_id" = sqlc.arg(chain_id)::uuid
WHERE "owner_id" = $1 AND "chain_id" IS NULL;

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