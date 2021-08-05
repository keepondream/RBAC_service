dev:
	cd ./internal/rbac && go run .

rbacWire: rbacEnt
	cd ./internal/rbac/app && wire

# 路由表, 权限表, 权限路由关系表(多对多), 分组表(角色:role, 角色组:role_group, 用户:user, 用户组:user_group, 菜单:menu, 菜单组:menu_group, 权限组:permission_group, 页面元素:element, 页面元素组:element_group 等等... ), 权限分组关系表(多对多)
rbacFirstEntSchema:
	cd ./internal/rbac/adapters && ent init Route Permission

rbacEnt:
	cd ./internal/rbac/adapters && go generate ./ent

openapi_http:
	oapi-codegen -generate types -o internal/rbac/ports/openapi_types.gen.go -package ports api/openapi/reference/rbac-api.json
	oapi-codegen -generate chi-server -o internal/rbac/ports/openapi_api.gen.go -package ports api/openapi/reference/rbac-api.json
	oapi-codegen -generate types -o internal/common/client/rbac/openapi_types.gen.go -package rbac api/openapi/reference/rbac-api.json
	oapi-codegen -generate client -o internal/common/client/rbac/openapi_client.gen.go -package rbac api/openapi/reference/rbac-api.json

local: rbacWire
	docker-compose down
	docker-compose up