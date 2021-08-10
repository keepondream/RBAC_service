// Package rbac provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package rbac

import (
	"time"
)

// Defines values for ErrCode.
const (
	ErrCodeN422000 ErrCode = "422.000"

	ErrCodeN422111 ErrCode = "422.111"

	ErrCodeN422999 ErrCode = "422.999"
)

// Defines values for ItemMethod.
const (
	ItemMethodCONNECT ItemMethod = "CONNECT"

	ItemMethodDELETE ItemMethod = "DELETE"

	ItemMethodGET ItemMethod = "GET"

	ItemMethodHEAD ItemMethod = "HEAD"

	ItemMethodOPTIONS ItemMethod = "OPTIONS"

	ItemMethodPATCH ItemMethod = "PATCH"

	ItemMethodPOST ItemMethod = "POST"

	ItemMethodPUT ItemMethod = "PUT"

	ItemMethodTRACE ItemMethod = "TRACE"
)

// Defines values for Order.
const (
	Asc Order = "asc"

	Desc Order = "desc"
)

// 错误码
//
// |Code|Description|
// |----|----|
// |422.000|数据不存在|
// |422.111|数据已存在|
// |422.999|参数有误|
type ErrCode string

// Group defines model for Group.
type Group struct {
	// Embedded fields due to inline allOf schema

	// 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	CreatedAt time.Time `json:"created_at"`

	// 自定义json对象或者数组json
	Data ItemData `json:"data"`
	Id   string   `json:"id"`

	// 分组名称
	Name string `json:"name"`

	// 父级ID
	ParentId string `json:"parent_id"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant ItemTenant `json:"tenant"`

	// 节点或者分组的类型可自定义 例如 role:角色, menu:菜单, element:页面元素 ...等等
	Type ItemType `json:"type"`

	// 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt time.Time `json:"updated_at"`
}

// GroupInfoResponse defines model for GroupInfoResponse.
type GroupInfoResponse struct {
	// Embedded struct due to allOf(#/components/schemas/Group)
	Group `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Children []Group `json:"children"`
	Nodes    []Node  `json:"nodes"`
	Parent   *Group  `json:"parent,omitempty"`
}

// GroupListResponse defines model for GroupListResponse.
type GroupListResponse struct {
	Items []GroupInfoResponse `json:"items"`
	Total string              `json:"total"`
}

// 自定义json对象或者数组json
type ItemData interface{}

// URL请求方式 GET,HEAD,POST,PUT,PATCH,DELETE,CONNECT,OPTIONS,TRACE
type ItemMethod string

// ItemPolicy defines model for ItemPolicy.
type ItemPolicy struct {

	// URL请求方式 GET,HEAD,POST,PUT,PATCH,DELETE,CONNECT,OPTIONS,TRACE
	Method ItemMethod `json:"method"`

	// 权限ID
	PermissionId string `json:"permission_id"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant ItemTenant `json:"tenant"`

	// 路由path
	Uri string `json:"uri"`
}

// ItemRelation defines model for ItemRelation.
type ItemRelation struct {

	// ID
	Id string `json:"id"`

	// 名称
	Name string `json:"name"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant ItemTenant `json:"tenant"`

	// 节点或者分组的类型可自定义 例如 role:角色, menu:菜单, element:页面元素 ...等等
	Type *ItemType `json:"type,omitempty"`
}

// ItemRoute defines model for ItemRoute.
type ItemRoute struct {

	// 权限IDs
	PermissionIds []string     `json:"permission_ids"`
	Policies      []ItemPolicy `json:"policies"`
}

// 域标识,可自定义用于区分哪个平台使用
type ItemTenant string

// 节点或者分组的类型可自定义 例如 role:角色, menu:菜单, element:页面元素 ...等等
type ItemType string

// Node defines model for Node.
type Node struct {

	// 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	CreatedAt time.Time `json:"created_at"`

	// 自定义json对象或者数组json
	Data ItemData `json:"data"`
	Id   string   `json:"id"`

	// 节点名称
	Name string `json:"name"`

	// 父节点ID
	ParentId string `json:"parent_id"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant ItemTenant `json:"tenant"`

	// 节点或者分组的类型可自定义 例如 role:角色, menu:菜单, element:页面元素 ...等等
	Type ItemType `json:"type"`

	// 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt time.Time `json:"updated_at"`
}

// NodeInfoResponse defines model for NodeInfoResponse.
type NodeInfoResponse struct {
	// Embedded struct due to allOf(#/components/schemas/Node)
	Node `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Children    []Node       `json:"children"`
	Parent      *Node        `json:"parent,omitempty"`
	Permissions []Permission `json:"permissions"`
}

