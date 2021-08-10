package service

import (
	"context"

	"github.com/keepondream/RBAC_service/internal/rbac/adapters/repo"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
)

type User struct {
	*Service
}

func NewUser(s *Service) *User {
	return &User{
		Service: s,
	}
}

type Userer interface {
	Create(ctx context.Context, params ports.PostUsersJSONBody) (*ports.UserInfoResponse, error)
	IsUnique(ctx context.Context, tenant, uuid string) error
	GetById(ctx context.Context, id string) (*ports.UserInfoResponse, error)
	GetByUuid(ctx context.Context, tenant, uuid string) (*ports.UserInfoResponse, error)
	DeleteByUuid(ctx context.Context, tenant, uuid string) error
	List(ctx context.Context, params ports.GetUsersParams) (*ports.UserListResponse, error)
	Update(ctx context.Context, params ports.PatchUsersUuidTenantJSONBody, tenant, uuid string) (*ports.UserInfoResponse, error)
}

func (s *User) Create(ctx context.Context, params ports.PostUsersJSONBody) (*ports.UserInfoResponse, error) {
	r := repo.NewUser(s.Repo)
	res, err := r.Create(ctx, params)
	if err != nil {
		return nil, err
	}

	s.SyncCasbinForUser(ctx, res.Uuid, string(res.Tenant))

	return res, nil
}

func (s *User) IsUnique(ctx context.Context, tenant string, uuid string) error {
	r := repo.NewUser(s.Repo)
	return r.IsUnique(ctx, tenant, uuid)
}

func (s *User) GetById(ctx context.Context, id string) (*ports.UserInfoResponse, error) {
	r := repo.NewUser(s.Repo)
	return r.GetById(ctx, id)
}

func (s *User) GetByUuid(ctx context.Context, tenant string, uuid string) (*ports.UserInfoResponse, error) {
	r := repo.NewUser(s.Repo)
	return r.GetByUuid(ctx, tenant, uuid)
}

func (s *User) DeleteByUuid(ctx context.Context, tenant string, uuid string) error {
	r := repo.NewUser(s.Repo)
	res, err := r.GetByUuid(ctx, tenant, uuid)
	if err != nil {
		return err
	}

	err = r.DeleteByUuid(ctx, tenant, uuid)
	if err != nil {
		return err
	}

	s.SyncCasbinForUser(ctx, res.Uuid, string(res.Tenant))

	return nil
}

func (s *User) List(ctx context.Context, params ports.GetUsersParams) (*ports.UserListResponse, error) {
	r := repo.NewUser(s.Repo)
	return r.List(ctx, params)
}

func (s *User) Update(ctx context.Context, params ports.PatchUsersUuidTenantJSONBody, tenant string, uuid string) (*ports.UserInfoResponse, error) {
	r := repo.NewUser(s.Repo)
	res, err := r.Update(ctx, params, tenant, uuid)
	if err != nil {
		return nil, err
	}

	s.SyncCasbinForUser(ctx, res.Uuid, string(res.Tenant))

	return res, nil
}
