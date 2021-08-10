package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/keepondream/RBAC_service/internal/common/utils"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/group"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/node"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
	"github.com/spf13/cast"
)

type Group struct {
	*Repo
}

func NewGroup(r *Repo) *Group {
	return &Group{
		Repo: r,
	}
}

func (r *Group) Ent2Port(model *ent.Group) *ports.Group {
	resp := ports.Group{
		Data:      model.Data,
		Id:        cast.ToString(model.ID),
		Name:      model.Name,
		Tenant:    ports.ItemTenant(model.Tenant),
		Type:      ports.ItemType(model.Type),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}

	return &resp
}

func (r *Group) Model2Response(ctx context.Context, model *ent.Group) *ports.GroupInfoResponse {
	resp := ports.GroupInfoResponse{
		Group:    *r.Ent2Port(model),
		Nodes:    []ports.Node{},
		Children: []ports.Group{},
	}

	if model.Edges.Parent != nil {
		resp.Parent = r.Ent2Port(model.Edges.Parent)
	}

	for _, c := range model.Edges.Children {
		resp.Children = append(resp.Children, *r.Ent2Port(c))
	}

	nodeRepo := NewNode(r.Repo)
	for _, n := range model.Edges.Nodes {
		resp.Nodes = append(resp.Nodes, *nodeRepo.Ent2Port(n))
	}

	return &resp
}

func (r *Group) Create(ctx context.Context, params ports.PostGroupsJSONBody) (*ports.GroupInfoResponse, error) {
	nodeIds := []int{}
	for _, v := range params.NodeIds {
		nodeIds = append(nodeIds, cast.ToInt(v))
	}
	nodes := r.EntClient.Node.Query().Where(node.Tenant(string(params.Tenant))).Where(node.IDIn(nodeIds...)).AllX(ctx)
	modelQuery := r.EntClient.Group.Create().
		SetName(params.Name).
		SetTenant(string(params.Tenant)).
		SetType(string(params.Type)).
		AddNodes(nodes...)
	if params.Data != nil {
		modelQuery.SetData((*interface{})(params.Data))
	}
	if params.ParentId != nil {
		parent, err := r.EntClient.Group.Query().Where(group.Tenant(string(params.Tenant))).Where(group.ID(cast.ToInt(params.ParentId))).First(ctx)
		if err != nil {
			return nil, err
		}
		modelQuery.SetParent(parent)
	}

	model, err := modelQuery.Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.GetById(ctx, cast.ToString(model.ID))
}

func (r *Group) IsUnique(ctx context.Context, tenant string, name string, group_type string) error {
	exist, err := r.EntClient.Group.Query().
		Where(group.Tenant(tenant)).
		Where(group.Name(name)).
		Where(group.Type(group_type)).Exist(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("exist not unique")
	}

	return nil
}

func (r *Group) GetById(ctx context.Context, id string) (*ports.GroupInfoResponse, error) {
	model, err := r.EntClient.Group.Query().
		WithParent().
		WithChildren().
		WithNodes().
		Where(group.ID(cast.ToInt(id))).First(ctx)
	if err != nil {
		return nil, err
	}

	return r.Model2Response(ctx, model), nil
}

func (r *Group) DeleteById(ctx context.Context, id string) error {
	return r.EntClient.Group.DeleteOneID(cast.ToInt(id)).Exec(ctx)
}

