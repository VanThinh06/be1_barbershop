-- name: CreateCustomer :one
INSERT INTO "Customers" (
   "name",
    email,
    phone,
    hashed_password,
    is_social_auth
  )
VALUES (
    $1,
    $2,
    $3,
    $4,
    sqlc.arg(is_social_auth)::bool
  )
RETURNING *;

-- name: GetUserCustomer :one
SELECT *
FROM "Customers"
WHERE
    (
        (sqlc.arg(type_username)::varchar = 'email' AND email = $1)
        OR
        (sqlc.arg(type_username)::varchar = 'phone' AND phone = $1)
    )
LIMIT 1;

-- name: GetCustomer :one
SELECT *
  FROM "Customers"
  WHERE "id" = sqlc.arg('id');

-- name: UpdateCustomer :one
UPDATE "Customers"
set name = coalesce(sqlc.narg('name'), name),
  email = coalesce(sqlc.narg('email'), email),
  phone = coalesce(sqlc.narg('phone'), phone),
  gender_id = coalesce(sqlc.narg('gender_id'), gender_id),
  avatar = coalesce(sqlc.narg('avatar'), avatar),
  "update_at" = NOW()
  WHERE "id" = sqlc.arg('id')
RETURNING *;


-- name: ChangePasswordCustomer :one
UPDATE "Customers" 
set   
  "hashed_password" = sqlc.arg('hashed_password')::varchar(100),
  "password_changed_at" = NOW()
WHERE "id" = sqlc.arg('id')
RETURNING *;