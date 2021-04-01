//+build wireinject

package wire

import (
	google_wire "github.com/google/wire"
	"github.com/keepondream/RBAC_service/handle"
	"github.com/keepondream/RBAC_service/route"
	"github.com/keepondream/RBAC_service/server"
	"github.com/keepondream/RBAC_service/store"
	"github.com/keepondream/RBAC_service/utils"
)

func InitRBACServer(configFile string) (*server.RBACServer, error) {
	google_wire.Build(
		utils.LoadConfig,
		server.NewDB,
		store.NewStore,
		server.NewRBACServer,
		server.NewEnforcer,
		server.NewLogger,
		handle.NewHandle,
		route.NewRoute,
	)
	return &server.RBACServer{}, nil
}
