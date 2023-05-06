-- name: CreateStore :one

INSERT INTO store (name_store,manager_id, employee_id, location, status)
VALUES ($1,
        $2,
        $3,
        $4,
        $5
        ) RETURNING *;