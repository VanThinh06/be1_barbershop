-- name: CreateSessionsCustomer :one
INSERT INTO "SessionsCustomer" (
                               customer_id,
                               refresh_token,
                               user_agent,
                               client_ip,
                               fcm_device,
                               is_blocked,
                               expires_at,
                               timezone,
                               coordinates
                               )
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        ST_GeographyFromText('POINT(' || sqlc.arg(longitude)::float8 || ' ' || sqlc.arg(latitude)::float8 || ')')
        ) RETURNING *;

-- name: GetSessionsCustomer :one
SELECT *
FROM "SessionsBarber"
WHERE id = $1
LIMIT 1;

-- name: UpdateSessionsCustomer :exec
UPDATE "SessionsBarber"
SET refresh_token = $2,
 "expires_at" = now() + interval '30 minutes'
WHERE id = $1;

