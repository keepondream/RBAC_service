// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package app

import (
	"github.com/google/wire"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/cache"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/repo"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
	"github.com/keepondream/RBAC_service/internal/rbac/service"
)

// Injectors from wire.go:

func NewApplication() (*App, error) {
	cacheCache := cache.NewCache()
	db := NewDB()
	client := NewEntClient(db)
	repoRepo := repo.NewRepo(client)
	adapter := NewEntAdapter(db)
	enforcer := NewEnforcer(adapter)
	serviceService := service.NewService(cacheCache, client, repoRepo, enforcer)
	route := service.NewRoute(serviceService)
	permission := service.NewPermission(serviceService)
	node := service.NewNode(serviceService)
	group := service.NewGroup(serviceService)
	httpServer := ports.NewHttpServer(route, permission, node, group)
	app := NewApp(httpServer)
	return app, nil
}

// wire.go:

var NewRoute = wire.NewSet(service.NewRoute, wire.Bind(new(ports.Router), new(*service.Route)))

var NewPermission = wire.NewSet(service.NewPermission, wire.Bind(new(ports.Permissioner), new(*service.Permission)))

var NewNode = wire.NewSet(service.NewNode, wire.Bind(new(ports.Noder), new(*service.Node)))

var NewGroup = wire.NewSet(service.NewGroup, wire.Bind(new(ports.Grouper), new(*service.Group)))
