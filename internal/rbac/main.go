package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"github.com/keepondream/RBAC_service/internal/common/server"
	"github.com/keepondream/RBAC_service/internal/rbac/app"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
)

func main() {
	fmt.Println("starting rbac server ......")

	app, err := app.NewApplication()
	if err != nil {
		log.Panic(err)
	}

	// 实例swagger规则
	swagger, err := ports.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}
	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	server.RunHTTPServer("/", func(router chi.Router) http.Handler {
		// 注入swagger规则校验中间件
		router.Use(middleware.OapiRequestValidator(swagger))
		return ports.HandlerFromMux(app.HttpServer, router)
	})

}
