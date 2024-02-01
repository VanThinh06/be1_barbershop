-- name: CreateSessionBarber :one
INSERT INTO "SessionsBarber" (
  barber_id,
  refresh_token,
  user_agent,
  client_ip,
  fcm_device,
  is_blocked,
  expires_at
)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  now() + interval '30 minutes' 
)
RETURNING *;

-- name: GetSessionBarber :one
SELECT *
FROM "SessionsBarber"
WHERE id = $1
LIMIT 1;

-- name: UpdateSessionRefreshToken :exec
UPDATE "SessionsBarber"
SET refresh_token = $2,
 "expires_at" = now() + interval '30 minutes'
WHERE id = $1;

