dev:
	cd ./internal/rbac && go run .

rbacWire: rbacEnt yapi
	cd ./internal/rbac/app && wire

# 最终抽象: 路由表, 权限表, 节点表(可以->角色/菜单/页面元素等等), 分组表(可以->角色组/菜单组/页面元素组等等), 用户表(只记录用户唯一标识和域以及是否为超级管理员)
# rbacFirstEntSchema:
# 	cd ./internal/rbac/adapters && ent init Route Permission

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

# http://yapi.smart-xwork.cn 
yapi:
	yapi import