-- name: CheckBarberRolePermission :one
SELECT EXISTS (
    SELECT 1
    FROM "BarberRoles" br
    JOIN "RolePermissions" rp ON br."role_id" = rp."role_id"
    JOIN "Permissions" p ON rp."permission_id" = p."id"
    WHERE br."barber_id" = $1 AND br."barber_shop_id" = $2 AND p."id" = $3
) AS "has_permission";