package ports

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/keepondream/RBAC_service/internal/common/utils"
)

type Router interface {
	GetById(ctx context.Context, id string) (*RouteInfoResponse, error)
	List(ctx context.Context, params GetRoutesParams) (*RouteListResponse, error)
	Create(ctx context.Context, params PostRoutesJSONBody) (*RouteInfoResponse, error)
	DeleteById(ctx context.Context, id string) error
	Update(ctx context.Context, params PatchRoutesIdJSONBody, id string) (*RouteInfoResponse, error)
	IsUnique(ctx context.Context, tenant, uri, method string) error
}

// 路由列表
// (GET /routes)
func (h *HttpServer) GetRoutes(w http.ResponseWriter, r *http.Request, params GetRoutesParams) {
	res, err := h.RouteService.List(r.Context(), params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 200, utils.WithData(res))
}

// 创建路由
// (POST /routes)
func (h *HttpServer) PostRoutes(w http.ResponseWriter, r *http.Request) {
	var params PostRoutesJSONRequestBody
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}

	err = h.RouteService.IsUnique(r.Context(), string(params.Tenant), params.Uri, string(params.Method))
	if err != nil {
		utils.Render(w, r, 422, utils.WithCode(string(ErrCodeN422111)), utils.WithField("tenant,uri,method"), utils.WithError(err))
		return
	}
	res, err := h.RouteService.Create(r.Context(), PostRoutesJSONBody(params))
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 201, utils.WithData(res))
}

// 删除路由
// (DELETE /routes/{id})
func (h *HttpServer) DeleteRoutesId(w http.ResponseWriter, r *http.Request, id string) {
	_, err := h.RouteService.GetById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(fmt.Errorf("not found")))
		return
	}
	err = h.RouteService.DeleteById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 204)
}

// 路由详情
// (GET /routes/{id})
func (h *HttpServer) GetRoutesId(w http.ResponseWriter, r *http.Request, id string) {
	res, err := h.RouteService.GetById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(fmt.Errorf("not found")))
		return
	}
	utils.Render(w, r, 200, utils.WithData(res))
}

// 修改路由信息
// (PATCH /routes/{id})
func (h *HttpServer) PatchRoutesId(w http.ResponseWriter, r *http.Request, id string) {
	var params PatchRoutesIdJSONRequestBody
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	_, err = h.RouteService.GetById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(fmt.Errorf("not found")))
		return
	}
	res, err := h.RouteService.Update(r.Context(), PatchRoutesIdJSONBody(params), id)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 200, utils.WithData(res))
}
