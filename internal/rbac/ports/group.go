package ports

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/keepondream/RBAC_service/internal/common/utils"
)

type Grouper interface {
	Create(ctx context.Context, params PostGroupsJSONBody) (*GroupInfoResponse, error)
	IsUnique(ctx context.Context, tenant, name, group_type string) error
	GetById(ctx context.Context, id string) (*GroupInfoResponse, error)
	DeleteById(ctx context.Context, id string) error
	List(ctx context.Context, params GetGroupsParams) (*GroupListResponse, error)
	Update(ctx context.Context, params PatchGroupsIdJSONBody, id string) (*GroupInfoResponse, error)
}

// 分组列表
// (GET /groups)
func (h *HttpServer) GetGroups(w http.ResponseWriter, r *http.Request, params GetGroupsParams) {
	res, err := h.GroupService.List(r.Context(), params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 200, utils.WithData(res))
}

// 创建分组
// (POST /groups)
func (h *HttpServer) PostGroups(w http.ResponseWriter, r *http.Request) {
	var params PostGroupsJSONRequestBody
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}

	err = h.GroupService.IsUnique(r.Context(), string(params.Tenant), params.Name, string(params.Type))
	if err != nil {
		utils.Render(w, r, 422, utils.WithCode(string(ErrCodeN422111)), utils.WithField("tenant,name,type"), utils.WithError(err))
		return
	}

	res, err := h.GroupService.Create(r.Context(), PostGroupsJSONBody(params))
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}

	utils.Render(w, r, 201, utils.WithData(res))
}

// 删除分组
// (DELETE /groups/{id})
func (h *HttpServer) DeleteGroupsId(w http.ResponseWriter, r *http.Request, id string) {
	_, err := h.GroupService.GetById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(err))
		return
	}
	err = h.GroupService.DeleteById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 204)
}

// 获取分组详情
// (GET /groups/{id})
func (h *HttpServer) GetGroupsId(w http.ResponseWriter, r *http.Request, id string) {
	res, err := h.GroupService.GetById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(err))
		return
	}
	utils.Render(w, r, 200, utils.WithData(res))
}

// 更新分组信息
// (PATCH /groups/{id})
func (h *HttpServer) PatchGroupsId(w http.ResponseWriter, r *http.Request, id string) {
	var params PatchGroupsIdJSONRequestBody
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	_, err = h.GroupService.GetById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(err))
		return
	}
	res, err := h.GroupService.Update(r.Context(), PatchGroupsIdJSONBody(params), id)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 200, utils.WithData(res))
}
