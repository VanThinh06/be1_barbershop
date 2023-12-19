-- name: CreateChain :one
INSERT INTO "Chains" (
    owner_id,
   "name"
  )
VALUES (
    $1,
    $2
  )
RETURNING *;