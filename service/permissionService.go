package service

import (
	"context"
	"database/sql"

	"github.com/keepondream/RBAC_service/store/rbacStore"
)

type Permission struct {
	Id      string `json:"id"`
	Sign    string `json:"-"`
	Tenant  string `json:"tenant"`
	Uri     string `json:"uri"`
	Method  string `json:"method"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func (s *Service) AddPermission(ctx context.Context, p Permission) (bool, error) {
	return s.Enforcer.AddPolicy(
		p.Sign,
		p.Tenant,
		p.Uri,
		p.Method,
		p.Name,
		p.Comment,
	)
}

func (s *Service) DeletePermission(ctx context.Context, id, tenant string) error {
	oldP, err := s.GetPermission(ctx, id, tenant)
	if err != nil {
		return err
	}

	_, err = s.Enforcer.RemoveFilteredPolicy(1,
		oldP.V1.String,
		oldP.V2.String,
		oldP.V3.String,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetPermission(ctx context.Context, id, tenant string) (rbacStore.RbacCasbinRule, error) {
	return s.Store.GetInfoByIDTenant(ctx, rbacStore.GetInfoByIDTenantParams{
		ID: id,
		V1: sql.NullString{
			String: tenant,
			Valid:  true,
		},
	})
}

type ListPermission struct {
	Total int64        `json:"total"`
	List  []Permission `json:"list"`
}

func (s *Service) ListPermission(ctx context.Context, page, pageSize int32, tenant string) (ListPermission, error) {
	resp := ListPermission{
		Total: 0,
		List:  []Permission{},
	}
	total, err := s.Store.TotalBySignTenant(ctx, rbacStore.TotalBySignTenantParams{
		V0: sql.NullString{
			String: AllPermission,
			Valid:  true,
		},
		V1: sql.NullString{
			String: tenant,
			Valid:  true,
		},
	})
	if err != nil {
		return resp, err
	}
	resp.Total = total

	list, err := s.Store.ListBySignTenant(ctx, rbacStore.ListBySignTenantParams{
		V0: sql.NullString{
			String: AllPermission,
			Valid:  true,
		},
		V1: sql.NullString{
			String: tenant,
			Valid:  true,
		},
		Limit:  pageSize,
		Offset: (page - 1) * pageSize,
	})
	if err != nil && err != sql.ErrNoRows {
		return resp, err
	}

	for _, v := range list {
		resp.List = append(resp.List, Permission{
			Id:      v.ID,
			Sign:    v.V0.String,
			Tenant:  v.V1.String,
			Uri:     v.V2.String,
			Method:  v.V3.String,
			Name:    v.V4.String,
			Comment: v.V5.String,
		})
	}

	return resp, nil
}
