package ports

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/keepondream/RBAC_service/internal/common/utils"
)

type HttpServer struct {
	Validate     *validator.Validate
	RouteService Router
}

func NewHttpServer(routeService Router) *HttpServer {
	return &HttpServer{
		Validate:     utils.NewValidate(),
		RouteService: routeService,
	}
}

func (h *HttpServer) Test() {

	fmt.Println("test  in http of ports")
}
