.PHONY:

DB_USER=root
DB_PWD=root
DB_HOST=192.168.102.240
DB_PORT=8588
DB_NAME=root
SERVER_NAME=RBAC_service
SERVER_PORT=9010
ADMINER_PORT=8589
SWAGGER=true

postgres:
	docker ps | grep db-$(SERVER_NAME) || docker run --rm -itd -e POSTGRES_PASSWORD=$(DB_PWD) -e POSTGRES_USER=$(DB_USER) -e POSTGRES_DB=$(DB_NAME) -p $(DB_PORT):5432 --name db-$(SERVER_NAME) postgres 

adminer:
	docker ps | grep adminer-$(SERVER_NAME) || docker run --rm -itd -p $(ADMINER_PORT):8080 --name adminer-$(SERVER_NAME) adminer

browse: postgres adminer
	open 'http://$(DB_HOST):$(ADMINER_PORT)?pgsql=$(DB_HOST):$(DB_PORT)&username=$(DB_USER)&db=$(DB_NAME)&ns=public'

check_sqlc:
	@which sqlc || go get github.com/kyleconroy/sqlc/cmd/sqlc && go mod tidy 
	@sqlc version

sqlc_init: check_sqlc
	sqlc init -f ./sqlc.yaml

sqlc_generate: check_sqlc
	sqlc generate

check_swag:
	@which swag || go get github.com/swaggo/swag/cmd/swag && go mod tidy

swag: check_swag
	swag init -d rbac/rbac_api --parseDependency --parseInternal -g index.go -o ./docs/rbac

# create_table=init_id_generator
create_table=check_casbin_rules_table
# brew install golang-migrate for Mac cli used
migrate:
	migrate create -ext sql -dir migrations/rbac $(create_table)

check_wire:
	@which wire || go get github.com/google/wire/cmd/wire && go mod tidy

wire: check_wire
	wire ./wire

run: postgres adminer wire
	DB_USER=$(DB_USER) DB_PWD=$(DB_PWD) DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) DB_NAME=$(DB_NAME) SERVER_PORT=$(SERVER_PORT) SWAGGER=$(SWAGGER)  go run ./cmd/rbac