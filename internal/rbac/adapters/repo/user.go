package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/keepondream/RBAC_service/internal/common/utils"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/group"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/node"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/permission"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/user"
	"github.com/keepondream/RBAC_service/internal/rbac/ports"
	"github.com/spf13/cast"
)

type User struct {
	*Repo
}

func NewUser(r *Repo) *User {
	return &User{
		Repo: r,
	}
}

func (r *User) Ent2Port(model *ent.User) *ports.User {
	resp := ports.User{
		CreatedAt: model.CreatedAt,
		Data:      model.Data,
		Id:        cast.ToString(model.ID),
		IsSuper:   model.IsSuper,
		ParentId:  cast.ToString(model.ParentID),
		Tenant:    ports.ItemTenant(model.Tenant),
		UpdatedAt: model.UpdatedAt,
		Uuid:      model.UUID,
	}

	return &resp
}

func (r *User) Model2Response(ctx context.Context, model *ent.User) *ports.UserInfoResponse {
	resp := ports.UserInfoResponse{
		User:        *r.Ent2Port(model),
		Children:    []ports.User{},
		Groups:      []ports.Group{},
		Nodes:       []ports.Node{},
		Permissions: []ports.Permission{},
	}

	if model.Edges.Parent != nil {
		resp.Parent = r.Ent2Port(model.Edges.Parent)
	}

	for _, c := range model.Edges.Children {
		resp.Children = append(resp.Children, *r.Ent2Port(c))
	}

	groupRepo := NewGroup(r.Repo)
	for _, g := range model.Edges.Groups {
		resp.Groups = append(resp.Groups, *groupRepo.Ent2Port(g))
	}

	nodeRepo := NewNode(r.Repo)
	for _, n := range model.Edges.Nodes {
		resp.Nodes = append(resp.Nodes, *nodeRepo.Ent2Port(n))
	}

	permissionRepo := NewPermission(r.Repo)
	for _, p := range model.Edges.Permissions {
		resp.Permissions = append(resp.Permissions, *permissionRepo.Ent2Port(p))
	}

	return &resp
}

func (r *User) Create(ctx context.Context, params ports.PostUsersJSONBody) (*ports.UserInfoResponse, error) {
	modelQuery := r.EntClient.User.Create().
		SetUUID(params.Uuid).
		SetTenant(string(params.Tenant)).
		SetIsSuper(params.IsSuper)

	if params.Data != nil {
		modelQuery.SetData((*interface{})(params.Data))
	}

	if params.ParentId != nil {
		parent, err := r.EntClient.User.Query().Where(user.Tenant(string(params.Tenant))).Where(user.ID(cast.ToInt(params.ParentId))).First(ctx)
		if err != nil {
			return nil, err
		}
		modelQuery.SetParent(parent)
	}

	nodeIds := []int{}
	for _, v := range params.NodeIds {
		nodeIds = append(nodeIds, cast.ToInt(v))
	}
	nodes := r.EntClient.Node.Query().Where(node.Tenant(string(params.Tenant))).
		Where(node.IDIn(nodeIds...)).AllX(ctx)
	modelQuery.AddNodes(nodes...)

	groupIds := []int{}
	for _, v := range params.GroupIds {
		groupIds = append(groupIds, cast.ToInt(v))
	}
	groups := r.EntClient.Group.Query().Where(group.Tenant(string(params.Tenant))).
		Where(group.IDIn(groupIds...)).AllX(ctx)
	modelQuery.AddGroups(groups...)

	permissionIds := []int{}
	for _, v := range params.PermissionIds {
		permissionIds = append(permissionIds, cast.ToInt(v))
	}
	permissions := r.EntClient.Permission.Query().Where(permission.Tenant(string(params.Tenant))).
		Where(permission.IDIn(permissionIds...)).AllX(ctx)
	modelQuery.AddPermissions(permissions...)

	model, err := modelQuery.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, cast.ToString(model.ID))
}

