-- name: CreateBarberShopService :one
INSERT INTO "BarberShopServices" (
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
    "BarberShopServices" bs ON sc."id" = bs."category_id"
LEFT JOIN 
    "CategoryPositions" cp ON sc."id" = cp."category_id"
WHERE 
    bs."barber_shop_id" = $1
    AND (cp."visible" = false)  
ORDER BY
    cp."position",  -- Sắp xếp theo vị trí của category
    sc."id",
    bs."gender_id"; -- Để phân loại theo gender_id nếu cần thiết


-- name: UpdateBarberShopService :exec
UPDATE "BarberShopServices"
SET 
    name = COALESCE($2, name),
    timer = COALESCE($3, timer),
    category_id = COALESCE($4, category_id),
    gender_id = COALESCE($5, gender_id),
    price = COALESCE($6, price),
    description = COALESCE($7, description),
    image_url = COALESCE($8, image_url),
    is_active = COALESCE($9, is_active),
    discount_price = $10,
    discount_start_time = $11,
    discount_end_time = $12
WHERE "id" = $1;


-- -- name: DeleteBarberShopServices :exec
-- DELETE FROM "BarberShopServices"
-- WHERE
--   "id" = $1;


-- COMBO SERVICE
-- name: CreateComboServices :one
INSERT INTO "ComboServices" (
barber_shop_id,
name,
gender_id,
timer,
price,
description,
image_url
  )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;


-- name: CreateComboServiceItems :one
INSERT INTO "ComboServiceItems" (
combo_service_id,
barber_shop_service_id
  )
VALUES ($1, $2)
RETURNING *;


-- name: GetComboService :many
SELECT 
  cs.id,
  cs.name AS combo_service_name,
  cs.description AS combo_service_description,
  cs.price AS combo_service_price,
  cs.image_url AS combo_service_image_url,
  cs.timer AS combo_service_timer,
  cs.gender_id AS combo_service_gender,
  cs.is_active AS combo_service_is_active,
  bss.id AS barber_shop_service_id,
  bss.name AS barber_shop_service_name,
  bss.price AS barber_shop_service_price
FROM 
  "ComboServiceItems" csi
JOIN 
  "ComboServices" cs ON csi.combo_service_id = cs.id
JOIN 
  "BarberShopServices" bss ON csi.barber_shop_service_id = bss.id
WHERE 
  cs.id = $1;

-- name: ListComboServices :many
SELECT 
  cs.id,
  cs.gender_id AS combo_service_gender,
  cs.name AS combo_service_name,
  cs.timer AS combo_service_timer,
  cs.price AS combo_service_price,
  cs.discount_price AS combo_service_discount_price,
  cs.discount_start_time AS combo_service_discount_start_time,
  cs.discount_end_time AS combo_service_discount_end_time,
  cs.description AS combo_service_description,
  cs.image_url AS combo_service_image_url,
  cs.is_active AS combo_service_is_active,
  ARRAY_AGG(csi.barber_shop_service_id) AS barber_shop_service_ids
FROM 
  "ComboServiceItems" csi
JOIN 
  "ComboServices" cs ON csi.combo_service_id = cs.id
WHERE 
  cs.barber_shop_id = $1
GROUP BY 
  cs.id, cs.gender_id, cs.name, cs.timer, cs.price, cs.discount_price, 
  cs.discount_start_time, cs.discount_end_time, cs.description, cs.image_url, cs.is_active;
