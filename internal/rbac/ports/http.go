package ports

import (
	"github.com/go-playground/validator/v10"
	"github.com/keepondream/RBAC_service/internal/common/utils"
)

type HttpServer struct {
	Validate *validator.Validate
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		Validate: utils.NewValidate(),
	}
}
