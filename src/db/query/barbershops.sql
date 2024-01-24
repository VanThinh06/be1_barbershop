-- name: CreateBarberShop :one
INSERT INTO "BarberShops" (
                           barbershop_chain_id,
                           name,
                           is_main_branch,
                           branch_count,
                           coordinates,
                           address,
                           image
                           )
VALUES (
        sqlc.arg(barbershop_chain_id),
        sqlc.arg(name),
        sqlc.narg(is_main_branch),
        sqlc.arg(branch_count),
        ST_GeographyFromText('POINT(' || sqlc.arg(longitude)::float8 || ' ' || sqlc.arg(latitude)::float8 || ')'),
        sqlc.arg(address),
        sqlc.narg(image)
        ) RETURNING *; 

-- name: GetBarberShop :one
SELECT *
FROM "BarberShops"
WHERE id = $1;

-- name: ListBarberShopsNearby :many
SELECT
    id,
    barbershop_chain_id,
    name,
    branch_count,
    coordinates,
    address,
    image,
    status,
    rate,
    "reputation",
    CAST(ST_X(ST_GeomFromWKB(coordinates::geometry)) AS float8) AS longitude,
    CAST(ST_Y(ST_GeomFromWKB(coordinates::geometry)) AS float8) AS latitude,
    CAST(ST_Distance(
        ST_SetSRID(ST_MakePoint(sqlc.arg(current_longitude)::float, sqlc.arg(current_latitude)::float), 4326),
        coordinates::geography
    ) AS float) AS distance
FROM "BarberShops"
WHERE  ST_Distance(coordinates, ST_SetSRID(ST_MakePoint(sqlc.arg(current_longitude)::float, sqlc.arg(current_latitude)::float), 4326)) <= sqlc.arg(distance_in_meters)::int
ORDER BY ST_Distance(coordinates, ST_SetSRID(ST_MakePoint(sqlc.arg(current_longitude)::float, sqlc.arg(current_latitude)::float), 4326));

-- name: UpdateBarberShop :one
UPDATE "BarberShops"
SET 
    name = coalesce(sqlc.narg('name'), name),
    is_main_branch = coalesce(sqlc.narg('is_main_branch'), is_main_branch),
    branch_count = coalesce(sqlc.narg('branch_count'), branch_count),
    coordinates = coalesce(ST_GeographyFromText('POINT(' || sqlc.narg(longitude)::float8 || ' ' || sqlc.narg(latitude)::float8 || ')'), coordinates),
    address = coalesce(sqlc.narg('address'), address),
    image = coalesce(sqlc.narg('image'), image),
    status = coalesce(sqlc.narg('status'), status),
    rate = coalesce(sqlc.narg('rate'), rate),
    start_time = coalesce(sqlc.narg('start_time'), start_time),
    end_time = coalesce(sqlc.narg('end_time'), end_time),
    break_time = coalesce(sqlc.narg('break_time'), break_time),
    interval_scheduler = coalesce(sqlc.narg('interval_scheduler'), interval_scheduler),
    update_at = now()
WHERE "id" = $1
RETURNING *;

-- name: DeleteBarberShops :exec
DELETE FROM "BarberShops"
WHERE id = $1;

-- name: QueryBarberShops :many
SELECT bs.*
FROM "BarberShops" bs
WHERE bs."name" = $1
  OR bs."barbershop_chain_id" IN (
    SELECT c."barbershop_chain_id"
    FROM "BarberShopChains" c
    WHERE c."name" = $1
);
