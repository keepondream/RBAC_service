package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/keepondream/RBAC_service/store"
	"github.com/keepondream/RBAC_service/utils"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Config   *utils.Config
	Enforcer *casbin.Enforcer
	Logger   *logrus.Logger
	Store    *store.Store
}

func NewService(config *utils.Config, e *casbin.Enforcer, logger *logrus.Logger, store *store.Store) *Service {
	return &Service{
		Config:   config,
		Enforcer: e,
		Logger:   logger,
		Store:    store,
	}
}
