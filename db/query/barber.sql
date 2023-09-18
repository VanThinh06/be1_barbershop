-- name: CreateBarber :one
INSERT INTO barber(username, full_name, email, hashed_password, role, store_work, type_barber)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7
        ) RETURNING *;

-- name: GetBarber :one

SELECT *
FROM barber
WHERE username = $1
LIMIT 1
;

-- name: UpdateBarber :one

UPDATE "barber"
set 
  status = $2,
  full_name = $3,
  "role" = $4,
  store_work = $5,
  type_barber = $6,
  email = $7,
  update_at = $8,
  avatar = $9
WHERE username = $1
RETURNING *;