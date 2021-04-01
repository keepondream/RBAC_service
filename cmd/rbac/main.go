package main

import (
	"log"

	"github.com/keepondream/RBAC_service/wire"
)

func main() {
	log.Println("start server ....")

	server, err := wire.InitRBACServer("./app.env")
	if err != nil {
		log.Fatal("init rbac server failed err is ", err)
	}

	server.Start()

}
