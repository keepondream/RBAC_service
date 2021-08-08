//+build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/cache"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/repo"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
	"github.com/keepondream/RBAC_service/internal/rbac/service"
)

var NewRoute = wire.NewSet(service.NewRoute, wire.Bind(new(ports.Router), new(*service.Route)))
var NewPermission = wire.NewSet(service.NewPermission, wire.Bind(new(ports.Permissioner), new(*service.Permission)))
var NewNode = wire.NewSet(service.NewNode, wire.Bind(new(ports.Noder), new(*service.Node)))

func NewApplication() (*App, error) {
	wire.Build(
		NewApp,
		NewDB,
		NewEntClient,
		NewEntAdapter,
		NewEnforcer,
		service.NewService,
		cache.NewCache,
		ports.NewHttpServer,
		repo.NewRepo,
		NewRoute,
		NewPermission,
		NewNode,
	)

	return &App{}, nil
}
