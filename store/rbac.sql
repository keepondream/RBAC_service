-- name: GetInfoByIDTenant :one
SELECT *
FROM rbac_casbin_rules
WHERE id = $1
  AND v1 = $2;