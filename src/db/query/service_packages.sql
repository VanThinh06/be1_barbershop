
-- name: CreateServicePackage :one
INSERT INTO "ServicePackages" (
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


-- name: CreateServicePackageItem :one
INSERT INTO "ServicePackageItems" (
service_package_id,
service_item_id
  )
VALUES ($1, $2)
RETURNING *;


-- name: GetServicePackage :many
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
  "ServicePackageItems" csi
JOIN 
  "ServicePackages" cs ON csi.combo_service_id = cs.id
JOIN 
  "ServiceItems" bss ON csi.barber_shop_service_id = bss.id
WHERE 
  cs.id = $1;


-- name: ListServicePackages :many
SELECT * FROM "view_service_packages"
WHERE barber_shop_id = $1;


-- name: UpdateServicePackage :one
UPDATE "ServicePackages"
SET
  name = COALESCE(sqlc.narg('name'), name),
  gender_id = COALESCE(sqlc.narg('gender_id'), gender_id),
  timer = COALESCE(sqlc.narg('timer'), timer),
  price = COALESCE(sqlc.narg('price'), price),
  discount_price = COALESCE(sqlc.narg('discount_price'), 0),
  discount_start_time = sqlc.narg('discount_start_time'),
  discount_end_time = sqlc.narg('discount_end_time'),
  description = COALESCE(sqlc.narg('description'), ''),
  image_url = COALESCE(sqlc.narg('image_url'), ''),
  is_active = COALESCE(sqlc.narg('is_active'), is_active)
WHERE id = sqlc.arg('id')
RETURNING *;

-- name: UpsertServicePackageItem :exec
WITH upserted AS (
  INSERT INTO "ServicePackageItems" (service_package_id, service_item_id)
  SELECT $1, unnest($2::uuid[])
  ON CONFLICT (service_package_id, service_item_id) DO NOTHING
)
DELETE FROM "ServicePackageItems" AS csi
WHERE csi.service_package_id = $1
  AND csi.service_item_id NOT IN (
    SELECT unnest($2::uuid[])
);
