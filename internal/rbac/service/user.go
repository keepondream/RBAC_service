package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/keepondream/RBAC_service/internal/rbac/adapters/repo"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
	"github.com/spf13/cast"
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

func (s *User) GetAllRoutes(ctx context.Context, tenant string, uuid string) (*ports.UserAllRoutesResponse, error) {
	resp := ports.UserAllRoutesResponse{
		PermissionIds: []string{},
		Policies:      []ports.ItemPolicy{},
	}
	prefix := fmt.Sprintf("%s%s", USERPREFIX, uuid)
	data, err := s.GetAllPolicyForPrefix(prefix, tenant)
	if err != nil {
		return nil, err
	}

	// 用于判断permissionids是否重复,类似去重功能
	set := make(map[string]struct{})

	for _, policy := range data {
		resp.Policies = append(resp.Policies, ports.ItemPolicy{
			PermissionId: policy[0],
			Tenant:       ports.ItemTenant(policy[1]),
			Uri:          policy[2],
			Method:       ports.ItemMethod(policy[3]),
		})

		if _, ok := set[policy[0]]; !ok {
			set[policy[0]] = struct{}{}
			resp.PermissionIds = append(resp.PermissionIds, policy[0])
		}
	}

	return &resp, nil
}

func (s *User) GetAllRelations(ctx context.Context, tenant string, uuid string) (*ports.UserAllRelationsResponse, error) {
	resp := ports.UserAllRelationsResponse{
		PermissionItems: []ports.ItemRelation{},
		NodeItems:       []ports.ItemRelation{},
		GroupItems:      []ports.ItemRelation{},
	}
	prefix := fmt.Sprintf("%s%s", USERPREFIX, uuid)
	data, err := s.GetAllRelationsForPrefix(prefix, tenant)
	if err != nil {
		return nil, err
	}

	for _, relation := range data {
		if strings.HasPrefix(relation, NODEPREFIX) {
			nodeModel, err := s.EntClient.Node.Get(ctx, cast.ToInt(relation[len(NODEPREFIX):]))
			if err != nil {
				return nil, err
			}
			resp.NodeItems = append(resp.NodeItems, ports.ItemRelation{
				Id:     cast.ToString(nodeModel.ID),
				Name:   nodeModel.Name,
				Tenant: ports.ItemTenant(nodeModel.Tenant),
				Type:   (*ports.ItemType)(&nodeModel.Type),
			})

			continue
		}

		if strings.HasPrefix(relation, GROUPPREFIX) {
			groupModel, err := s.EntClient.Group.Get(ctx, cast.ToInt(relation[len(GROUPPREFIX):]))
			if err != nil {
				return nil, err
			}
			resp.GroupItems = append(resp.GroupItems, ports.ItemRelation{
				Id:     cast.ToString(groupModel.ID),
				Name:   groupModel.Name,
				Tenant: ports.ItemTenant(groupModel.Tenant),
				Type:   (*ports.ItemType)(&groupModel.Type),
			})

			continue
		}

		permissionModel, err := s.EntClient.Permission.Get(ctx, cast.ToInt(relation))
		if err != nil {
			return nil, err
		}
		resp.PermissionItems = append(resp.PermissionItems, ports.ItemRelation{
			Id:     cast.ToString(permissionModel.ID),
			Name:   permissionModel.Name,
			Tenant: ports.ItemTenant(permissionModel.Tenant),
		})
	}

	return &resp, nil
}
