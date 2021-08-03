package service

import "github.com/keepondream/RBAC_service/internal/rbac/adapters/cache"

type Service struct {
	Cache       *cache.Cache
	MenuService MenuService
}

func NewService() *Service {
	return &Service{}
}
