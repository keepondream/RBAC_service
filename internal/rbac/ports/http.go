package ports

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/keepondream/RBAC_service/internal/common/utils"
)

type HttpServer struct {
	Validate          *validator.Validate
	RouteService      Router
	PermissionService Permissioner
}

// NewHttpServer 注入依赖服务
func NewHttpServer(
	routeService Router,
	permissionService Permissioner,
) *HttpServer {
	return &HttpServer{
		Validate:          utils.NewValidate(),
		RouteService:      routeService,
		PermissionService: permissionService,
	}
}

func (h *HttpServer) Test() {

	fmt.Println("test  in http of ports")
}