// NodeListResponse defines model for NodeListResponse.
type NodeListResponse struct {
	Items []NodeInfoResponse `json:"items"`
	Total string             `json:"total"`
}

// Permission defines model for Permission.
type Permission struct {

	// 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	CreatedAt time.Time `json:"created_at"`

	// 自定义json对象或者数组json
	Data ItemData `json:"data"`
	Id   string   `json:"id"`

	// 权限名称
	Name string `json:"name"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant ItemTenant `json:"tenant"`

	// 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt time.Time `json:"updated_at"`
}

// PermissionInfoResponse defines model for PermissionInfoResponse.
type PermissionInfoResponse struct {
	// Embedded struct due to allOf(#/components/schemas/Permission)
	Permission `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Routes []Route `json:"routes"`
}

// PermissionListResponse defines model for PermissionListResponse.
type PermissionListResponse struct {
	Items []PermissionInfoResponse `json:"items"`
	Total string                   `json:"total"`
}

// Route defines model for Route.
type Route struct {

	// 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	CreatedAt time.Time `json:"created_at"`

	// 自定义json对象或者数组json
	Data ItemData `json:"data"`

	// 路由ID
	Id string `json:"id"`

	// URL请求方式 GET,HEAD,POST,PUT,PATCH,DELETE,CONNECT,OPTIONS,TRACE
	Method ItemMethod `json:"method"`

	// 路由名称
	Name string `json:"name"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant ItemTenant `json:"tenant"`

	// 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt time.Time `json:"updated_at"`

	// 拦截URL
	Uri string `json:"uri"`
}

// RouteInfoResponse defines model for RouteInfoResponse.
type RouteInfoResponse Route

// RouteListResponse defines model for RouteListResponse.
type RouteListResponse struct {
	Items []Route `json:"items"`
	Total string  `json:"total"`
}

// User defines model for User.
type User struct {

	// 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	CreatedAt time.Time `json:"created_at"`

	// 自定义json对象或者数组json
	Data ItemData `json:"data"`
	Id   string   `json:"id"`

	// 是否为超级管理员,该标识意味着当前用户不需要鉴权,有系统最大权限
	IsSuper bool `json:"is_super"`

	// 父级ID
	ParentId string `json:"parent_id"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant ItemTenant `json:"tenant"`

	// 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	UpdatedAt time.Time `json:"updated_at"`

	// 用户唯一标识,由认证服务或者第三方服务提供的唯一标识
	Uuid string `json:"uuid"`
}

// UserAllRelationsResponse defines model for UserAllRelationsResponse.
type UserAllRelationsResponse struct {

	// 分组IDs
	GroupItems []ItemRelation `json:"group_items"`

	// 节点IDs
	NodeItems []ItemRelation `json:"node_items"`

	// 权限IDs
	PermissionItems []ItemRelation `json:"permission_items"`
}

// UserAllRoutesResponse defines model for UserAllRoutesResponse.
type UserAllRoutesResponse ItemRoute

// UserInfoResponse defines model for UserInfoResponse.
type UserInfoResponse struct {
	// Embedded struct due to allOf(#/components/schemas/User)
	User `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	Children    []User       `json:"children"`
	Groups      []Group      `json:"groups"`
	Nodes       []Node       `json:"nodes"`
	Parent      *User        `json:"parent,omitempty"`
	Permissions []Permission `json:"permissions"`
}

// UserListResponse defines model for UserListResponse.
type UserListResponse struct {
	Items []UserInfoResponse `json:"items"`
	Total string             `json:"total"`
}

// EndTime defines model for end_time.
type EndTime time.Time

// Order defines model for order.
type Order string

// Page defines model for page.
type Page string

// PerPage defines model for per_page.
type PerPage string

// Query defines model for query.
type Query string

// Sort defines model for sort.
type Sort string

// StartTime defines model for start_time.
type StartTime time.Time

// ErrResponse defines model for ErrResponse.
type ErrResponse struct {

	// 错误码
	//
	// |Code|Description|
	// |----|----|
	// |422.000|数据不存在|
	// |422.111|数据已存在|
	// |422.999|参数有误|
	Code *ErrCode `json:"code,omitempty"`

	// 错误字段
	Field *string `json:"field,omitempty"`

	// 错误描述
	Msg *string `json:"msg,omitempty"`
}

// GetGroupsParams defines parameters for GetGroups.
type GetGroupsParams struct {

	// 页码
	Page Page `json:"page"`

	// 分页数量 默认20, 最大100
	PerPage *PerPage `json:"per_page,omitempty"`

	// 排序方式 只支持 asc 或者 desc
	Order GetGroupsParamsOrder `json:"order"`

	// - 查询关键字以空格分割,多个维度用冒号连接关键字,每个关键字需要用encode `示例: 查询默认关键字为name "route1 route2"两个名称的路由, 域标识为 domain1 或者 domain2 的数据`  ``` http://host.com?query=${encodeURIComponent('route1')} ${encodeURIComponent('route2')} tenant:${encodeURIComponent('domain1')} tenant:${encodeURIComponent('domain2')} ```
	Query *Query `json:"query,omitempty"`

	// 排序字段 多个组合用逗号分隔 示例: id,name
	Sort *Sort `json:"sort,omitempty"`

	// 起始时间 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	StartTime *StartTime `json:"start_time,omitempty"`

	// 结束时间 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	EndTime *EndTime `json:"end_time,omitempty"`
}

