-- name: CreateSessionBarber :one
INSERT INTO "sessions_barber" (
                               id,
                               username,
                               refresh_token,
                               user_agent,
                               client_ip,
                               fcm_device,
                               is_blocked,
                               expires_at)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8
        ) RETURNING *;

-- name: GetSession :one
SELECT *
FROM "sessions_barber"
WHERE id = $1
LIMIT 1;

-- name: UpdateRefreshToken :one
UPDATE "sessions_barber"
set refresh_token = $2,
    expires_at = $3
WHERE id = $1
RETURNING *;

