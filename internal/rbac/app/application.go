package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc/integration/idtype/ent"
	"github.com/google/wire"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/repo"
	"github.com/keepondream/RBAC_service/internal/rbac/service"
)

type App struct {
	Service *service.Service
}

func NewApp() *App {
	return &App{}
}

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Panicf("connect postgres failed err : %v \n", err)
	}

	return db
}

func NewEntClient(db *sql.DB) *ent.Client {
	// 从db变量中构造一个ent.Driver对象。
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(drv))

	// 数据库迁移
	if err := client.Debug().Schema.Create(
		context.Background(),
		schema.WithDropColumn(true), // 迁移可以删除字段
		schema.WithDropIndex(true),  // 迁移可以删除索引
	); err != nil {
		log.Panicf("database migrate failed err:%v \n", err)
	}

	return client
}

var MenuRepoSet = wire.NewSet(repo.NewMenuRepo, wire.Bind(new(service.MenuService), new(*repo.MenuRepo)))