// GetGroupsParamsOrder defines parameters for GetGroups.
type GetGroupsParamsOrder string

// PostGroupsJSONBody defines parameters for PostGroups.
type PostGroupsJSONBody struct {

	// 自定义json对象或者数组json
	Data *ItemData `json:"data,omitempty"`

	// 名称
	Name string `json:"name"`

	// 节点IDs
	NodeIds []string `json:"node_ids"`

	// 父级ID,非必填选项
	ParentId *string `json:"parent_id,omitempty"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant ItemTenant `json:"tenant"`

	// 节点或者分组的类型可自定义 例如 role:角色, menu:菜单, element:页面元素 ...等等
	Type ItemType `json:"type"`
}

// PatchGroupsIdJSONBody defines parameters for PatchGroupsId.
type PatchGroupsIdJSONBody struct {

	// 自定义json对象或者数组json
	Data *ItemData `json:"data,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 节点IDs
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 父级ID,非必填选项
	ParentId *string `json:"parent_id,omitempty"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant *ItemTenant `json:"tenant,omitempty"`

	// 节点或者分组的类型可自定义 例如 role:角色, menu:菜单, element:页面元素 ...等等
	Type *ItemType `json:"type,omitempty"`
}

// GetNodesParams defines parameters for GetNodes.
type GetNodesParams struct {

	// 页码
	Page Page `json:"page"`

	// 分页数量 默认20, 最大100
	PerPage *PerPage `json:"per_page,omitempty"`

	// 排序方式 只支持 asc 或者 desc
	Order GetNodesParamsOrder `json:"order"`

	// - 查询关键字以空格分割,多个维度用冒号连接关键字,每个关键字需要用encode `示例: 查询默认关键字为name "route1 route2"两个名称的路由, 域标识为 domain1 或者 domain2 的数据`  ``` http://host.com?query=${encodeURIComponent('route1')} ${encodeURIComponent('route2')} tenant:${encodeURIComponent('domain1')} tenant:${encodeURIComponent('domain2')} ```
	Query *Query `json:"query,omitempty"`

	// 排序字段 多个组合用逗号分隔 示例: id,name
	Sort *Sort `json:"sort,omitempty"`

	// 起始时间 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	StartTime *StartTime `json:"start_time,omitempty"`

	// 结束时间 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	EndTime *EndTime `json:"end_time,omitempty"`
}

// GetNodesParamsOrder defines parameters for GetNodes.
type GetNodesParamsOrder string

// PostNodesJSONBody defines parameters for PostNodes.
type PostNodesJSONBody struct {

	// 自定义json对象或者数组json
	Data *ItemData `json:"data,omitempty"`

	// 名称
	Name string `json:"name"`

	// 父级ID,非必填选项
	ParentId *string `json:"parent_id,omitempty"`

	// 权限IDS
	PermissionIds []string `json:"permission_ids"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant ItemTenant `json:"tenant"`

	// 节点或者分组的类型可自定义 例如 role:角色, menu:菜单, element:页面元素 ...等等
	Type ItemType `json:"type"`
}

// PatchNodesIdJSONBody defines parameters for PatchNodesId.
type PatchNodesIdJSONBody struct {

	// 自定义json对象或者数组json
	Data *ItemData `json:"data,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 父级ID
	ParentId *string `json:"parent_id,omitempty"`

	// 权限ID
	PermissionIds *[]string `json:"permission_ids,omitempty"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant *ItemTenant `json:"tenant,omitempty"`

	// 节点或者分组的类型可自定义 例如 role:角色, menu:菜单, element:页面元素 ...等等
	Type *ItemType `json:"type,omitempty"`
}

// GetPermissionsParams defines parameters for GetPermissions.
type GetPermissionsParams struct {

	// 页码
	Page Page `json:"page"`

	// 分页数量 默认20, 最大100
	PerPage *PerPage `json:"per_page,omitempty"`

	// 排序方式 只支持 asc 或者 desc
	Order GetPermissionsParamsOrder `json:"order"`

	// - 查询关键字以空格分割,多个维度用冒号连接关键字,每个关键字需要用encode `示例: 查询默认关键字为name "route1 route2"两个名称的路由, 域标识为 domain1 或者 domain2 的数据`  ``` http://host.com?query=${encodeURIComponent('route1')} ${encodeURIComponent('route2')} tenant:${encodeURIComponent('domain1')} tenant:${encodeURIComponent('domain2')} ```
	Query *Query `json:"query,omitempty"`

	// 排序字段 多个组合用逗号分隔 示例: id,name
	Sort *Sort `json:"sort,omitempty"`

	// 起始时间 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	StartTime *StartTime `json:"start_time,omitempty"`

	// 结束时间 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	EndTime *EndTime `json:"end_time,omitempty"`
}

