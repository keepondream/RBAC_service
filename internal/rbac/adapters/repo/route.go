package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/keepondream/RBAC_service/internal/common/utils"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/route"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
	"github.com/spf13/cast"
)

type Route struct {
	*Repo
}

func NewRoute(r *Repo) *Route {
	return &Route{
		Repo: r,
	}
}

func (r *Route) Ent2Port(model *ent.Route) *ports.Route {
	return &ports.Route{
		Id:        cast.ToString(model.ID),
		Name:      model.Name,
		Uri:       model.URI,
		Method:    ports.ItemMethod(model.Method),
		Tenant:    ports.ItemTenant(model.Tenant),
		Data:      model.Data,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func (r *Route) GetById(ctx context.Context, id string) (*ports.RouteInfoResponse, error) {
	model, err := r.EntClient.Route.Get(ctx, cast.ToInt(id))
	if err != nil {
		return nil, err
	}

	portsRoute := r.Ent2Port(model)

	return (*ports.RouteInfoResponse)(portsRoute), nil
}

func (r *Route) List(ctx context.Context, params ports.GetRoutesParams) (*ports.RouteListResponse, error) {
	list := ports.RouteListResponse{
		Items: []ports.Route{},
		Total: "0",
	}

	pageSize := ""
	if params.PerPage != nil {
		pageSize = string(*params.PerPage)
	}

	offset, limit := utils.ParsePagination(string(params.Page), pageSize)

	modelQuery := r.EntClient.Route.Query()

	if params.Query != nil {
		conditions := utils.ParseQuery(string(*params.Query))
		if values, ok := conditions["default"]; ok {
			modelQuery.Where(route.NameIn(values...))
		}
		if values, ok := conditions[route.FieldID]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(route.IDIn(ids...))
		}
		if values, ok := conditions[route.FieldTenant]; ok {
			modelQuery.Where(route.TenantIn(values...))
		}
		if values, ok := conditions[route.FieldMethod]; ok {
			methods := []route.Method{}
			for _, v := range values {
				methods = append(methods, route.Method(v))
			}
			modelQuery.Where(route.MethodIn(methods...))
		}
		if values, ok := conditions[route.FieldURI]; ok {
			modelQuery.Where(route.URIIn(values...))
		}
	}

	if params.StartTime != nil {
		start := time.Time(*params.StartTime)
		if !start.IsZero() {
			modelQuery.Where(route.CreatedAtGTE(start))
		}
	}

	if params.EndTime != nil {
		end := time.Time(*params.EndTime)
		if !end.IsZero() {
			modelQuery.Where(route.CreatedAtLTE(end))
		}
	}

	total := modelQuery.CountX(ctx)
	list.Total = cast.ToString(total)
	if params.Sort != nil {
		fields := utils.ParseSort(string(*params.Sort))
		if string(params.Order) == string(ports.Asc) {
			modelQuery.Order(ent.Asc(fields...))
		} else {
			modelQuery.Order(ent.Desc(fields...))
		}
	} else {
		modelQuery.Order(ent.Desc(route.FieldID))
	}

	models := modelQuery.Offset(offset).Limit(limit).AllX(ctx)

	for _, v := range models {
		portsRoute := r.Ent2Port(v)
		list.Items = append(list.Items, *portsRoute)
	}

	return &list, nil
}

func (r *Route) DeleteById(ctx context.Context, id string) error {
	return r.EntClient.Route.DeleteOneID(cast.ToInt(id)).Exec(ctx)
}

func (r *Route) Update(ctx context.Context, params ports.PatchRoutesIdJSONBody, id string) (*ports.RouteInfoResponse, error) {
	modelQuery := r.EntClient.Route.UpdateOneID(cast.ToInt(id))
	if params.Data != nil {
		modelQuery.SetData((*interface{})(params.Data))
	}
	if params.Method != nil {
		modelQuery.SetMethod(route.Method(*params.Method))
	}
	if params.Name != nil {
		modelQuery.SetName(*params.Name)
	}
	if params.Tenant != nil {
		modelQuery.SetTenant(string(*params.Tenant))
	}
	if params.Uri != nil {
		modelQuery.SetURI(*params.Uri)
	}

	model, err := modelQuery.Save(ctx)
	if err != nil {
		return nil, err
	}

	portsRoute := r.Ent2Port(model)
	return (*ports.RouteInfoResponse)(portsRoute), nil
}

func (r *Route) Create(ctx context.Context, params ports.PostRoutesJSONBody) (*ports.RouteInfoResponse, error) {
	model, err := r.EntClient.Route.Create().
		SetTenant(string(params.Tenant)).
		SetName(params.Name).
		SetURI(params.Uri).
		SetMethod(route.Method(params.Method)).
		SetData((*interface{})(&params.Data)).Save(ctx)
	if err != nil {
		return nil, err
	}

	portsRoute := r.Ent2Port(model)
	return (*ports.RouteInfoResponse)(portsRoute), nil
}

func (r *Route) IsUnique(ctx context.Context, tenant string, uri string, method string) error {
	exist, err := r.EntClient.Route.Query().
		Where(route.Tenant(tenant)).
		Where(route.URI(uri)).
		Where(route.MethodEQ(route.Method(method))).
		Exist(ctx)
	if err != nil {
		return err
	}

	if exist {
		return fmt.Errorf("exist not unique")
	}

	return nil
}
