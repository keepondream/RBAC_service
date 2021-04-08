package service

import (
	"context"
	"database/sql"

	"github.com/keepondream/RBAC_service/store/rbacStore"
)

func (s *Service) AddPermission(ctx context.Context, sign, tenant, uri, method, name, comment string) (bool, error) {
	return s.Enforcer.AddPolicy(
		sign,
		tenant,
		uri,
		method,
		name,
		comment,
	)
}

func (s *Service) EditPermission(ctx context.Context, id, tenant, uri, method, name, comment string) error {
	oldP, err := s.GetPermission(ctx, id, tenant)
	if err != nil {
		return err
	}
	oldPolicy := []string{
		oldP.V0.String,
		oldP.V1.String,
		oldP.V2.String,
		oldP.V3.String,
		oldP.V4.String,
		oldP.V5.String,
	}
	newPolicy := []string{
		oldP.V0.String,
		oldP.V1.String,
		uri,
		method,
		name,
		comment,
	}

	_, err = s.Enforcer.UpdatePolicy(oldPolicy, newPolicy)
	if err != nil {
		return err
	}

	return nil
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
