package ports

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/keepondream/RBAC_service/internal/common/utils"
)

type Userer interface {
	Create(ctx context.Context, params PostUsersJSONBody) (*UserInfoResponse, error)
	IsUnique(ctx context.Context, tenant, uuid string) error
	GetById(ctx context.Context, id string) (*UserInfoResponse, error)
	GetByUuid(ctx context.Context, tenant, uuid string) (*UserInfoResponse, error)
	DeleteByUuid(ctx context.Context, tenant, uuid string) error
	List(ctx context.Context, params GetUsersParams) (*UserListResponse, error)
	Update(ctx context.Context, params PatchUsersUuidTenantJSONBody, tenant, uuid string) (*UserInfoResponse, error)
	GetAllRoutes(ctx context.Context, tenant, uuid string) (*UserAllRoutesResponse, error)
	GetAllRelations(ctx context.Context, tenant, uuid string) (*UserAllRelationsResponse, error)
}

// 绑定用户列表
// (GET /users)
func (h *HttpServer) GetUsers(w http.ResponseWriter, r *http.Request, params GetUsersParams) {
	res, err := h.UserService.List(r.Context(), params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 200, utils.WithData(res))
}

// 绑定用户
// (POST /users)
func (h *HttpServer) PostUsers(w http.ResponseWriter, r *http.Request) {
	var params PostUsersJSONRequestBody
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}

	err = h.UserService.IsUnique(r.Context(), string(params.Tenant), params.Uuid)
	if err != nil {
		utils.Render(w, r, 422, utils.WithCode(string(ErrCodeN422111)), utils.WithField("tenant,uuid"), utils.WithError(err))
		return
	}

	if params.ParentId != nil && *params.ParentId != "" {
		parent, err := h.UserService.GetById(r.Context(), *params.ParentId)
		if err != nil {
			utils.Render(w, r, 422, utils.WithCode(string(ErrCodeN422000)), utils.WithField("parent_id"), utils.WithError(err))
			return
		}

		if parent.Tenant != params.Tenant {
			utils.Render(w, r, 422, utils.WithCode(string(ErrCodeN422999)), utils.WithField("parent_id"), utils.WithError(fmt.Errorf("parent tenant not equal params.tenant:%s", params.Tenant)))
			return
		}
	}

	res, err := h.UserService.Create(r.Context(), PostUsersJSONBody(params))
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}

	utils.Render(w, r, 201, utils.WithData(res))
}

// 删除绑定用户
// (DELETE /users/{uuid}/{tenant})
func (h *HttpServer) DeleteUsersUuidTenant(w http.ResponseWriter, r *http.Request, uuid string, tenant string) {
	_, err := h.UserService.GetByUuid(r.Context(), tenant, uuid)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(err))
		return
	}
	err = h.UserService.DeleteByUuid(r.Context(), tenant, uuid)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 204)
}

// 获取绑定用户详情(包含拥有的权限,角色,子级用户,角色组,菜单组等等)
// (GET /users/{uuid}/{tenant})
func (h *HttpServer) GetUsersUuidTenant(w http.ResponseWriter, r *http.Request, uuid string, tenant string) {
	res, err := h.UserService.GetByUuid(r.Context(), tenant, uuid)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(err))
		return
	}
	utils.Render(w, r, 200, utils.WithData(res))
}

// 更新绑定用户(角色,权限,子级,父级,分组等等)
// (PATCH /users/{uuid}/{tenant})
func (h *HttpServer) PatchUsersUuidTenant(w http.ResponseWriter, r *http.Request, uuid string, tenant string) {
	var params PatchUsersUuidTenantJSONRequestBody
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	userModel, err := h.UserService.GetByUuid(r.Context(), tenant, uuid)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(err))
		return
	}

	if params.ParentId != nil && *params.ParentId != "" {
		parent, err := h.UserService.GetById(r.Context(), *params.ParentId)
		if err != nil {
			utils.Render(w, r, 422, utils.WithCode(string(ErrCodeN422000)), utils.WithField("parent_id"), utils.WithError(err))
			return
		}

		if parent.Tenant != userModel.Tenant {
			utils.Render(w, r, 422, utils.WithCode(string(ErrCodeN422999)), utils.WithField("parent_id"), utils.WithError(fmt.Errorf("parent tenant not equal current.tenant:%s", userModel.Tenant)))
			return
		}
	}

	res, err := h.UserService.Update(r.Context(), PatchUsersUuidTenantJSONBody(params), tenant, uuid)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}

	utils.Render(w, r, 200, utils.WithData(res))
}

// 获取用户所有的关系图(节点,分组,权限)
// (GET /users/{uuid}/{tenant}/relations)
func (h *HttpServer) GetUsersUuidTenantRelations(w http.ResponseWriter, r *http.Request, uuid string, tenant string) {
	_, err := h.UserService.GetByUuid(r.Context(), tenant, uuid)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(err))
		return
	}
	res, err := h.UserService.GetAllRelations(r.Context(), tenant, uuid)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}

	utils.Render(w, r, 200, utils.WithData(res))
}

// 获取用户的所有路由(包含权限,角色,菜单,角色组等等...)
// (GET /users/{uuid}/{tenant}/routes)
func (h *HttpServer) GetUsersUuidTenantRoutes(w http.ResponseWriter, r *http.Request, uuid string, tenant string) {
	_, err := h.UserService.GetByUuid(r.Context(), tenant, uuid)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(err))
		return
	}
	res, err := h.UserService.GetAllRoutes(r.Context(), tenant, uuid)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}

	utils.Render(w, r, 200, utils.WithData(res))
}
