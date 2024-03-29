-- name: CreateSessionsCustomer :one
INSERT INTO "SessionsCustomer" (
                               customer_id,
                               refresh_token,
                               user_agent,
                               client_ip,
                               fcm_device,
                               is_blocked,
                               expires_at,
                               coordinates
                               )
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        now() + interval '30 minutes',
        ST_GeographyFromText('POINT(' || sqlc.arg(longitude)::float8 || ' ' || sqlc.arg(latitude)::float8 || ')')
        ) RETURNING *;

-- name: GetSessionsCustomer :one
SELECT *
FROM "SessionsCustomer"
WHERE id = $1;

-- name: UpdateSessionsCustomer :exec
UPDATE "SessionsCustomer"
SET refresh_token = $2,
 "expires_at" = now() + interval '30 minutes'
WHERE id = $1
RETURNING * ;

