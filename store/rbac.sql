-- name: GetInfoByIDTenant :one
SELECT *
FROM rbac_casbin_rules
WHERE id = $1
  AND v1 = $2;
-- name: ListBySignTenant :many
SELECT *
FROM rbac_casbin_rules
WHERE v0 = $1
  AND v1 = $2
LIMIT $3 OFFSET $4;
-- name: TotalBySignTenant :one
SELECT count(*)
FROM rbac_casbin_rules
WHERE v0 = $1
  AND v1 = $2;