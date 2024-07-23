-- name: InsertBarberShop :one
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
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32,
      ST_GeographyFromText('POINT(' || sqlc.arg(longitude)::float8 || ' ' || sqlc.arg(latitude)::float8 || ')')
) RETURNING *; 

-- name: GetBarberShopById :one
SELECT *
FROM "view_barber_shops" bs
WHERE bs.id = $1;

-- name: ListBarberShopsByBarberId :many
SELECT DISTINCT 
    bs.id,
    bs.name,
    bs.specific_location,
    bs.phone,
    bs.email,
    bs.website_url,
    bs.avatar_url,
    bs.cover_photo_url,
    bs.facade_photo_url,
    bs.is_main_branch,
    bs.is_reputation,
    bs.province_name,
    bs.district_name,
    bs.ward_name,
    br.role_id AS role_id
FROM "view_barber_shops" bs
LEFT JOIN "BarberRoles" br ON bs."id" = br."barber_shop_id"
WHERE br."barber_id" = $1;

-- name: UpdateBarberShopById :one
UPDATE "BarberShops"
SET 
    name = COALESCE(sqlc.narg(name), name),
    is_main_branch = COALESCE($2, is_main_branch),
    coordinates = COALESCE(ST_GeographyFromText('POINT(' || sqlc.narg(longitude)::float8 || ' ' || sqlc.narg(latitude)::float8 || ')'), coordinates),
    interval_scheduler = COALESCE($3, interval_scheduler),
    is_reputation = COALESCE($4, is_reputation),
    is_verified = COALESCE($5, is_verified),
    update_at = now()
WHERE "id" = $1
RETURNING *;

-- name: DeleteBarberShopById :exec
DELETE FROM "BarberShops"
WHERE id = $1;




-- name: SearchBarberShopsByName :many
SELECT
    bs.id,
    bs.barber_shop_chain_id,
    bs.name,
    bs.coordinates,
    bs."is_reputation",
    CAST(ST_X(ST_GeomFromWKB(bs.coordinates::geometry)) AS float8) AS longitude,
    CAST(ST_Y(ST_GeomFromWKB(bs.coordinates::geometry)) AS float8) AS latitude,
    CAST(ST_Distance(
        ST_SetSRID(ST_MakePoint($2::float, $3::float), 4326),
        bs.coordinates::geography
    ) AS float) AS distance
FROM "BarberShops" bs
JOIN "BarberShopChains" bsc ON bs.barber_shop_chain_id = bsc.id
WHERE bsc."name" = $1
ORDER BY ST_Distance(bs.coordinates, ST_SetSRID(ST_MakePoint($2::float, $3::float), 4326));

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
        ST_SetSRID(ST_MakePoint($2::float, $3::float), 4326),
        coordinates::geography
    ) AS float) AS distance
FROM "BarberShops"
WHERE ST_Distance(coordinates, ST_SetSRID(ST_MakePoint($2::float, $3::float), 4326)) <= $1::int
ORDER BY ST_Distance(coordinates, ST_SetSRID(ST_MakePoint($2::float, $3::float), 4326));
