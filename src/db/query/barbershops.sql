
-- name: GetBarberShop :one
SELECT *
FROM "BarberShops"
WHERE id = $1;

-- name: SearchByNameBarberShops :many
SELECT
    bs.id,
    bs.barber_shop_chain_id,
    bs.name,
    bs.coordinates,
    bs."is_reputation",
    CAST(ST_X(ST_GeomFromWKB(bs.coordinates::geometry)) AS float8) AS longitude,
    CAST(ST_Y(ST_GeomFromWKB(bs.coordinates::geometry)) AS float8) AS latitude,
    CAST(ST_Distance(
        ST_SetSRID(ST_MakePoint(sqlc.arg(current_longitude)::float, sqlc.arg(current_latitude)::float), 4326),
        bs.coordinates::geography
    ) AS float) AS distance
FROM "BarberShops" bs
JOIN "BarberShopChains" bsc ON bs.barber_shop_chain_id = bsc.id
WHERE bsc."name" = $1
ORDER BY ST_Distance(bs.coordinates, ST_SetSRID(ST_MakePoint(sqlc.arg(current_longitude)::float, sqlc.arg(current_latitude)::float), 4326));



-- name: ListNearbyBarberShops :many
SELECT
    id,
    barber_shop_chain_id,
    name,
    coordinates,
    "is_reputation",
    CAST(ST_X(ST_GeomFromWKB(coordinates::geometry)) AS float8) AS longitude,
    CAST(ST_Y(ST_GeomFromWKB(coordinates::geometry)) AS float8) AS latitude,
    CAST(ST_Distance(
        ST_SetSRID(ST_MakePoint(sqlc.arg(current_longitude)::float, sqlc.arg(current_latitude)::float), 4326),
        coordinates::geography
    ) AS float) AS distance
FROM "BarberShops"
WHERE  ST_Distance(coordinates, ST_SetSRID(ST_MakePoint(sqlc.arg(current_longitude)::float, sqlc.arg(current_latitude)::float), 4326)) <= sqlc.arg(distance_in_meters)::int
ORDER BY ST_Distance(coordinates, ST_SetSRID(ST_MakePoint(sqlc.arg(current_longitude)::float, sqlc.arg(current_latitude)::float), 4326));


-- name: CreateBarberShop :one
INSERT INTO "BarberShops" (
                           barber_shop_chain_id,
                           name,
                           province_id,
                           district_id,
                           ward_id,
                           specific_location,
                           phone,
                           email,
                           website_url,
                           avatar_url,
                           cover_photo_url,
                           facade_photo_url,
                           representative_name,
                           citizen_id,
                           representative_phone,
                           representative_email,
                           representative_phone_other,
                           start_time_monday,
                           end_time_monday,
                           start_time_tuesday,
                           end_time_tuesday,
                           start_time_wednesday,
                           end_time_wednesday,
                           start_time_thursday,
                           end_time_thursday,
                           start_time_friday,
                           end_time_friday,
                           start_time_saturday,
                           end_time_saturday,
                           start_time_sunday,
                           end_time_sunday,
                           interval_scheduler,
                           coordinates
                           )
VALUES (
         sqlc.narg(barber_shop_chain_id),
$1,
$2,
$3,
$4,
$5,
$6,
$7,
         sqlc.narg(website_url),

$8,
$9,
$10,
$11,
$12,
$13,
$14,
         sqlc.narg(representative_phone_other),

$15,
$16,
$17,
$18,
$19,
$20,
$21,
$22,
$23,
$24,
$25,
$26,
$27,
$28,
$29,
        ST_GeographyFromText('POINT(' || sqlc.arg(longitude)::float8 || ' ' || sqlc.arg(latitude)::float8 || ')')
        ) RETURNING *; 

-- name: UpdateBarberShop :one
UPDATE "BarberShops"
SET 
    name = coalesce(sqlc.narg('name'), name),
    is_main_branch = coalesce(sqlc.narg('is_main_branch'), is_main_branch),
    coordinates = coalesce(ST_GeographyFromText('POINT(' || sqlc.narg(longitude)::float8 || ' ' || sqlc.narg(latitude)::float8 || ')'), coordinates),
    interval_scheduler = coalesce(sqlc.narg('interval_scheduler'), interval_scheduler),
    is_reputation = coalesce(sqlc.arg('is_reputation'), is_reputation),
    is_verified = coalesce(sqlc.arg('is_verified'), is_verified),
    update_at = now()
WHERE "id" = $1
RETURNING *;

-- name: DeleteBarberShops :exec
DELETE FROM "BarberShops"
WHERE id = $1;


