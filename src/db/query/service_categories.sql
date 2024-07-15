-- name: CreateServiceCategory :one
INSERT INTO "ServiceCategories" ("name", "barber_shop_id")
VALUES ($1, $2)
RETURNING *;

-- name: ListServiceCategories :many
SELECT 
    sc.*,
    cp.position,
    cp.visible
FROM 
    "ServiceCategories" sc
LEFT JOIN 
    "CategoryPositions" cp
ON 
    sc.id = cp.category_id 
WHERE sc."barber_shop_id" = $1
ORDER BY cp.position
;


-- name: UpdateServiceCategory :exec
UPDATE "ServiceCategories"
SET "name" = coalesce(sqlc.narg('name'), name)
WHERE "id" = $1;


-- name: DeleteServiceCategories :exec
DELETE FROM "ServiceCategories"
WHERE "id" IN (
  $1 -- Danh sách các ID
);


-- name: ListCategoryPosition :many
SELECT 
    sc.*,
    cp.position,
    cp.visible
FROM 
    "ServiceCategories" sc
LEFT JOIN 
    "CategoryPositions" cp
ON 
    sc.id = cp.category_id 
WHERE sc."barber_shop_id" = $1 OR sc."barber_shop_id" IS NULL
ORDER BY cp.position
;

-- name: UpdateCategoryPosition :exec
INSERT INTO "CategoryPositions" ("barber_shop_id", "category_id", "position", "visible")
VALUES ($1, $2, $3, $4)
ON CONFLICT ("barber_shop_id", "category_id") 
DO UPDATE SET 
    "position" = EXCLUDED."position",
    "visible" = EXCLUDED."visible";