// GetPermissionsParamsOrder defines parameters for GetPermissions.
type GetPermissionsParamsOrder string

// PostPermissionsJSONBody defines parameters for PostPermissions.
type PostPermissionsJSONBody struct {

	// 自定义json对象或者数组json
	Data ItemData `json:"data"`

	// 权限名称
	Name string `json:"name"`

	// 权限对应的路由IDS
	RouteIds []string `json:"route_ids"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant ItemTenant `json:"tenant"`
}

// PatchPermissionsIdJSONBody defines parameters for PatchPermissionsId.
type PatchPermissionsIdJSONBody struct {

	// 自定义json对象或者数组json
	Data *ItemData `json:"data,omitempty"`

	// 权限名称
	Name *string `json:"name,omitempty"`

	// 权限绑定的相关路由ID,字段不传则不更改
	RouteIds *[]string `json:"route_ids,omitempty"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant *ItemTenant `json:"tenant,omitempty"`
}

// GetRoutesParams defines parameters for GetRoutes.
type GetRoutesParams struct {

	// 页码
	Page Page `json:"page"`

	// 分页数量 默认20, 最大100
	PerPage *PerPage `json:"per_page,omitempty"`

	// 排序方式 只支持 asc 或者 desc
	Order GetRoutesParamsOrder `json:"order"`

	// 排序字段 多个组合用逗号分隔 示例: id,name
	Sort *Sort `json:"sort,omitempty"`

	// - 查询关键字以空格分割,多个维度用冒号连接关键字,每个关键字需要用encode `示例: 查询默认关键字为name "route1 route2"两个名称的路由, 域标识为 domain1 或者 domain2 的数据`  ``` http://host.com?query=${encodeURIComponent('route1')} ${encodeURIComponent('route2')} tenant:${encodeURIComponent('domain1')} tenant:${encodeURIComponent('domain2')} ```
	Query *Query `json:"query,omitempty"`

	// 起始时间 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	StartTime *StartTime `json:"start_time,omitempty"`

	// 结束时间 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	EndTime *EndTime `json:"end_time,omitempty"`
}

// GetRoutesParamsOrder defines parameters for GetRoutes.
type GetRoutesParamsOrder string

// PostRoutesJSONBody defines parameters for PostRoutes.
type PostRoutesJSONBody struct {

	// 自定义json对象或者数组json
	Data ItemData `json:"data"`

	// URL请求方式 GET,HEAD,POST,PUT,PATCH,DELETE,CONNECT,OPTIONS,TRACE
	Method ItemMethod `json:"method"`

	// 路由名称
	Name string `json:"name"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant ItemTenant `json:"tenant"`

	// 路由路径,拦截URL,path 例如: /users/:id
	Uri string `json:"uri"`
}

// PatchRoutesIdJSONBody defines parameters for PatchRoutesId.
type PatchRoutesIdJSONBody struct {

	// 自定义json对象或者数组json
	Data *ItemData `json:"data,omitempty"`

	// URL请求方式 GET,HEAD,POST,PUT,PATCH,DELETE,CONNECT,OPTIONS,TRACE
	Method *ItemMethod `json:"method,omitempty"`

	// 路由名称
	Name *string `json:"name,omitempty"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant *ItemTenant `json:"tenant,omitempty"`

	// 路由路径,path,拦截URL
	Uri *string `json:"uri,omitempty"`
}

