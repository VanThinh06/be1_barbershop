-- name: CreateSchedulerWork :one

INSERT INTO schedulerwork(barber_id, users_id, timerstart, timerend, service, total_price, status)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7) RETURNING *;

