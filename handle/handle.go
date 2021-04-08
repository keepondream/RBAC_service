package handle

import (
	"github.com/casbin/casbin/v2"
	"github.com/keepondream/RBAC_service/service"
	"github.com/keepondream/RBAC_service/store"
	"github.com/keepondream/RBAC_service/utils"
	"github.com/sirupsen/logrus"
)

type Handle struct {
	Config   *utils.Config
	Enforcer *casbin.Enforcer
	Logger   *logrus.Logger
	Store    *store.Store
	Service  *service.Service
}

func NewHandle(config *utils.Config, e *casbin.Enforcer, logger *logrus.Logger, store *store.Store, service *service.Service) *Handle {
	return &Handle{
		Config:   config,
		Enforcer: e,
		Logger:   logger,
		Store:    store,
		Service:  service,
	}
}
