package service

import (
	"context"
	"fmt"

	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/group"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/node"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/permission"
	"github.com/keepondream/RBAC_service/internal/rbac/adapters/ent/user"
	"github.com/spf13/cast"
)

const (
	NODEPREFIX  = "node:"  // casbin 中 node节点前缀
	GROUPPREFIX = "group:" // casbin 中 分组前缀
	USERPREFIX  = "user:"  // casbin 中 用户前缀
)

// SyncPolicyForPermission 同步权限对应的策略,这是鉴权最底层的核心, policy 策略层
func (s *Service) SyncPolicyForPermission(ctx context.Context, permission_id string) error {
	// 清除当前权限所有的策略
	s.CasbinE.RemoveFilteredPolicy(0, permission_id)
	permissionModel, err := s.EntClient.Permission.Query().WithRoutes().Where(permission.ID(cast.ToInt(permission_id))).First(ctx)
	if err != nil {
		return err
	}

	for _, routeModel := range permissionModel.Edges.Routes {
		// 组装策略进行添加
		s.CasbinE.AddPolicy(
			cast.ToString(permissionModel.ID),
			permissionModel.Tenant,
			routeModel.URI,
			routeModel.Method.String(),
		)
	}

	return nil
}

// SyncCasbinForNode 同步节点对应的权限和父级关系 , 凌驾于 policy上一层, 用于描述 节点与权限的关系,即为鉴权提供 节点所有的策略使用
func (s *Service) SyncCasbinForNode(ctx context.Context, node_id, tenant string) error {
	// 组装node对应的key,因为 node ID 会与 group id 重复,这里为了区分增加 固定前缀
	prefix := fmt.Sprintf("%s%s", NODEPREFIX, node_id)

	// 清除节点所有的权限,包含所属关系  , 第一参数为 节点ID, 第二参数为 域,
	s.CasbinE.DeleteRolesForUserInDomain(prefix, tenant)

	// 清除节点的children关系
	s.CasbinE.RemoveFilteredGroupingPolicy(1, prefix, tenant)

	// 获取节点和对应域下的所有权限
	nodeModel, err := s.EntClient.Node.Query().
		WithParent(func(nq *ent.NodeQuery) {
			node.Tenant(tenant)
		}).
		WithPermissions(func(pq *ent.PermissionQuery) {
			permission.Tenant(tenant)
		}).WithChildren(func(nq *ent.NodeQuery) {
		node.Tenant(tenant)
	}).Where(node.ID(cast.ToInt(node_id))).First(ctx)

	if err != nil {
		return err
	}

	// 获取节点的所有权限
	for _, permissionModel := range nodeModel.Edges.Permissions {
		// 为节点分配对应域的权限  第一参数为 节点ID , 第二参数为 权限ID, 第三参数为域
		s.CasbinE.AddRoleForUserInDomain(prefix, cast.ToString(permissionModel.ID), tenant)
	}

	// 绑定节点父级的关系
	if nodeModel.Edges.Parent != nil {
		// 为几点分配对应父级的权限 第一参数为 节点ID , 第二参数为 父级节点ID, 第三参数为域
		parentPrefix := fmt.Sprintf("%s%d", NODEPREFIX, nodeModel.Edges.Parent.ID)
		s.CasbinE.AddRoleForUserInDomain(prefix, parentPrefix, tenant)
	}

	// 绑定节点children关系
	for _, childrenModel := range nodeModel.Edges.Children {
		childrenPrefix := fmt.Sprintf("%s%d", NODEPREFIX, childrenModel.ID)
		s.CasbinE.AddRoleForUserInDomain(childrenPrefix, prefix, tenant)
	}

	return nil
}

