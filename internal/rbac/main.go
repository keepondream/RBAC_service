package main

import (
	"fmt"
	"log"

	"github.com/keepondream/RBAC_service/internal/rbac/app"
)

func main() {
	fmt.Println("starting rbac server ......")

	app, err := app.NewApplication()
	if err != nil {
		log.Panic(err)
	}

	app.HttpServer.Test()
}
