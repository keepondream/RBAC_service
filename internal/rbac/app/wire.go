//+build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/cache"
	"github.com/keepondream/RBAC_service/internal/rbac/service"
)

func NewApplication() (*App, error) {

	wire.Build(
		NewApp,
		NewDB,
		NewEntClient,
		service.NewService,
		cache.NewCache,
		MenuRepoSet,
	)

	return &App{}, nil
}
