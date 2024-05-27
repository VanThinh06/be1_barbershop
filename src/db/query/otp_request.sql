-- name: InsertOTPRequest :exec
INSERT INTO "OTPRequests" (
    barber_id,
    otp) 
VALUES ($1, $2);

-- name: SelectRequestOTPNumber :one
SELECT COUNT(*)
FROM "OTPRequests"
WHERE barber_id = $1
AND date_trunc('day', requested_at) = date_trunc('day', now());


-- name: SelectOTPRequest :one
SELECT id, otp, requested_at, is_confirm
FROM "OTPRequests" 
WHERE barber_id = $1 AND otp = $2 
ORDER BY requested_at DESC 
LIMIT 1;

-- name: ConfirmOTPRequest :exec
UPDATE "OTPRequests"
SET is_confirm = $1
WHERE id = $2;