func (r *Group) List(ctx context.Context, params ports.GetGroupsParams) (*ports.GroupListResponse, error) {
	list := ports.GroupListResponse{
		Items: []ports.GroupInfoResponse{},
		Total: "0",
	}
	pageSize := ""
	if params.PerPage != nil {
		pageSize = string(*params.PerPage)
	}

	offset, limit := utils.ParsePagination(string(params.Page), pageSize)

	modelQuery := r.EntClient.Group.Query().
		WithParent().
		WithChildren().
		WithNodes()

	if params.Query != nil {
		conditions := utils.ParseQuery(string(*params.Query))
		if values, ok := conditions["default"]; ok {
			modelQuery.Where(group.NameIn(values...))
		}
		if values, ok := conditions[group.FieldID]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(group.IDIn(ids...))
		}
		if values, ok := conditions[group.FieldTenant]; ok {
			modelQuery.Where(group.TenantIn(values...))
		}
		if values, ok := conditions[group.FieldParentID]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(group.ParentIDIn(ids...))
		}
		// 根据父级相关属性查询
		if values, ok := conditions[fmt.Sprintf("%s.%s", group.EdgeParent, group.FieldID)]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(group.HasParentWith(group.IDIn(ids...)))
		}
		if values, ok := conditions[fmt.Sprintf("%s.%s", group.EdgeParent, group.FieldName)]; ok {
			modelQuery.Where(group.HasParentWith(group.NameIn(values...)))
		}
		// 根据子级相关属性查询
		if values, ok := conditions[fmt.Sprintf("%s.%s", group.EdgeChildren, group.FieldID)]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(group.HasChildrenWith(group.IDIn(ids...)))
		}
		if values, ok := conditions[fmt.Sprintf("%s.%s", group.EdgeChildren, group.FieldName)]; ok {
			modelQuery.Where(group.HasChildrenWith(group.NameIn(values...)))
		}
		// 根据节点相关属性查询
		if values, ok := conditions[fmt.Sprintf("%s.%s", group.EdgeNodes, node.FieldID)]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(group.HasNodesWith(node.IDIn(ids...)))
		}
		if values, ok := conditions[fmt.Sprintf("%s.%s", group.EdgeNodes, node.FieldName)]; ok {
			modelQuery.Where(group.HasNodesWith(node.NameIn(values...)))
		}
	}
	if params.StartTime != nil {
		start := time.Time(*params.StartTime)
		if !start.IsZero() {
			modelQuery.Where(group.CreatedAtGTE(start))
		}
	}

	if params.EndTime != nil {
		end := time.Time(*params.EndTime)
		if !end.IsZero() {
			modelQuery.Where(group.CreatedAtLTE(end))
		}
	}

	total := modelQuery.CountX(ctx)
	list.Total = cast.ToString(total)

	if params.Sort != nil {
		fields := utils.ParseSort(string(*params.Sort))
		if params.Order == ports.GetGroupsParamsOrder(ports.Asc) {
			modelQuery.Order(ent.Asc(fields...))
		} else {
			modelQuery.Order(ent.Desc(fields...))
		}
	} else {
		modelQuery.Order(ent.Desc(node.FieldID))
	}

	models := modelQuery.Offset(offset).Limit(limit).AllX(ctx)

	for _, v := range models {
		list.Items = append(list.Items, *r.Model2Response(ctx, v))
	}

	return &list, nil
}

func (r *Group) Update(ctx context.Context, params ports.PatchGroupsIdJSONBody, id string) (*ports.GroupInfoResponse, error) {
	model, err := r.EntClient.Group.Get(ctx, cast.ToInt(id))
	if err != nil {
		return nil, err
	}
	modelQuery := model.Update()
	if params.Name != nil {
		modelQuery.SetName(*params.Name)
	}
	if params.Data != nil {
		modelQuery.SetData((*interface{})(params.Data))
	}
	if params.Tenant != nil {
		modelQuery.SetTenant(string(*params.Tenant))
	}
	if params.Type != nil {
		modelQuery.SetType(string(*params.Type))
	}
	if params.ParentId != nil {
		if *params.ParentId == "" {
			modelQuery.ClearParent()
		} else {
			parent, err := r.EntClient.Group.Query().Where(group.Tenant(model.Tenant)).Where(group.ID(cast.ToInt(params.ParentId))).First(ctx)
			if err != nil {
				return nil, err
			}
			modelQuery.ClearParent().SetParent(parent)
		}
	}

	if params.NodeIds != nil {
		nodeIds := []int{}
		for _, v := range *params.NodeIds {
			nodeIds = append(nodeIds, cast.ToInt(v))
		}

		nodes := r.EntClient.Node.Query().Where(node.Tenant(model.Tenant)).
			Where(node.IDIn(nodeIds...)).AllX(ctx)
		modelQuery.ClearNodes().AddNodes(nodes...)
	}

	model, err = modelQuery.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, cast.ToString(model.ID))
}
