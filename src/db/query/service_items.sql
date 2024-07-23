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
    cp."position",  
    sc."id",
    bs."gender_id"; 


-- name: UpdateServiceItem :one
UPDATE "ServiceItems"
SET
  name = COALESCE(sqlc.narg('name'), name),
  gender_id = COALESCE(sqlc.narg('gender_id'), gender_id),
  timer = COALESCE(sqlc.narg('timer'), timer),
  category_id = COALESCE(sqlc.narg('category_id'), category_id),
  price = COALESCE(sqlc.narg('price'), price),
  discount_price = COALESCE(sqlc.narg('discount_price'), discount_price),
  discount_start_time = COALESCE(sqlc.narg('discount_start_time'), discount_start_time),
  discount_end_time = COALESCE(sqlc.narg('discount_end_time'), discount_end_time),
  description = COALESCE(sqlc.narg('description'), description),
  image_url = COALESCE(sqlc.narg('image_url'), image_url),
  is_active = COALESCE(sqlc.narg('is_active'), is_active)
WHERE id = sqlc.arg('id')
RETURNING *;


-- name: DeleteServiceItem :exec
DELETE FROM "ServiceItems"
WHERE
  "id" = $1;

