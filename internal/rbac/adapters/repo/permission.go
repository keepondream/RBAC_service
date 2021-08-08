package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/keepondream/RBAC_service/internal/common/utils"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/permission"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/route"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
	"github.com/spf13/cast"
)

type Permission struct {
	*Repo
}

func NewPermission(r *Repo) *Permission {
	return &Permission{
		Repo: r,
	}
}

func (r *Permission) Ent2Port(model *ent.Permission) *ports.Permission {
	resp := ports.Permission{
		Id:   cast.ToString(model.ID),
		Name: model.Name,
		ItemTenant: ports.ItemTenant{
			Tenant: model.Tenant,
		},
		ItemData: ports.ItemData{
			Data: model.Data,
		},
		ItemCreatedat: ports.ItemCreatedat{
			CreatedAt: model.CreatedAt,
		},
		ItemUpdatedat: ports.ItemUpdatedat{
			UpdatedAt: model.UpdatedAt,
		},
	}

	return &resp
}

func (r *Permission) Model2Response(ctx context.Context, model *ent.Permission) *ports.PermissionInfoResponse {
	routes := []ports.Route{}
	routeRepo := NewRoute(r.Repo)
	for _, route := range model.QueryRoutes().AllX(ctx) {
		routes = append(routes, *routeRepo.Ent2Port(route))
	}
	return &ports.PermissionInfoResponse{
		Permission: *r.Ent2Port(model),
		Routes:     &routes,
	}
}

func (r *Permission) Create(ctx context.Context, params ports.PostPermissionsJSONBody) (*ports.PermissionInfoResponse, error) {
	routeIds := []int{}
	for _, v := range params.RouteIds {
		routeIds = append(routeIds, cast.ToInt(v))
	}
	routes := r.EntClient.Route.Query().Where(route.Tenant(params.Tenant)).Where(route.IDIn(routeIds...)).AllX(ctx)
	model, err := r.EntClient.Permission.Create().
		SetName(params.Name).
		SetTenant(params.Tenant).
		SetData(&params.Data).
		AddRoutes(routes...).Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.Model2Response(ctx, model), nil
}

func (r *Permission) IsUnique(ctx context.Context, tenant string, name string) error {
	exist, err := r.EntClient.Permission.Query().
		Where(permission.Tenant(tenant)).
		Where(permission.Name(name)).
		Exist(ctx)
	if err != nil {
		return err
	}

	if exist {
		return fmt.Errorf("exist not unique")
	}

	return nil
}

func (r *Permission) GetById(ctx context.Context, id string) (*ports.PermissionInfoResponse, error) {
	model, err := r.EntClient.Permission.Get(ctx, cast.ToInt(id))
	if err != nil {
		return nil, err
	}

	return r.Model2Response(ctx, model), nil
}

func (r *Permission) List(ctx context.Context, params ports.GetPermissionsParams) (*ports.PermissionListResponse, error) {
	list := ports.PermissionListResponse{
		Items: []ports.PermissionInfoResponse{},
		Total: "0",
	}

	pageSize := ""
	if params.PerPage != nil {
		pageSize = string(*params.PerPage)
	}

	offset, limit := utils.ParsePagination(string(params.Page), pageSize)

	modelQuery := r.EntClient.Permission.Query().WithRoutes()

	if params.Query != nil {
		conditions := utils.ParseQuery(string(*params.Query))
		if values, ok := conditions["default"]; ok {
			modelQuery.Where(permission.NameIn(values...))
		}
		if values, ok := conditions[permission.FieldID]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(permission.IDIn(ids...))
		}
		if values, ok := conditions[permission.FieldTenant]; ok {
			modelQuery.Where(permission.TenantIn(values...))
		}
		// 根据路由相关属性查询
		if values, ok := conditions[fmt.Sprintf("%s.%s", permission.EdgeRoutes, route.FieldID)]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(permission.HasRoutesWith(route.IDIn(ids...)))
		}
		if values, ok := conditions[fmt.Sprintf("%s.%s", permission.EdgeRoutes, route.FieldName)]; ok {
			modelQuery.Where(permission.HasRoutesWith(route.NameIn(values...)))
		}
		if values, ok := conditions[fmt.Sprintf("%s.%s", permission.EdgeRoutes, route.FieldTenant)]; ok {
			modelQuery.Where(permission.HasRoutesWith(route.TenantIn(values...)))
		}
		if values, ok := conditions[fmt.Sprintf("%s.%s", permission.EdgeRoutes, route.FieldMethod)]; ok {
			methods := []route.Method{}
			for _, v := range values {
				methods = append(methods, route.Method(v))
			}
			modelQuery.Where(permission.HasRoutesWith(route.MethodIn(methods...)))
		}
	}

	if params.StartTime != nil {
		start := time.Time(*params.StartTime)
		if !start.IsZero() {
			modelQuery.Where(permission.CreatedAtGTE(start))
		}
	}

	if params.EndTime != nil {
		end := time.Time(*params.EndTime)
		if !end.IsZero() {
			modelQuery.Where(permission.CreatedAtLTE(end))
		}
	}

	total := modelQuery.CountX(ctx)
	list.Total = cast.ToString(total)

	if params.Sort != nil {
		fields := utils.ParseSort(string(*params.Sort))
		if params.Order == ports.GetPermissionsParamsOrder(ports.Asc) {
			modelQuery.Order(ent.Asc(fields...))
		} else {
			modelQuery.Order(ent.Desc(fields...))
		}
	} else {
		modelQuery.Order(ent.Desc(permission.FieldID))
	}
	models := modelQuery.Offset(offset).Limit(limit).AllX(ctx)

	for _, v := range models {
		list.Items = append(list.Items, *r.Model2Response(ctx, v))
	}

	return &list, nil
}

func (r *Permission) DeleteById(ctx context.Context, id string) error {
	return r.EntClient.Permission.DeleteOneID(cast.ToInt(id)).Exec(ctx)
}

func (r *Permission) Update(ctx context.Context, params ports.PatchPermissionsIdJSONBody, id string) (*ports.PermissionInfoResponse, error) {
	modelQuery := r.EntClient.Permission.UpdateOneID(cast.ToInt(id))

	if params.Name != nil {
		modelQuery.SetName(*params.Name)
	}
	if params.Data != nil {
		modelQuery.SetData(&params.Data.Data)
	}
	if params.Tenant != nil {
		modelQuery.SetTenant(params.Tenant.Tenant)
	}
	if params.RouteIds != nil {
		routeIds := []int{}
		for _, v := range *params.RouteIds {
			routeIds = append(routeIds, cast.ToInt(v))
		}
		routes := r.EntClient.Route.Query().Where(route.Tenant(params.Tenant.Tenant)).Where(route.IDIn(routeIds...)).AllX(ctx)
		modelQuery.ClearRoutes().AddRoutes(routes...)
	}

	model, err := modelQuery.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.Model2Response(ctx, model), nil
}
