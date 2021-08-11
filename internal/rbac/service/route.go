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
	GetById(ctx context.Context, id string) (*ports.RouteInfoResponse, error)
	List(ctx context.Context, params ports.GetRoutesParams) (*ports.RouteListResponse, error)
	DeleteById(ctx context.Context, id string) error
	Update(ctx context.Context, params ports.PatchRoutesIdJSONBody, id string) (*ports.RouteInfoResponse, error)
	Create(ctx context.Context, params ports.PostRoutesJSONBody) (*ports.RouteInfoResponse, error)
	IsUnique(ctx context.Context, tenant string, uri string, method string) error
}

func (s *Route) GetById(ctx context.Context, id string) (*ports.RouteInfoResponse, error) {
	r := repo.NewRoute(s.Repo)
	return r.GetById(ctx, id)
}

func (s *Route) List(ctx context.Context, params ports.GetRoutesParams) (*ports.RouteListResponse, error) {
	r := repo.NewRoute(s.Repo)
	return r.List(ctx, params)
}

func (s *Route) DeleteById(ctx context.Context, id string) error {
	r := repo.NewRoute(s.Repo)
	res, err := r.GetById(ctx, id)
	if err != nil {
		return err
	}
	err = r.DeleteById(ctx, id)
	if err != nil {
		return err
	}

	// 同步路由对应的权限策略
	s.SyncCasbinForRoute(ctx, string(res.Tenant), res.Uri, string(res.Method), res.Id)

	return nil
}

func (s *Route) Update(ctx context.Context, params ports.PatchRoutesIdJSONBody, id string) (*ports.RouteInfoResponse, error) {
	r := repo.NewRoute(s.Repo)
	res, err := r.Update(ctx, params, id)
	if err != nil {
		return nil, err
	}

	// 同步路由对应的权限策略
	s.SyncCasbinForRoute(ctx, string(res.Tenant), res.Uri, string(res.Method), res.Id)

	return res, nil
}

func (s *Route) Create(ctx context.Context, params ports.PostRoutesJSONBody) (*ports.RouteInfoResponse, error) {
	r := repo.NewRoute(s.Repo)
	res, err := r.Create(ctx, params)
	if err != nil {
		return nil, err
	}

	// 同步路由对应的权限策略
	s.SyncCasbinForRoute(ctx, string(res.Tenant), res.Uri, string(res.Method), res.Id)

	return res, nil
}

func (s *Route) IsUnique(ctx context.Context, tenant string, uri string, method string) error {
	r := repo.NewRoute(s.Repo)
	return r.IsUnique(ctx, tenant, uri, method)
}
