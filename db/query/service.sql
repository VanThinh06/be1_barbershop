-- name: CreateService :one

INSERT INTO "service"( store_id,
                      work,
                       timer,
                       price,
                       description,
                       image)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6 ) RETURNING *;

-- name: GetService :many

SELECT * FROM "service"
WHERE store_id = $1 LIMIT 10;