func (r *User) IsUnique(ctx context.Context, tenant string, uuid string) error {
	exist, err := r.EntClient.User.Query().
		Where(user.Tenant(tenant)).Where(user.UUID(uuid)).Exist(ctx)
	if err != nil {
		return err
	}

	if exist {
		return fmt.Errorf("exist not unique")
	}

	return nil
}

func (r *User) GetById(ctx context.Context, id string) (*ports.UserInfoResponse, error) {
	model, err := r.EntClient.User.Query().
		WithParent().
		WithChildren().
		WithNodes().
		WithGroups().
		WithPermissions().
		Where(user.ID(cast.ToInt(id))).First(ctx)
	if err != nil {
		return nil, err
	}

	return r.Model2Response(ctx, model), nil
}

func (r *User) GetByUuid(ctx context.Context, tenant string, uuid string) (*ports.UserInfoResponse, error) {
	id, err := r.EntClient.User.Query().Where(user.Tenant(tenant)).Where(user.UUID(uuid)).FirstID(ctx)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, cast.ToString(id))
}

func (r *User) DeleteByUuid(ctx context.Context, tenant string, uuid string) error {
	id, err := r.EntClient.User.Query().Where(user.Tenant(tenant)).Where(user.UUID(uuid)).FirstID(ctx)
	if err != nil {
		return err
	}

	return r.EntClient.User.DeleteOneID(id).Exec(ctx)
}

func (r *User) List(ctx context.Context, params ports.GetUsersParams) (*ports.UserListResponse, error) {
	list := ports.UserListResponse{
		Items: []ports.UserInfoResponse{},
		Total: "0",
	}

	pageSize := ""
	if params.PerPage != nil {
		pageSize = string(*params.PerPage)
	}

	offset, limit := utils.ParsePagination(string(params.Page), pageSize)

	modelQuery := r.EntClient.User.Query().
		WithParent().
		WithChildren().
		WithNodes().
		WithGroups().
		WithPermissions()
	if params.Query != nil {
		conditions := utils.ParseQuery(string(*params.Query))
		if values, ok := conditions["default"]; ok {
			modelQuery.Where(user.UUIDIn(values...))
		}
		if values, ok := conditions[user.FieldID]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(user.IDIn(ids...))
		}
		if values, ok := conditions[user.FieldTenant]; ok {
			modelQuery.Where(user.TenantIn(values...))
		}
		if values, ok := conditions[user.FieldParentID]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(user.ParentIDIn(ids...))
		}
		// 根据超管标识查询
		if values, ok := conditions[user.FieldIsSuper]; ok {
			modelQuery.Where(user.IsSuperEQ(cast.ToBool(values[0])))
		}
		// 根据父级相关属性查询
		if values, ok := conditions[fmt.Sprintf("%s.%s", user.EdgeParent, user.FieldID)]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(user.HasParentWith(user.IDIn(ids...)))
		}
		if values, ok := conditions[fmt.Sprintf("%s.%s", user.EdgeParent, user.FieldUUID)]; ok {
			modelQuery.Where(user.HasParentWith(user.UUIDIn(values...)))
		}
		// 根据子级相关属性查询
		if values, ok := conditions[fmt.Sprintf("%s.%s", user.EdgeChildren, user.FieldID)]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(user.HasChildrenWith(user.IDIn(ids...)))
		}
		if values, ok := conditions[fmt.Sprintf("%s.%s", user.EdgeChildren, user.FieldUUID)]; ok {
			modelQuery.Where(user.HasChildrenWith(user.UUIDIn(values...)))
		}
		// 根据节点相关属性查询
		if values, ok := conditions[fmt.Sprintf("%s.%s", user.EdgeNodes, node.FieldID)]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(user.HasNodesWith(node.IDIn(ids...)))
		}
		if values, ok := conditions[fmt.Sprintf("%s.%s", user.EdgeNodes, node.FieldName)]; ok {
			modelQuery.Where(user.HasNodesWith(node.NameIn(values...)))
		}
		// 根据权限相关属性查询
		if values, ok := conditions[fmt.Sprintf("%s.%s", user.EdgePermissions, permission.FieldID)]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(user.HasPermissionsWith(permission.IDIn(ids...)))
		}
		if values, ok := conditions[fmt.Sprintf("%s.%s", user.EdgePermissions, permission.FieldName)]; ok {
			modelQuery.Where(user.HasPermissionsWith(permission.NameIn(values...)))
		}
		// 根据分组相关属性查询
		if values, ok := conditions[fmt.Sprintf("%s.%s", user.EdgeGroups, group.FieldID)]; ok {
			ids := []int{}
			for _, v := range values {
				ids = append(ids, cast.ToInt(v))
			}
			modelQuery.Where(user.HasGroupsWith(group.IDIn(ids...)))
		}
		if values, ok := conditions[fmt.Sprintf("%s.%s", user.EdgeGroups, group.FieldName)]; ok {
			modelQuery.Where(user.HasGroupsWith(group.NameIn(values...)))
		}
	}

	if params.StartTime != nil {
		start := time.Time(*params.StartTime)
		if !start.IsZero() {
			modelQuery.Where(user.CreatedAtGTE(start))
		}
	}

	if params.EndTime != nil {
		end := time.Time(*params.EndTime)
		if !end.IsZero() {
			modelQuery.Where(user.CreatedAtLTE(end))
		}
	}

	total := modelQuery.CountX(ctx)
	list.Total = cast.ToString(total)

	if params.Sort != nil {
		fields := utils.ParseSort(string(*params.Sort))
		if params.Order == ports.GetUsersParamsOrder(ports.Asc) {
			modelQuery.Order(ent.Asc(fields...))
		} else {
			modelQuery.Order(ent.Desc(fields...))
		}
	} else {
		modelQuery.Order(ent.Desc(user.FieldID))
	}

	models := modelQuery.Offset(offset).Limit(limit).AllX(ctx)

	for _, v := range models {
		list.Items = append(list.Items, *r.Model2Response(ctx, v))
	}

	return &list, nil
}

