package ports

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/keepondream/RBAC_service/internal/common/utils"
)

type Noder interface {
	Create(ctx context.Context, params PostNodesJSONBody) (*NodeInfoResponse, error)
	IsUnique(ctx context.Context, tenant, name, node_type string) error
	GetById(ctx context.Context, id string) (*NodeInfoResponse, error)
	DeleteById(ctx context.Context, id string) error
	List(ctx context.Context, params GetNodesParams) (*NodeListResponse, error)
	Update(ctx context.Context, params PatchNodesIdJSONBody, id string) (*NodeInfoResponse, error)
}

// 节点列表
// (GET /nodes)
func (h *HttpServer) GetNodes(w http.ResponseWriter, r *http.Request, params GetNodesParams) {
	res, err := h.NodeService.List(r.Context(), params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 200, utils.WithData(res))
}

// 创建节点
// (POST /nodes)
func (h *HttpServer) PostNodes(w http.ResponseWriter, r *http.Request) {
	var params PostNodesJSONRequestBody
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}

	err = h.NodeService.IsUnique(r.Context(), params.Tenant, params.Name, params.Type)
	if err != nil {
		utils.Render(w, r, 422, utils.WithCode(string(ErrCodeN422111)), utils.WithField("tenant,name,type"), utils.WithError(err))
		return
	}

	if params.ParentId != nil && *params.ParentId != "" {
		parent, err := h.NodeService.GetById(r.Context(), *params.ParentId)
		if err != nil {
			utils.Render(w, r, 422, utils.WithCode(string(ErrCodeN422000)), utils.WithField("parent_id"), utils.WithError(err))
			return
		}

		if parent.Tenant != params.Tenant {
			utils.Render(w, r, 422, utils.WithCode(string(ErrCodeN422999)), utils.WithField("parent_id"), utils.WithError(fmt.Errorf("parent tenant not equal params.tenant:%s", params.Tenant)))
			return
		}
	}

	res, err := h.NodeService.Create(r.Context(), PostNodesJSONBody(params))
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}

	utils.Render(w, r, 200, utils.WithData(res))
}

// 删除节点
// (DELETE /nodes/{id})
func (h *HttpServer) DeleteNodesId(w http.ResponseWriter, r *http.Request, id string) {
	_, err := h.NodeService.GetById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(err))
		return
	}
	err = h.NodeService.DeleteById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 204)
}

// 获取节点详情
// (GET /nodes/{id})
func (h *HttpServer) GetNodesId(w http.ResponseWriter, r *http.Request, id string) {
	res, err := h.NodeService.GetById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(err))
		return
	}
	utils.Render(w, r, 200, utils.WithData(res))
}

// 更新节点信息
// (PATCH /nodes/{id})
func (h *HttpServer) PatchNodesId(w http.ResponseWriter, r *http.Request, id string) {
	var params PatchNodesIdJSONRequestBody
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	_, err = h.NodeService.GetById(r.Context(), id)
	if err != nil {
		utils.Render(w, r, 404, utils.WithError(err))
		return
	}

	if params.ParentId != nil && *params.ParentId != "" {
		parent, err := h.NodeService.GetById(r.Context(), *params.ParentId)
		if err != nil {
			utils.Render(w, r, 422, utils.WithCode(string(ErrCodeN422000)), utils.WithField("parent_id"), utils.WithError(err))
			return
		}

		if parent.Tenant != params.Tenant.Tenant {
			utils.Render(w, r, 422, utils.WithCode(string(ErrCodeN422999)), utils.WithField("parent_id"), utils.WithError(fmt.Errorf("parent tenant not equal params.tenant:%s", params.Tenant)))
			return
		}
	}

	res, err := h.NodeService.Update(r.Context(), PatchNodesIdJSONBody(params), id)
	if err != nil {
		utils.Render(w, r, 400, utils.WithError(err))
		return
	}
	utils.Render(w, r, 200, utils.WithData(res))
}
