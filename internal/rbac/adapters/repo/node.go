package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/keepondream/RBAC_service/internal/common/utils"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/node"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/permission"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
	"github.com/spf13/cast"
)

type Node struct {
	*Repo
}

func NewNode(r *Repo) *Node {
	return &Node{
		Repo: r,
	}
}

func (r *Node) Ent2Port(model *ent.Node) *ports.Node {
	resp := ports.Node{
		Id:        cast.ToString(model.ID),
		Name:      model.Name,
		ParentId:  cast.ToString(model.ParentID),
		Data:      model.Data,
		Tenant:    ports.ItemTenant(model.Tenant),
		Type:      ports.ItemType(model.Type),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}

	return &resp
}
func (r *Node) Model2Response(ctx context.Context, model *ent.Node) *ports.NodeInfoResponse {
	resp := ports.NodeInfoResponse{
		Node:        *r.Ent2Port(model),
		Children:    []ports.Node{},
		Permissions: []ports.Permission{},
	}

	if model.Edges.Parent != nil {
		resp.Parent = r.Ent2Port(model.Edges.Parent)
	}

	for _, c := range model.Edges.Children {
		resp.Children = append(resp.Children, *r.Ent2Port(c))
	}

	permissionRepo := NewPermission(r.Repo)
	for _, p := range model.Edges.Permissions {
		resp.Permissions = append(resp.Permissions, *permissionRepo.Ent2Port(p))
	}

	return &resp
}

func (r *Node) Create(ctx context.Context, params ports.PostNodesJSONBody) (*ports.NodeInfoResponse, error) {
	permissionIds := []int{}
	for _, v := range params.PermissionIds {
		permissionIds = append(permissionIds, cast.ToInt(v))
	}
	permissions := r.EntClient.Permission.Query().Where(permission.Tenant(string(params.Tenant))).Where(permission.IDIn(permissionIds...)).AllX(ctx)
	modelQuery := r.EntClient.Node.Create().
		SetName(params.Name).
		SetTenant(string(params.Tenant)).
		SetType(string(params.Type)).
		AddPermissions(permissions...)
	if params.Data != nil {
		modelQuery.SetData((*interface{})(params.Data))
	}
	if params.ParentId != nil {
		parent, err := r.EntClient.Node.Query().Where(node.Tenant(string(params.Tenant))).Where(node.ID(cast.ToInt(params.ParentId))).First(ctx)
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

func (r *Node) IsUnique(ctx context.Context, tenant string, name string, node_type string) error {
	exist, err := r.EntClient.Node.Query().
		Where(node.Tenant(tenant)).
		Where(node.Name(name)).
		Where(node.Type(node_type)).
		Exist(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("exist not unique")
	}

	return nil
}

func (r *Node) GetById(ctx context.Context, id string) (*ports.NodeInfoResponse, error) {
	model, err := r.EntClient.Node.Query().
		WithParent().
		WithChildren().
		WithPermissions().
		Where(node.IDEQ(cast.ToInt(id))).First(ctx)
	if err != nil {
		return nil, err
	}

	return r.Model2Response(ctx, model), nil
}

func (r *Node) DeleteById(ctx context.Context, id string) error {
	return r.EntClient.Node.DeleteOneID(cast.ToInt(id)).Exec(ctx)
}

func (r *Node) List(ctx context.Context, params ports.GetNodesParams) (*ports.NodeListResponse, error) {
	list := ports.NodeListResponse{
		Items: []ports.NodeInfoResponse{},
		Total: "0",
	}

	pageSize := ""
	if params.PerPage != nil {
		pageSize = string(*params.PerPage)
	}

	offset, limit := utils.ParsePagination(string(params.Page), pageSize)

	modelQuery := r.EntClient.Node.Query().
		WithParent().
		WithChildren().
		WithPermissions()

	if params.Query != nil {
		conditions := utils.ParseQuery(string(*params.Query))
		if values, ok := conditions["default"]; ok {
			modelQuery.Where(node.NameIn(values...))
		}
		if values, ok := conditions[node.FieldID]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(node.IDIn(ids...))
		}
		if values, ok := conditions[node.FieldTenant]; ok {
			modelQuery.Where(node.TenantIn(values...))
		}
		if values, ok := conditions[node.FieldParentID]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(node.ParentIDIn(ids...))
		}
		// 根据父级相关属性查询
		if values, ok := conditions[fmt.Sprintf("%s.%s", node.EdgeParent, node.FieldID)]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(node.HasParentWith(node.IDIn(ids...)))
		}
		if values, ok := conditions[fmt.Sprintf("%s.%s", node.EdgeParent, node.FieldName)]; ok {
			modelQuery.Where(node.HasParentWith(node.NameIn(values...)))
		}
		// 根据子级相关属性查询
		if values, ok := conditions[fmt.Sprintf("%s.%s", node.EdgeChildren, node.FieldID)]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(node.HasChildrenWith(node.IDIn(ids...)))
		}
		if values, ok := conditions[fmt.Sprintf("%s.%s", node.EdgeChildren, node.FieldName)]; ok {
			modelQuery.Where(node.HasChildrenWith(node.NameIn(values...)))
		}
		// 根据权限相关属性查询
		if values, ok := conditions[fmt.Sprintf("%s.%s", node.EdgePermissions, permission.FieldID)]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(node.HasPermissionsWith(permission.IDIn(ids...)))
		}
		if values, ok := conditions[fmt.Sprintf("%s.%s", node.EdgePermissions, permission.FieldName)]; ok {
			modelQuery.Where(node.HasPermissionsWith(permission.NameIn(values...)))
		}
	}

	if params.StartTime != nil {
		start := time.Time(*params.StartTime)
		if !start.IsZero() {
			modelQuery.Where(node.CreatedAtGTE(start))
		}
	}

	if params.EndTime != nil {
		end := time.Time(*params.EndTime)
		if !end.IsZero() {
			modelQuery.Where(node.CreatedAtLTE(end))
		}
	}

	total := modelQuery.CountX(ctx)
	list.Total = cast.ToString(total)

	if params.Sort != nil {
		fields := utils.ParseSort(string(*params.Sort))
		if params.Order == ports.GetNodesParamsOrder(ports.Asc) {
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

func (r *Node) Update(ctx context.Context, params ports.PatchNodesIdJSONBody, id string) (*ports.NodeInfoResponse, error) {
	model, err := r.EntClient.Node.Get(ctx, cast.ToInt(id))
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
	if params.ParentId != nil {
		if *params.ParentId == "" {
			modelQuery.ClearParent()
		} else {
			parent, err := r.EntClient.Node.Query().Where(node.Tenant(model.Tenant)).Where(node.ID(cast.ToInt(params.ParentId))).First(ctx)
			if err != nil {
				return nil, err
			}
			modelQuery.ClearParent().SetParent(parent)
		}
	}
	if params.Tenant != nil {
		modelQuery.SetTenant(string(*params.Tenant))
	}
	if params.PermissionIds != nil {
		permissionIds := []int{}
		for _, v := range *params.PermissionIds {
			permissionIds = append(permissionIds, cast.ToInt(v))
		}
		permissions := r.EntClient.Permission.Query().Where(permission.Tenant(model.Tenant)).
			Where(permission.IDIn(permissionIds...)).AllX(ctx)
		modelQuery.ClearPermissions().AddPermissions(permissions...)
	}
	if params.Type != nil {
		modelQuery.SetType(string(*params.Type))
	}
	model, err = modelQuery.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, cast.ToString(model.ID))
}
