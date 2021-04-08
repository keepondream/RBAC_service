package server

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"database/sql"

	pgadapter "github.com/casbin/casbin-pg-adapter"
	"github.com/casbin/casbin/v2"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/keepondream/RBAC_service/handle"
	"github.com/keepondream/RBAC_service/utils"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type RBACServer struct {
	Handle *handle.Handle
	Router *gin.Engine
}

func NewDB(config *utils.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DB_USER, config.DB_PWD, config.DB_HOST, config.DB_PORT, config.DB_NAME))
	if err != nil {
		log.Fatal("connect postgres failed err is ", err)
	}

	return db, nil
}

func NewRBACServer(config *utils.Config, db *sql.DB, handle *handle.Handle, router *gin.Engine) (rbac *RBACServer) {
	log.Println("new rbac server ....")

	rbac = &RBACServer{
		Handle: handle,
		Router: router,
	}

	rbac.InitMigrate(db)

	return
}

func NewEnforcer(config *utils.Config) (e *casbin.Enforcer) {
	log.Println("make enforcer ....")
	a, err := pgadapter.NewAdapterByDB(pg.Connect(
		&pg.Options{
			User:     config.DB_USER,
			Password: config.DB_PWD,
			Database: config.DB_NAME,
			Addr:     fmt.Sprintf("%s:%s", config.DB_HOST, config.DB_PORT),
		},
	),
		pgadapter.WithTableName("rbac_casbin_rules"),
	)
	if err != nil {
		log.Fatal("new pgadapter failed err is ", err)
	}

	e, err = casbin.NewEnforcer(config.ENFORCER_MODEL_FILE, a)
	if err != nil {
		log.Fatal("new enforcer failed err is ", err)
	}

	e.EnableAutoSave(true)

	e.LoadPolicy()

	return
}

func NewLogger(config *utils.Config) (logger *logrus.Logger) {
	log.Println("make logger ....")
	logger = logrus.New()

	level := logrus.DebugLevel
	if config.LOG_LEVEL != "" {
		var err error
		level, err = logrus.ParseLevel(config.LOG_LEVEL)
		if err != nil {
			level = logrus.DebugLevel
		}
	}

	// Only log the warning severity or above.
	logger.SetLevel(level)

	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logger.SetOutput(os.Stdout)

	if config.LOG_FILE != "" {
		dir := filepath.Dir(config.LOG_FILE)
		_, err := os.Stat(dir)
		if err != nil && os.IsNotExist(err) {
			os.MkdirAll(dir, os.ModePerm)
		}

		file, err := os.OpenFile(config.LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err == nil {
			logger.SetOutput(file)
		}
	}

	return
}

func (rbac *RBACServer) InitMigrate(db *sql.DB) {
	log.Println("start init migrate ...")
	driver, err := postgres.WithInstance(db, &postgres.Config{
		MigrationsTable: rbac.Handle.Config.MIGRATIONS_TABLE,
	})
	if err != nil {
		log.Fatal("postgres driver failed err is ", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal("postgres instance failed err is ", err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("migrate up failed err is ", err)
	}

}

func (rbac *RBACServer) Start() {
	err := endless.ListenAndServe(":"+rbac.Handle.Config.SERVER_PORT, rbac.Router)

	if err != nil {
		log.Fatal(err)
	}
}
