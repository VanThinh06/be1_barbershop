-- name: CreateSessionsCustomer :one
INSERT INTO "SessionsCustomer" (
                               id,
                               customer_id,
                               refresh_token,
                               user_agent,
                               client_ip,
                               fcm_device,
                               is_blocked,
                               coordinates,
                               expires_at,
                               timezone,
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
        $10,
        ) RETURNING *;

-- name: GetSessionsCustomer :one
SELECT *
FROM "SessionsCustomer"
WHERE id = $1
LIMIT 1;

-- name: UpdateRefreshTokenSessionsCustomer :one
UPDATE "SessionsCustomer"
set refresh_token = $2,
    expires_at = $3
WHERE id = $1
RETURNING *;

