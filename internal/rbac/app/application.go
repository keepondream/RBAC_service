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
	"github.com/casbin/casbin/v2"
	entadapter "github.com/casbin/ent-adapter"
	entadapterent "github.com/casbin/ent-adapter/ent"

	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
)

type App struct {
	HttpServer *ports.HttpServer
}

func NewApp(httpServer *ports.HttpServer) *App {
	return &App{
		HttpServer: httpServer,
	}
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

func NewEntAdapter(db *sql.DB) *entadapter.Adapter {
	// 从db变量中构造一个entadapterent.Driver对象。
	drv := entsql.OpenDB(dialect.Postgres, db)
	client := entadapterent.NewClient(entadapterent.Driver(drv))

	adapter, err := entadapter.NewAdapterWithClient(client)
	if err != nil {
		log.Panicf("init ent adapter err:%v \n", err)
	}

	return adapter
}

func NewEnforcer(adapter *entadapter.Adapter) *casbin.Enforcer {
	casbinModelConfPath := os.Getenv("CASBIN_MODEL_CONF_PATH")
	e, err := casbin.NewEnforcer(casbinModelConfPath, adapter)
	if err != nil {
		log.Panicf("init casbin enforcer failed err:%v \n", err)
	}

	// 每次运行从据库中重新加载策略。
	e.LoadPolicy()

	// 启用自动保存,将内存策略同步至数据库
	e.EnableAutoSave(true)

	return e
}

// var MenuRepoSet = wire.NewSet(repo.NewMenuRepo, wire.Bind(new(service.MenuService), new(*repo.MenuRepo)))
