package main

import (
	"fmt"
	"log"
	"net/http"

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

	server.RunHTTPServer("/", func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(app.HttpServer, router)
	})

}
