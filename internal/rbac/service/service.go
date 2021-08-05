package service

import (
	"github.com/casbin/casbin/v2"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/cache"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/repo"
)

type Service struct {
	Cache     *cache.Cache
	EntClient *ent.Client
	Repo      *repo.Repo
	CasbinE   *casbin.Enforcer
}

func NewService(cache *cache.Cache, entClient *ent.Client, repo *repo.Repo, e *casbin.Enforcer) *Service {
	return &Service{
		Cache:     cache,
		EntClient: entClient,
		Repo:      repo,
		CasbinE:   e,
	}
}
