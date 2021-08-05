package service

import (
	"context"

	"github.com/keepondream/RBAC_service/internal/rbac/adapters/repo"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
)

type Route struct {
	*Service
}

func NewRoute(s *Service) *Route {
	return &Route{
		Service: s,
	}
}

type Router interface {
	GetById(ctx context.Context, id string) (*ports.RouteDetail, error)
	List(ctx context.Context, params ports.GetRoutesParams) (*ports.RouteList, error)
	DeleteById(ctx context.Context, id string) error
	Update(ctx context.Context, params ports.PatchRoutesIdJSONBody, id string) (*ports.RouteDetail, error)
	Create(ctx context.Context, params ports.PostRoutesJSONBody) (*ports.RouteDetail, error)
}

func (s *Route) GetById(ctx context.Context, id string) (*ports.RouteDetail, error) {
	r := repo.NewRoute(s.Repo)
	return r.GetById(ctx, id)
}

func (s *Route) List(ctx context.Context, params ports.GetRoutesParams) (*ports.RouteList, error) {
	r := repo.NewRoute(s.Repo)
	return r.List(ctx, params)
}

func (s *Route) DeleteById(ctx context.Context, id string) error {
	r := repo.NewRoute(s.Repo)
	return r.DeleteById(ctx, id)
}

func (s *Route) Update(ctx context.Context, params ports.PatchRoutesIdJSONBody, id string) (*ports.RouteDetail, error) {
	r := repo.NewRoute(s.Repo)
	return r.Update(ctx, params, id)
}

func (s *Route) Create(ctx context.Context, params ports.PostRoutesJSONBody) (*ports.RouteDetail, error) {
	r := repo.NewRoute(s.Repo)
	return r.Create(ctx, params)
}
