package ports

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/keepondream/RBAC_service/internal/common/utils"
)

type Permissioner interface {
	Create(ctx context.Context, params PostPermissionsJSONBody) (*PermissionInfoResponse, error)
	IsUnique(ctx context.Context, tenant, name string) error
	GetById(ctx context.Context, id string) (*PermissionInfoResponse, error)
	DeleteById(ctx context.Context, id string) error
	List(ctx context.Context, params GetPermissionsParams) (*PermissionListResponse, error)
	Update(ctx context.Context, params PatchPermissionsIdJSONBody, id string) (*PermissionInfoResponse, error)
}

// 权限列表
// (GET /permissions)
func (h *HttpServer) GetPermissions(w http.ResponseWriter, r *http.Request, params GetPermissionsParams) {
	res, err := h.PermissionService.List(r.Context(), params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 200, utils.WithData(res))
}

// 创建权限
// (POST /permissions)
func (h *HttpServer) PostPermissions(w http.ResponseWriter, r *http.Request) {
	var params PostPermissionsJSONRequestBody
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}

	err = h.PermissionService.IsUnique(r.Context(), params.Tenant, params.Name)
	if err != nil {
		utils.Render(w, r, 422, utils.WithCode(string(ErrCodeN422111)), utils.WithField("tenant,name"), utils.WithError(err))
		return
	}

	res, err := h.PermissionService.Create(r.Context(), PostPermissionsJSONBody(params))
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 201, utils.WithData(res))
}

// 删除权限
// (DELETE /permissions/{id})
func (h *HttpServer) DeletePermissionsId(w http.ResponseWriter, r *http.Request, id string) {
	_, err := h.PermissionService.GetById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(fmt.Errorf("not found")))
		return
	}
	err = h.PermissionService.DeleteById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 204)
}

// 获取权限详情
// (GET /permissions/{id})
func (h *HttpServer) GetPermissionsId(w http.ResponseWriter, r *http.Request, id string) {
	res, err := h.PermissionService.GetById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(fmt.Errorf("not found")))
		return
	}
	utils.Render(w, r, 200, utils.WithData(res))
}

// 更新权限信息
// (PATCH /permissions/{id})
func (h *HttpServer) PatchPermissionsId(w http.ResponseWriter, r *http.Request, id string) {
	var params PatchPermissionsIdJSONRequestBody
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	_, err = h.PermissionService.GetById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(fmt.Errorf("not found")))
		return
	}

	res, err := h.PermissionService.Update(r.Context(), PatchPermissionsIdJSONBody(params), id)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}

	utils.Render(w, r, 200, utils.WithData(res))
}
