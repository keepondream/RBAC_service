package repo

import (
	"context"
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
		Id:   cast.ToString(model.ID),
		Name: model.Name,
		Uri:  model.URI,
		ItemMethod: ports.ItemMethod{
			Method: ports.ItemMethodMethod(model.Method),
		},
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
}

func (r *Route) GetById(ctx context.Context, id string) (*ports.RouteDetail, error) {
	model, err := r.EntClient.Route.Query().Where(route.ID(cast.ToInt(id))).First(ctx)
	if err != nil {
		return nil, err
	}

	portsRoute := r.Ent2Port(model)

	return (*ports.RouteDetail)(portsRoute), nil
}

func (r *Route) List(ctx context.Context, params ports.GetRoutesParams) (*ports.RouteList, error) {
	list := ports.RouteList{
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
		modelQuery.Where(route.CreatedAtGTE(time.Time(*params.StartTime)))
	}

	if params.EndTime != nil {
		modelQuery.Where(route.CreatedAtLTE(time.Time(*params.EndTime)))
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

	panic("not implemented") // TODO: Implement
}

func (r *Route) DeleteById(ctx context.Context, id string) error {
	return r.EntClient.Route.DeleteOneID(cast.ToInt(id)).Exec(ctx)
}

func (r *Route) Update(ctx context.Context, params ports.PatchRoutesIdJSONBody, id string) (*ports.RouteDetail, error) {
	modelQuery := r.EntClient.Route.UpdateOneID(cast.ToInt(id))
	if params.Data != nil {
		modelQuery.SetData(&params.Data.Data)
	}
	if params.Method != nil {
		modelQuery.SetMethod(route.Method(params.Method.Method))
	}
	if params.Name != nil {
		modelQuery.SetName(*params.Name)
	}
	if params.Tenant != nil {
		modelQuery.SetTenant(params.Tenant.Tenant)
	}
	if params.Uri != nil {
		modelQuery.SetURI(*params.Uri)
	}

	model, err := modelQuery.Save(ctx)
	if err != nil {
		return nil, err
	}

	portsRoute := r.Ent2Port(model)
	return (*ports.RouteDetail)(portsRoute), nil
}

func (r *Route) Create(ctx context.Context, params ports.PostRoutesJSONBody) (*ports.RouteDetail, error) {
	model, err := r.EntClient.Route.Create().
		SetTenant(params.Tenant).
		SetName(params.Name).
		SetURI(params.Uri).
		SetMethod(route.Method(params.Method)).
		SetData(&params.Data).Save(ctx)
	if err != nil {
		return nil, err
	}

	portsRoute := r.Ent2Port(model)
	return (*ports.RouteDetail)(portsRoute), nil
}
