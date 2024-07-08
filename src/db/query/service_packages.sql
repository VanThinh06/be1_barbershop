
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


-- name: GetServicePackage :one
SELECT * FROM "view_service_packages"
WHERE id = $1;


-- name: ListServicePackages :many
SELECT * FROM "view_service_packages"
WHERE barber_shop_id = $1;


-- name: UpdateServicePackage :one
UPDATE "ServicePackages"
SET
  name = COALESCE($2, name),
  gender_id = COALESCE($3, gender_id),
  timer = COALESCE($4, timer),
  price = COALESCE($5, price),
  discount_price = COALESCE($6, 0),
  discount_start_time = $7,
  discount_end_time = $8,
  description = COALESCE($9, ''),
  image_url = COALESCE($10, ''),
  is_active = COALESCE($11, is_active)
WHERE id = $1
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
