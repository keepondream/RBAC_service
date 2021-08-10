package service

import (
	"context"

	"github.com/keepondream/RBAC_service/internal/rbac/adapters/repo"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
)

type Permission struct {
	*Service
}

func NewPermission(s *Service) *Permission {
	return &Permission{
		Service: s,
	}
}

type Permissioner interface {
	Create(ctx context.Context, params ports.PostPermissionsJSONBody) (*ports.PermissionInfoResponse, error)
	IsUnique(ctx context.Context, tenant string, name string) error
	GetById(ctx context.Context, id string) (*ports.PermissionInfoResponse, error)
	List(ctx context.Context, params ports.GetPermissionsParams) (*ports.RouteListResponse, error)
	DeleteById(ctx context.Context, id string) error
	Update(ctx context.Context, params ports.PatchPermissionsIdJSONBody, id string) (*ports.PermissionInfoResponse, error)
}

func (s *Permission) Create(ctx context.Context, params ports.PostPermissionsJSONBody) (*ports.PermissionInfoResponse, error) {
	r := repo.NewPermission(s.Repo)
	res, err := r.Create(ctx, params)
	if err != nil {
		return nil, err
	}

	// 同步权限所有的casbin策略
	s.SyncPolicyForPermission(ctx, res.Id)

	return res, nil
}

func (s *Permission) IsUnique(ctx context.Context, tenant string, name string) error {
	r := repo.NewPermission(s.Repo)
	return r.IsUnique(ctx, tenant, name)
}

func (s *Permission) GetById(ctx context.Context, id string) (*ports.PermissionInfoResponse, error) {
	r := repo.NewPermission(s.Repo)
	return r.GetById(ctx, id)
}

func (s *Permission) List(ctx context.Context, params ports.GetPermissionsParams) (*ports.PermissionListResponse, error) {
	r := repo.NewPermission(s.Repo)
	return r.List(ctx, params)
}

func (s *Permission) DeleteById(ctx context.Context, id string) error {
	r := repo.NewPermission(s.Repo)
	err := r.DeleteById(ctx, id)
	if err != nil {
		return err
	}

	// 同步权限所有的casbin策略
	s.SyncPolicyForPermission(ctx, id)

	return nil
}

func (s *Permission) Update(ctx context.Context, params ports.PatchPermissionsIdJSONBody, id string) (*ports.PermissionInfoResponse, error) {
	r := repo.NewPermission(s.Repo)
	res, err := r.Update(ctx, params, id)
	if err != nil {
		return nil, err
	}

	// 同步权限所有的casbin策略
	s.SyncPolicyForPermission(ctx, res.Id)

	return res, nil
}
