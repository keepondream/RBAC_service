package service

import (
	"context"

	"github.com/keepondream/RBAC_service/internal/rbac/adapters/repo"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
)

type Group struct {
	*Service
}

func NewGroup(s *Service) *Group {
	return &Group{
		Service: s,
	}
}

type Grouper interface {
	Create(ctx context.Context, params ports.PostGroupsJSONBody) (*ports.GroupInfoResponse, error)
	IsUnique(ctx context.Context, tenant, name, group_type string) error
	GetById(ctx context.Context, id string) (*ports.GroupInfoResponse, error)
	DeleteById(ctx context.Context, id string) error
	List(ctx context.Context, params ports.GetGroupsParams) (*ports.GroupListResponse, error)
	Update(ctx context.Context, params ports.PatchGroupsIdJSONBody, id string) (*ports.GroupInfoResponse, error)
}

func (s *Group) Create(ctx context.Context, params ports.PostGroupsJSONBody) (*ports.GroupInfoResponse, error) {
	r := repo.NewGroup(s.Repo)
	return r.Create(ctx, params)
}

func (s *Group) IsUnique(ctx context.Context, tenant string, name string, group_type string) error {
	r := repo.NewGroup(s.Repo)
	return r.IsUnique(ctx, tenant, name, group_type)
}

func (s *Group) GetById(ctx context.Context, id string) (*ports.GroupInfoResponse, error) {
	r := repo.NewGroup(s.Repo)
	return r.GetById(ctx, id)
}

func (s *Group) DeleteById(ctx context.Context, id string) error {
	r := repo.NewGroup(s.Repo)
	return r.DeleteById(ctx, id)
}

func (s *Group) List(ctx context.Context, params ports.GetGroupsParams) (*ports.GroupListResponse, error) {
	r := repo.NewGroup(s.Repo)
	return r.List(ctx, params)
}

func (s *Group) Update(ctx context.Context, params ports.PatchGroupsIdJSONBody, id string) (*ports.GroupInfoResponse, error) {
	r := repo.NewGroup(s.Repo)
	return r.Update(ctx, params, id)
}
