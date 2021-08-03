dev:
	cd ./internal/rbac && go run .

rbacWire:
	cd ./internal/rbac/app && wire



rbacEnt:
	cd ./internal/rbac && ent init Menu Route