// GetUsersParams defines parameters for GetUsers.
type GetUsersParams struct {

	// 页码
	Page Page `json:"page"`

	// 分页数量 默认20, 最大100
	PerPage *PerPage `json:"per_page,omitempty"`

	// 排序方式 只支持 asc 或者 desc
	Order GetUsersParamsOrder `json:"order"`

	// - 查询关键字以空格分割,多个维度用冒号连接关键字,每个关键字需要用encode `示例: 查询默认关键字为name "route1 route2"两个名称的路由, 域标识为 domain1 或者 domain2 的数据`  ``` http://host.com?query=${encodeURIComponent('route1')} ${encodeURIComponent('route2')} tenant:${encodeURIComponent('domain1')} tenant:${encodeURIComponent('domain2')} ```
	Query *Query `json:"query,omitempty"`

	// 排序字段 多个组合用逗号分隔 示例: id,name
	Sort *Sort `json:"sort,omitempty"`

	// 起始时间 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	StartTime *StartTime `json:"start_time,omitempty"`

	// 结束时间 零时区时间格式: YYYY-MM-DDTHH:MM:SSZ
	EndTime *EndTime `json:"end_time,omitempty"`
}

// GetUsersParamsOrder defines parameters for GetUsers.
type GetUsersParamsOrder string

// PostUsersJSONBody defines parameters for PostUsers.
type PostUsersJSONBody struct {

	// 自定义json对象或者数组json
	Data *ItemData `json:"data,omitempty"`

	// 绑定分组IDS,可以为 角色组,菜单组页面元素组等等
	GroupIds []string `json:"group_ids"`

	// 是否为超级管理员,该标识意味着当前用户不需要鉴权,有系统最大权限
	IsSuper bool `json:"is_super"`

	// 绑定节点IDs, 可以为 角色ID,菜单ID,页面元素ID等等
	NodeIds []string `json:"node_ids"`

	// 父级ID
	ParentId *string `json:"parent_id,omitempty"`

	// 绑定权限IDS
	PermissionIds []string `json:"permission_ids"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant ItemTenant `json:"tenant"`

	// 用户唯一标识,由认证服务或者第三方服务提供的唯一标识
	Uuid string `json:"uuid"`
}

// PatchUsersUuidTenantJSONBody defines parameters for PatchUsersUuidTenant.
type PatchUsersUuidTenantJSONBody struct {

	// 自定义json对象或者数组json
	Data *ItemData `json:"data,omitempty"`

	// 绑定分组IDS,可以为 角色组,菜单组页面元素组等等
	GroupIds *[]string `json:"group_ids,omitempty"`

	// 是否为超级管理员,该标识意味着当前用户不需要鉴权,有系统最大权限
	IsSuper *bool `json:"is_super,omitempty"`

	// 绑定节点IDs, 可以为 角色ID,菜单ID,页面元素ID等等
	NodeIds *[]string `json:"node_ids,omitempty"`

	// 父级ID
	ParentId *string `json:"parent_id,omitempty"`

	// 绑定权限IDS
	PermissionIds *[]string `json:"permission_ids,omitempty"`

	// 域标识,可自定义用于区分哪个平台使用
	Tenant *ItemTenant `json:"tenant,omitempty"`
}

// PostGroupsJSONRequestBody defines body for PostGroups for application/json ContentType.
type PostGroupsJSONRequestBody PostGroupsJSONBody

// PatchGroupsIdJSONRequestBody defines body for PatchGroupsId for application/json ContentType.
type PatchGroupsIdJSONRequestBody PatchGroupsIdJSONBody

// PostNodesJSONRequestBody defines body for PostNodes for application/json ContentType.
type PostNodesJSONRequestBody PostNodesJSONBody

// PatchNodesIdJSONRequestBody defines body for PatchNodesId for application/json ContentType.
type PatchNodesIdJSONRequestBody PatchNodesIdJSONBody

// PostPermissionsJSONRequestBody defines body for PostPermissions for application/json ContentType.
type PostPermissionsJSONRequestBody PostPermissionsJSONBody

// PatchPermissionsIdJSONRequestBody defines body for PatchPermissionsId for application/json ContentType.
type PatchPermissionsIdJSONRequestBody PatchPermissionsIdJSONBody

// PostRoutesJSONRequestBody defines body for PostRoutes for application/json ContentType.
type PostRoutesJSONRequestBody PostRoutesJSONBody

// PatchRoutesIdJSONRequestBody defines body for PatchRoutesId for application/json ContentType.
type PatchRoutesIdJSONRequestBody PatchRoutesIdJSONBody

// PostUsersJSONRequestBody defines body for PostUsers for application/json ContentType.
type PostUsersJSONRequestBody PostUsersJSONBody

// PatchUsersUuidTenantJSONRequestBody defines body for PatchUsersUuidTenant for application/json ContentType.
type PatchUsersUuidTenantJSONRequestBody PatchUsersUuidTenantJSONBody
