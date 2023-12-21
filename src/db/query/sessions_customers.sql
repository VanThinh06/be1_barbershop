-- name: CreateSessionsCustomer :one
INSERT INTO "SessionsCustomers" (
                               id,
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
        $9,
        ST_GeographyFromText('POINT(' || sqlc.arg(longitude)::float8 || ' ' || sqlc.arg(latitude)::float8 || ')')
        ) RETURNING *;

-- name: GetSessionsCustomer :one
SELECT *
FROM "SessionsCustomers"
WHERE id = $1
LIMIT 1;

-- name: UpdateRefreshTokenSessionsCustomer :one
UPDATE "SessionsCustomers"
set refresh_token = $2,
    expires_at = $3
WHERE id = $1
RETURNING *;

