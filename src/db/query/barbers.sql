-- name: CreateBarber :one
INSERT INTO "Barbers" (
    gender_id,
    email,
    phone,
    hashed_password,
    nick_name,
    full_name
  )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
  )
RETURNING *;

-- name: CreateBarberEmployee :one
INSERT INTO "Barbers" (
    phone,
    hashed_password,
    full_name,
    nick_name
  )
VALUES (
    $1,
    $2,
    $3,
    $4
  )
RETURNING *;

-- name: GetBarber :one
SELECT
  b.*, 
  br.*
FROM
  "Barbers" b
JOIN
  "BarberRoles" br ON b."id" = br."barber_id"
WHERE
  b."id" = $1
  AND br."barber_shop_id" = $2;

-- name: GetBarberEmployee :one
SELECT
  b.*, 
  br.*
FROM
  "Barbers" b
LEFT JOIN
  "BarberRoles" br ON b."id" = br."barber_id"
WHERE
  b."id" = $1
  AND (br."barber_shop_id" = $2 OR br."barber_shop_id" IS NULL);

-- name: GetUserBarber :one
SELECT 
  b.*
FROM "Barbers" b
WHERE  (
        (sqlc.arg(type_username)::varchar = 'email' AND email = $1)
        OR
        (sqlc.arg(type_username)::varchar = 'phone' AND phone = $1)
    );


-- name: ListEmployeesAdmin :many
SELECT *,
       (SELECT COUNT(*) FROM "Barbers" b
        JOIN "BarberRoles" br ON b."id" = br."barber_id"
        JOIN "Roles" r ON br."role_id" = r."id"
        WHERE br."barber_shop_id" = $1
          AND r."type" IN ('Staff', 'Management')) AS total_employees
FROM "Barbers" b
JOIN "BarberRoles" br ON b."id" = br."barber_id"
JOIN "Roles" r ON br."role_id" = r."id"
WHERE br."barber_shop_id" = $1
  AND r."type" IN ('Staff', 'Management')
ORDER BY br."role_id"
LIMIT $2 OFFSET $3;

-- name: UpdateBarber :one
UPDATE "Barbers"
SET 
  gender_id = coalesce(sqlc.narg('gender_id'), gender_id),
  email = coalesce(sqlc.narg('email'), email),
  phone = coalesce(sqlc.narg('phone'), phone),
  nick_name = coalesce(sqlc.narg('nick_name'), nick_name),
  full_name = coalesce(sqlc.narg('full_name'), full_name),
  haircut = coalesce(sqlc.narg('haircut'), haircut),
  work_status = coalesce(sqlc.narg('work_status'), work_status),
  avatar_url = coalesce(sqlc.narg('avatar_url'), avatar_url)
  WHERE "id" = $1
RETURNING *;

-- name: DeleteBarber :exec
DELETE FROM "Barbers"
WHERE "id" = $1;
