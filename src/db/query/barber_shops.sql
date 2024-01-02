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

-- name: UpdateBarberShop :one
UPDATE "BarberShops"
SET 
    name = coalesce(sqlc.narg('name'), name),
    facility = coalesce(sqlc.narg('facility'), facility),
    address = coalesce(sqlc.narg('address'), address),
    coordinates = coalesce(ST_GeographyFromText('POINT(' || sqlc.narg(longitude)::float8 || ' ' || sqlc.narg(latitude)::float8 || ')'), coordinates),
    image = coalesce(sqlc.narg('image'), image),
    start_time = coalesce(sqlc.narg('start_time'), start_time),
    end_time = coalesce(sqlc.narg('end_time'), end_time),
    break_time = coalesce(sqlc.narg('break_time'), break_time),
    status = coalesce(sqlc.narg('status'), status),
    interval_scheduler = coalesce(sqlc.narg('interval_scheduler'), interval_scheduler),
    updated_at = now()
WHERE "shop_id" = $1
RETURNING *;

-- name: GetBarberShop :one

SELECT *
FROM "BarberShops"
WHERE shop_id = $1;

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


-- name: QueryBarberShops :many
SELECT bs.*
FROM "BarberShops" bs
WHERE bs."name" = $1
  OR bs."chain_id" IN (
    SELECT c."chain_id"
    FROM "Chains" c
    WHERE c."name" = $1
);


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