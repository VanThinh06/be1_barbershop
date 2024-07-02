-- name: CreateServiceItem :one
INSERT INTO "ServiceItems" (
barber_shop_id,
category_id,
gender_id,
name,
timer,
price,
description,
image_url
  )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;


-- name: GetServiceItem :one
SELECT bs.*, sc."name" AS "category_name"
FROM "ServiceItems" bs
LEFT JOIN 
    "ServiceCategories" sc ON sc."id" = bs."category_id"
WHERE bs."id" = $1;


-- name: GetTimerServiceItem :one
SELECT SUM("timer") AS total_timer
FROM "ServiceItems"
WHERE "id" = ANY($1::uuid[]);


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
    bs."discount_end_time"
FROM 
    "ServiceCategories" sc
LEFT JOIN 
    "ServiceItems" bs ON sc."id" = bs."category_id"
LEFT JOIN 
    "CategoryPositions" cp ON sc."id" = cp."category_id"
WHERE 
    bs."barber_shop_id" = $1
    AND (cp."visible" = true)  
ORDER BY
    cp."position",  -- Sắp xếp theo vị trí của category
    sc."id",
    bs."gender_id"; -- Để phân loại theo gender_id nếu cần thiết


-- name: UpdateServiceItem :exec
UPDATE "ServiceItems"
SET 
    name = COALESCE($2, name),
    timer = COALESCE($3, timer),
    category_id = COALESCE($4, category_id),
    gender_id = COALESCE($5, gender_id),
    price = COALESCE($6, price),
    description = COALESCE($7, ''),
    image_url = COALESCE($8, ''),
    is_active = COALESCE($9, is_active),
    discount_price = COALESCE($10, 0),
    discount_start_time = $11,
    discount_end_time = $12
WHERE "id" = $1;


-- -- name: DeleteServiceItem :exec
-- DELETE FROM "ServiceItems"
-- WHERE
--   "id" = $1;

