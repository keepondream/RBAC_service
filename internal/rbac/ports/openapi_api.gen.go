// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.1 DO NOT EDIT.
package ports

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/go-chi/chi/v5"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 分组列表
	// (GET /groups)
	GetGroups(w http.ResponseWriter, r *http.Request, params GetGroupsParams)
	// 创建分组
	// (POST /groups)
	PostGroups(w http.ResponseWriter, r *http.Request)
	// 删除分组
	// (DELETE /groups/{id})
	DeleteGroupsId(w http.ResponseWriter, r *http.Request, id string)
	// 获取分组详情
	// (GET /groups/{id})
	GetGroupsId(w http.ResponseWriter, r *http.Request, id string)
	// 更新分组信息
	// (PATCH /groups/{id})
	PatchGroupsId(w http.ResponseWriter, r *http.Request, id string)
	// 节点列表
	// (GET /nodes)
	GetNodes(w http.ResponseWriter, r *http.Request, params GetNodesParams)
	// 创建节点
	// (POST /nodes)
	PostNodes(w http.ResponseWriter, r *http.Request)
	// 删除节点
	// (DELETE /nodes/{id})
	DeleteNodesId(w http.ResponseWriter, r *http.Request, id string)
	// 获取节点详情
	// (GET /nodes/{id})
	GetNodesId(w http.ResponseWriter, r *http.Request, id string)
	// 更新节点信息
	// (PATCH /nodes/{id})
	PatchNodesId(w http.ResponseWriter, r *http.Request, id string)
	// 权限列表
	// (GET /permissions)
	GetPermissions(w http.ResponseWriter, r *http.Request, params GetPermissionsParams)
	// 创建权限
	// (POST /permissions)
	PostPermissions(w http.ResponseWriter, r *http.Request)
	// 删除权限
	// (DELETE /permissions/{id})
	DeletePermissionsId(w http.ResponseWriter, r *http.Request, id string)
	// 获取权限详情
	// (GET /permissions/{id})
	GetPermissionsId(w http.ResponseWriter, r *http.Request, id string)
	// 更新权限信息
	// (PATCH /permissions/{id})
	PatchPermissionsId(w http.ResponseWriter, r *http.Request, id string)
	// 路由列表
	// (GET /routes)
	GetRoutes(w http.ResponseWriter, r *http.Request, params GetRoutesParams)
	// 创建路由
	// (POST /routes)
	PostRoutes(w http.ResponseWriter, r *http.Request)
	// 删除路由
	// (DELETE /routes/{id})
	DeleteRoutesId(w http.ResponseWriter, r *http.Request, id string)
	// 路由详情
	// (GET /routes/{id})
	GetRoutesId(w http.ResponseWriter, r *http.Request, id string)
	// 修改路由信息
	// (PATCH /routes/{id})
	PatchRoutesId(w http.ResponseWriter, r *http.Request, id string)
	// 绑定用户列表
	// (GET /users)
	GetUsers(w http.ResponseWriter, r *http.Request, params GetUsersParams)
	// 绑定用户
	// (POST /users)
	PostUsers(w http.ResponseWriter, r *http.Request)
	// 删除绑定用户
	// (DELETE /users/{uuid}/{tenant})
	DeleteUsersUuidTenant(w http.ResponseWriter, r *http.Request, uuid string, tenant string)
	// 获取绑定用户详情(包含拥有的权限,角色,子级用户,角色组,菜单组等等)
	// (GET /users/{uuid}/{tenant})
	GetUsersUuidTenant(w http.ResponseWriter, r *http.Request, uuid string, tenant string)
	// 更新绑定用户(角色,权限,子级,父级,分组等等)
	// (PATCH /users/{uuid}/{tenant})
	PatchUsersUuidTenant(w http.ResponseWriter, r *http.Request, uuid string, tenant string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// GetGroups operation middleware
func (siw *ServerInterfaceWrapper) GetGroups(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetGroupsParams

	// ------------- Required query parameter "page" -------------
	if paramValue := r.URL.Query().Get("page"); paramValue != "" {

	} else {
		http.Error(w, "Query argument page is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page", r.URL.Query(), &params.Page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter page: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "per_page" -------------
	if paramValue := r.URL.Query().Get("per_page"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "per_page", r.URL.Query(), &params.PerPage)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter per_page: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "order" -------------
	if paramValue := r.URL.Query().Get("order"); paramValue != "" {

	} else {
		http.Error(w, "Query argument order is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "order", r.URL.Query(), &params.Order)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter order: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "query" -------------
	if paramValue := r.URL.Query().Get("query"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "query", r.URL.Query(), &params.Query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter query: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "sort" -------------
	if paramValue := r.URL.Query().Get("sort"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sort", r.URL.Query(), &params.Sort)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter sort: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "start_time" -------------
	if paramValue := r.URL.Query().Get("start_time"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "start_time", r.URL.Query(), &params.StartTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter start_time: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "end_time" -------------
	if paramValue := r.URL.Query().Get("end_time"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "end_time", r.URL.Query(), &params.EndTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter end_time: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetGroups(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostGroups operation middleware
func (siw *ServerInterfaceWrapper) PostGroups(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostGroups(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteGroupsId operation middleware
func (siw *ServerInterfaceWrapper) DeleteGroupsId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteGroupsId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetGroupsId operation middleware
func (siw *ServerInterfaceWrapper) GetGroupsId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetGroupsId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PatchGroupsId operation middleware
func (siw *ServerInterfaceWrapper) PatchGroupsId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PatchGroupsId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetNodes operation middleware
func (siw *ServerInterfaceWrapper) GetNodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetNodesParams

	// ------------- Required query parameter "page" -------------
	if paramValue := r.URL.Query().Get("page"); paramValue != "" {

	} else {
		http.Error(w, "Query argument page is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page", r.URL.Query(), &params.Page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter page: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "per_page" -------------
	if paramValue := r.URL.Query().Get("per_page"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "per_page", r.URL.Query(), &params.PerPage)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter per_page: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "order" -------------
	if paramValue := r.URL.Query().Get("order"); paramValue != "" {

	} else {
		http.Error(w, "Query argument order is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "order", r.URL.Query(), &params.Order)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter order: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "query" -------------
	if paramValue := r.URL.Query().Get("query"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "query", r.URL.Query(), &params.Query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter query: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "sort" -------------
	if paramValue := r.URL.Query().Get("sort"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sort", r.URL.Query(), &params.Sort)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter sort: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "start_time" -------------
	if paramValue := r.URL.Query().Get("start_time"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "start_time", r.URL.Query(), &params.StartTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter start_time: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "end_time" -------------
	if paramValue := r.URL.Query().Get("end_time"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "end_time", r.URL.Query(), &params.EndTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter end_time: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetNodes(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostNodes operation middleware
func (siw *ServerInterfaceWrapper) PostNodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostNodes(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteNodesId operation middleware
func (siw *ServerInterfaceWrapper) DeleteNodesId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteNodesId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetNodesId operation middleware
func (siw *ServerInterfaceWrapper) GetNodesId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetNodesId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PatchNodesId operation middleware
func (siw *ServerInterfaceWrapper) PatchNodesId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PatchNodesId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetPermissions operation middleware
func (siw *ServerInterfaceWrapper) GetPermissions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPermissionsParams

	// ------------- Required query parameter "page" -------------
	if paramValue := r.URL.Query().Get("page"); paramValue != "" {

	} else {
		http.Error(w, "Query argument page is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page", r.URL.Query(), &params.Page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter page: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "per_page" -------------
	if paramValue := r.URL.Query().Get("per_page"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "per_page", r.URL.Query(), &params.PerPage)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter per_page: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "order" -------------
	if paramValue := r.URL.Query().Get("order"); paramValue != "" {

	} else {
		http.Error(w, "Query argument order is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "order", r.URL.Query(), &params.Order)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter order: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "query" -------------
	if paramValue := r.URL.Query().Get("query"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "query", r.URL.Query(), &params.Query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter query: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "sort" -------------
	if paramValue := r.URL.Query().Get("sort"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sort", r.URL.Query(), &params.Sort)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter sort: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "start_time" -------------
	if paramValue := r.URL.Query().Get("start_time"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "start_time", r.URL.Query(), &params.StartTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter start_time: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "end_time" -------------
	if paramValue := r.URL.Query().Get("end_time"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "end_time", r.URL.Query(), &params.EndTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter end_time: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPermissions(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostPermissions operation middleware
func (siw *ServerInterfaceWrapper) PostPermissions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostPermissions(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeletePermissionsId operation middleware
func (siw *ServerInterfaceWrapper) DeletePermissionsId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeletePermissionsId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetPermissionsId operation middleware
func (siw *ServerInterfaceWrapper) GetPermissionsId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetPermissionsId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PatchPermissionsId operation middleware
func (siw *ServerInterfaceWrapper) PatchPermissionsId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PatchPermissionsId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetRoutes operation middleware
func (siw *ServerInterfaceWrapper) GetRoutes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetRoutesParams

	// ------------- Required query parameter "page" -------------
	if paramValue := r.URL.Query().Get("page"); paramValue != "" {

	} else {
		http.Error(w, "Query argument page is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page", r.URL.Query(), &params.Page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter page: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "per_page" -------------
	if paramValue := r.URL.Query().Get("per_page"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "per_page", r.URL.Query(), &params.PerPage)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter per_page: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "order" -------------
	if paramValue := r.URL.Query().Get("order"); paramValue != "" {

	} else {
		http.Error(w, "Query argument order is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "order", r.URL.Query(), &params.Order)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter order: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "sort" -------------
	if paramValue := r.URL.Query().Get("sort"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sort", r.URL.Query(), &params.Sort)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter sort: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "query" -------------
	if paramValue := r.URL.Query().Get("query"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "query", r.URL.Query(), &params.Query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter query: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "start_time" -------------
	if paramValue := r.URL.Query().Get("start_time"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "start_time", r.URL.Query(), &params.StartTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter start_time: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "end_time" -------------
	if paramValue := r.URL.Query().Get("end_time"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "end_time", r.URL.Query(), &params.EndTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter end_time: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRoutes(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostRoutes operation middleware
func (siw *ServerInterfaceWrapper) PostRoutes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostRoutes(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteRoutesId operation middleware
func (siw *ServerInterfaceWrapper) DeleteRoutesId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteRoutesId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetRoutesId operation middleware
func (siw *ServerInterfaceWrapper) GetRoutesId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetRoutesId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PatchRoutesId operation middleware
func (siw *ServerInterfaceWrapper) PatchRoutesId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameter("simple", false, "id", chi.URLParam(r, "id"), &id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter id: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PatchRoutesId(w, r, id)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetUsers operation middleware
func (siw *ServerInterfaceWrapper) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUsersParams

	// ------------- Required query parameter "page" -------------
	if paramValue := r.URL.Query().Get("page"); paramValue != "" {

	} else {
		http.Error(w, "Query argument page is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "page", r.URL.Query(), &params.Page)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter page: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "per_page" -------------
	if paramValue := r.URL.Query().Get("per_page"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "per_page", r.URL.Query(), &params.PerPage)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter per_page: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Required query parameter "order" -------------
	if paramValue := r.URL.Query().Get("order"); paramValue != "" {

	} else {
		http.Error(w, "Query argument order is required, but not found", http.StatusBadRequest)
		return
	}

	err = runtime.BindQueryParameter("form", true, true, "order", r.URL.Query(), &params.Order)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter order: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "query" -------------
	if paramValue := r.URL.Query().Get("query"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "query", r.URL.Query(), &params.Query)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter query: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "sort" -------------
	if paramValue := r.URL.Query().Get("sort"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "sort", r.URL.Query(), &params.Sort)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter sort: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "start_time" -------------
	if paramValue := r.URL.Query().Get("start_time"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "start_time", r.URL.Query(), &params.StartTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter start_time: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "end_time" -------------
	if paramValue := r.URL.Query().Get("end_time"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "end_time", r.URL.Query(), &params.EndTime)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter end_time: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsers(w, r, params)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PostUsers operation middleware
func (siw *ServerInterfaceWrapper) PostUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostUsers(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// DeleteUsersUuidTenant operation middleware
func (siw *ServerInterfaceWrapper) DeleteUsersUuidTenant(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "uuid" -------------
	var uuid string

	err = runtime.BindStyledParameter("simple", false, "uuid", chi.URLParam(r, "uuid"), &uuid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter uuid: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "tenant" -------------
	var tenant string

	err = runtime.BindStyledParameter("simple", false, "tenant", chi.URLParam(r, "tenant"), &tenant)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter tenant: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteUsersUuidTenant(w, r, uuid, tenant)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// GetUsersUuidTenant operation middleware
func (siw *ServerInterfaceWrapper) GetUsersUuidTenant(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "uuid" -------------
	var uuid string

	err = runtime.BindStyledParameter("simple", false, "uuid", chi.URLParam(r, "uuid"), &uuid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter uuid: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "tenant" -------------
	var tenant string

	err = runtime.BindStyledParameter("simple", false, "tenant", chi.URLParam(r, "tenant"), &tenant)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter tenant: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsersUuidTenant(w, r, uuid, tenant)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// PatchUsersUuidTenant operation middleware
func (siw *ServerInterfaceWrapper) PatchUsersUuidTenant(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "uuid" -------------
	var uuid string

	err = runtime.BindStyledParameter("simple", false, "uuid", chi.URLParam(r, "uuid"), &uuid)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter uuid: %s", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "tenant" -------------
	var tenant string

	err = runtime.BindStyledParameter("simple", false, "tenant", chi.URLParam(r, "tenant"), &tenant)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid format for parameter tenant: %s", err), http.StatusBadRequest)
		return
	}

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PatchUsersUuidTenant(w, r, uuid, tenant)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/groups", wrapper.GetGroups)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/groups", wrapper.PostGroups)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/groups/{id}", wrapper.DeleteGroupsId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/groups/{id}", wrapper.GetGroupsId)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/groups/{id}", wrapper.PatchGroupsId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/nodes", wrapper.GetNodes)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/nodes", wrapper.PostNodes)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/nodes/{id}", wrapper.DeleteNodesId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/nodes/{id}", wrapper.GetNodesId)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/nodes/{id}", wrapper.PatchNodesId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/permissions", wrapper.GetPermissions)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/permissions", wrapper.PostPermissions)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/permissions/{id}", wrapper.DeletePermissionsId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/permissions/{id}", wrapper.GetPermissionsId)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/permissions/{id}", wrapper.PatchPermissionsId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/routes", wrapper.GetRoutes)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/routes", wrapper.PostRoutes)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/routes/{id}", wrapper.DeleteRoutesId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/routes/{id}", wrapper.GetRoutesId)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/routes/{id}", wrapper.PatchRoutesId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/users", wrapper.GetUsers)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/users", wrapper.PostUsers)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/users/{uuid}/{tenant}", wrapper.DeleteUsersUuidTenant)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/users/{uuid}/{tenant}", wrapper.GetUsersUuidTenant)
	})
	r.Group(func(r chi.Router) {
		r.Patch(options.BaseURL+"/users/{uuid}/{tenant}", wrapper.PatchUsersUuidTenant)
	})

	return r
}