// SyncCasbinForGroup 同步分组与节点之间的关系 , 在节点上再抽象一层分组,分组与分组的关系
func (s *Service) SyncCasbinForGroup(ctx context.Context, group_id, tenant string) error {
	prefix := fmt.Sprintf("%s%s", GROUPPREFIX, group_id)
	// 清除分组的所属关系  , 第一参数为 分组ID, 第二参数为 域,
	s.CasbinE.DeleteRolesForUserInDomain(prefix, tenant)

	// 清除分组的children关系
	s.CasbinE.RemoveFilteredGroupingPolicy(1, prefix)

	// 获取分组的所有子分组
	groupModel, err := s.EntClient.Group.Query().WithParent(func(gq *ent.GroupQuery) {
		group.Tenant(tenant)
	}).WithNodes(func(nq *ent.NodeQuery) {
		node.Tenant(tenant)
	}).WithChildren(func(gq *ent.GroupQuery) {
		group.Tenant(tenant)
	}).Where(group.ID(cast.ToInt(group_id))).First(ctx)

	if err != nil {
		return err
	}
	// 绑定分组父级的关系
	if groupModel.Edges.Parent != nil {
		parentPrefix := fmt.Sprintf("%s%d", GROUPPREFIX, groupModel.Edges.Parent.ID)
		s.CasbinE.AddRoleForUserInDomain(prefix, parentPrefix, tenant)
	}

	// 绑定分组所属节点的关系
	for _, nodeModel := range groupModel.Edges.Nodes {
		nodePrefix := fmt.Sprintf("%s%d", NODEPREFIX, nodeModel.ID)
		s.CasbinE.AddRoleForUserInDomain(prefix, nodePrefix, tenant)
	}

	// 绑定分组所有children的关系
	for _, childrenModel := range groupModel.Edges.Children {
		childrenPrefix := fmt.Sprintf("%s%d", GROUPPREFIX, childrenModel.ID)
		s.CasbinE.AddRoleForUserInDomain(childrenPrefix, prefix, tenant)
	}

	return nil
}

// SyncCasbinForUser 同步用户的权限,角色,菜单,角色组,父级,子级等等...所有关系
func (s *Service) SyncCasbinForUser(ctx context.Context, uuid, tenant string) error {
	prefix := fmt.Sprintf("%s%s", USERPREFIX, uuid)
	// 清除用户的所属关系(拥有关系即拥有的权限,拥有的角色,拥有的菜单,拥有的角色组,拥有的菜单组等等...)
	s.CasbinE.DeleteRolesForUserInDomain(prefix, tenant)

	// 清除用户的child关系(子级关系)
	s.CasbinE.RemoveFilteredPolicy(1, prefix)

	// 获取用户的所有关系图
	userModel, err := s.EntClient.User.Query().
		WithParent(func(uq *ent.UserQuery) {
			user.Tenant(tenant)
		}).
		WithChildren(func(uq *ent.UserQuery) {
			user.Tenant(tenant)
		}).
		WithNodes(func(nq *ent.NodeQuery) {
			node.Tenant(tenant)
		}).
		WithGroups(func(gq *ent.GroupQuery) {
			group.Tenant(tenant)
		}).
		WithPermissions(func(pq *ent.PermissionQuery) {
			permission.Tenant(tenant)
		}).Where(user.UUID(uuid)).Where(user.Tenant(tenant)).First(ctx)
	if err != nil {
		return err
	}

	// 绑定父级关系
	if userModel.Edges.Parent != nil {
		parentPrefix := fmt.Sprintf("%s%d", USERPREFIX, userModel.Edges.Parent.ID)
		s.CasbinE.AddRoleForUserInDomain(prefix, parentPrefix, tenant)
	}

	// 绑定子级关系
	for _, childrenModel := range userModel.Edges.Children {
		childrenPrefix := fmt.Sprintf("%s%d", USERPREFIX, childrenModel.ID)
		s.CasbinE.AddRoleForUserInDomain(childrenPrefix, prefix, tenant)
	}

	// 绑定节点关系
	for _, nodeModel := range userModel.Edges.Nodes {
		nodePrefix := fmt.Sprintf("%s%d", NODEPREFIX, nodeModel.ID)
		s.CasbinE.AddRoleForUserInDomain(prefix, nodePrefix, tenant)
	}

	// 绑定权限关系
	for _, permissionModel := range userModel.Edges.Permissions {
		s.CasbinE.AddRoleForUserInDomain(prefix, cast.ToString(permissionModel.ID), tenant)
	}

	// 绑定分组关系
	for _, groupModel := range userModel.Edges.Groups {
		groupPrefix := fmt.Sprintf("%s%d", GROUPPREFIX, groupModel.ID)
		s.CasbinE.AddRoleForUserInDomain(prefix, groupPrefix, tenant)
	}

	return nil
}

// GetAllPolicyForPrefix 获取(节点/分组/权限) 下所有递归,继承等隐式路由策略
func (s *Service) GetAllPolicyForPrefix(prefix, tenant string) ([][]string, error) {
	// 获取节点or分组or权限 下所有路由策略 , 递归,包含对应继承关系的权限, 递归查到底  ***
	return s.CasbinE.GetImplicitPermissionsForUser(prefix, tenant)
}

// GetAllRelationsForPrefix 获取(节点/分组/权限) 下所有递归关系,不包含路由策略
func (s *Service) GetAllRelationsForPrefix(prefix, tenant string) ([]string, error) {
	// 获取当前用户继承的所有一级继承,包一级继承下的所有继承的继承
	return s.CasbinE.GetImplicitRolesForUser(prefix, tenant)
}
