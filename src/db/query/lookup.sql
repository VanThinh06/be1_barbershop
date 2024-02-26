-- name: GetProvinces :many
SELECT *
FROM "Provinces";

-- name: GetDistricts :many
SELECT *
FROM "Districts"
WHERE province_id = $1;


-- name: GetWards :many
SELECT *
FROM "Wards"
WHERE district_id = $1;

