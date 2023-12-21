-- name: CreateSessionBarber :one
INSERT INTO "SessionsBarbers" (
    id,
                               barber_id,
                               refresh_token,
                               user_agent,
                               client_ip,
                               fcm_device,
                               is_blocked,
                               expires_at,
                               timezone
                               )
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8,
        sqlc.arg(timezone)::varchar
        ) RETURNING *;

-- name: GetSessionsBarber :one
SELECT *
FROM "SessionsBarbers"
WHERE id = $1
LIMIT 1;

-- name: UpdateRefreshTokenSessionsBarber :one
UPDATE "SessionsBarbers"
set refresh_token = $2,
    expires_at = $3
WHERE id = $1
RETURNING *;

