-- name: CreateBarberShopService :one
INSERT INTO "BarberShopServices" (
barber_shop_id,
category_id,
gender_id,
name,
timer,
price,
description,
image_url,
combo_services
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetBarberShopService :one
SELECT bs.*, sc."name" AS "category_name"
FROM "BarberShopServices" bs
LEFT JOIN 
    "ServiceCategories" sc ON sc."id" = bs."category_id"
WHERE bs."id" = $1;

-- name: GetTimerBarberShopServices :one
SELECT SUM("timer") AS total_timer
FROM "BarberShopServices"
WHERE "id" = ANY($1::uuid[]);

-- name: ListBarberShopServices :many
SELECT bs."id", bs."name", sc."name" as "category_name"
FROM "BarberShopServices" bs
JOIN "ServiceCategories" sc ON bs."category_id" = sc."id"
WHERE bs."barber_shop_id" = $1
AND bs."combo_services" IS NULL;

-- name: ListServicesByCategory :many
SELECT 
    sc."id" AS "category_id", 
    sc."name" AS "category_name", 
    bs."id" AS "service_id", 
    bs."name" AS "service_name",
    bs."timer",
    bs."price",
    bs."description",
    bs."image_url",
    bs."gender_id",
    bs."is_active",
    bs."discount_price",
    bs."discount_start_time",
    bs."discount_end_time",
    combo_services
FROM 
    "ServiceCategories" sc
LEFT JOIN 
    "BarberShopServices" bs ON sc."id" = bs."category_id"
WHERE 
    bs."barber_shop_id" = $1
    AND (bs."combo_services" IS NULL OR bs."combo_services" = '{}')
ORDER BY
	sc."id",
    bs."gender_id";


-- name: ListComboServices :many
SELECT 
    bs."id" AS "service_id",
    bs."name" AS "service_name",
    bs."gender_id",
    bs."timer",
    bs."price",
    bs."description",
    bs."image_url",
    bs."is_active",
    bs."discount_price",
    bs."discount_start_time",
    bs."discount_end_time",
    combo_services
FROM 
    "BarberShopServices" bs
WHERE 
    bs."barber_shop_id" = $1
    AND bs."combo_services" IS NOT NULL AND array_length(bs."combo_services", 1) > 0
ORDER BY
    bs."gender_id";

-- name: UpdateBarberShopService :exec
UPDATE "BarberShopServices"
SET 
    name = coalesce(sqlc.narg('name'), name),
    timer = coalesce(sqlc.narg('timer'), timer),
    category_id = coalesce(sqlc.narg('category_id'), category_id),
    gender_id = coalesce(sqlc.narg('gender_id'), gender_id),
    price = coalesce(sqlc.narg('price'), price),
    description = coalesce(sqlc.narg('description'), description),
    image_url = coalesce(sqlc.narg('image_url'), image_url),
    is_active = coalesce(sqlc.narg('is_active'), is_active),
    discount_price = sqlc.narg('discount_price'),
    discount_start_time = sqlc.narg('discount_start_time'),
    discount_end_time = sqlc.narg('discount_end_time'),
    combo_services = CASE 
                        WHEN COALESCE(sqlc.narg('combo_services'), '{}')::text[] != '{}' THEN COALESCE(sqlc.narg('combo_services'), '{}')::text[]
                        ELSE combo_services
                    END
WHERE "id" = $1;


-- -- name: DeleteBarberShopServices :exec
-- DELETE FROM "BarberShopServices"
-- WHERE
--   "id" = $1;