func (r *User) Update(ctx context.Context, params ports.PatchUsersUuidTenantJSONBody, tenant string, uuid string) (*ports.UserInfoResponse, error) {
	model, err := r.EntClient.User.Query().Where(user.Tenant(tenant)).Where(user.UUID(uuid)).First(ctx)
	if err != nil {
		return nil, err
	}
	modelQuery := model.Update()
	if params.Data != nil {
		modelQuery.SetData((*interface{})(params.Data))
	}
	if params.IsSuper != nil {
		modelQuery.SetIsSuper(*params.IsSuper)
	}
	if params.Tenant != nil {
		modelQuery.SetTenant(string(*params.Tenant))
	}
	if params.ParentId != nil {
		if *params.ParentId == "" {
			modelQuery.ClearParent()
		} else {
			parent, err := r.EntClient.User.Query().Where(user.Tenant(model.Tenant)).Where(user.ID(cast.ToInt(params.ParentId))).First(ctx)
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

	if params.PermissionIds != nil {
		permissionIds := []int{}
		for _, v := range *params.PermissionIds {
			permissionIds = append(permissionIds, cast.ToInt(v))
		}
		permissions := r.EntClient.Permission.Query().Where(permission.Tenant(model.Tenant)).
			Where(permission.IDIn(permissionIds...)).AllX(ctx)
		modelQuery.ClearPermissions().AddPermissions(permissions...)
	}
	if params.GroupIds != nil {
		groupIds := []int{}
		for _, v := range *params.GroupIds {
			groupIds = append(groupIds, cast.ToInt(v))
		}
		groups := r.EntClient.Group.Query().Where(group.Tenant(model.Tenant)).
			Where(group.IDIn(groupIds...)).AllX(ctx)
		modelQuery.ClearGroups().AddGroups(groups...)
	}

	model, err = modelQuery.Save(ctx)
	if err != nil {
		return nil, err
	}

	return r.GetById(ctx, cast.ToString(model.ID))
}
