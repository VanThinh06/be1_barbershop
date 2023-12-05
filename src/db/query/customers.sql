-- name: CreateCustomer :one
INSERT INTO "Customers" (
   "name",
    email,
    phone,
    gender,
    hashed_password,
    is_social_auth
  )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    sqlc.arg(is_social_auth)::bool
  )
RETURNING *;

-- name: GetContactCustomer :one
SELECT *
FROM "Customers"
WHERE
    (
        (sqlc.arg(type_username)::varchar = 'email' AND email = $1)
        OR
        (sqlc.arg(type_username)::varchar = 'phone' AND phone = $1)
    )
LIMIT 1;

-- name: UpdateCustomer :one
UPDATE "Customers"
set name = coalesce(sqlc.narg('name'), name),
  email = coalesce(sqlc.narg('email'), email),
  phone = coalesce(sqlc.narg('phone'), phone),
  gender = coalesce(sqlc.narg('gender'), gender),
  avatar = coalesce(sqlc.narg('avatar'), avatar),
  "hashed_password" = coalesce(sqlc.narg('hashed_password'), hashed_password),
  "password_changed_at" = coalesce(
    sqlc.narg('password_changed_at'),
    password_changed_at
  ),
  "update_at" = sqlc.arg('update_at')
  WHERE "customer_id" = sqlc.arg('customer_id')
RETURNING *;
