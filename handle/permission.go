package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/keepondream/RBAC_service/service"
	"github.com/keepondream/RBAC_service/utils"
	"github.com/spf13/cast"
)

func (h *Handle) ListPermissions(c *gin.Context) {
	tenant := c.Param("tenant")
	page := cast.ToInt32(c.DefaultQuery("page", "1"))
	pageSize := cast.ToInt32(c.DefaultQuery("page_size", "10"))

	if page <= 0 || pageSize <= 0 {
		utils.Failed(c, utils.WithErr(utils.Err_HTTP_Query))
		return
	}

	resp, err := h.Service.ListPermission(c.Request.Context(), page, pageSize, tenant)
	if err != nil {
		utils.Failed(c, utils.WithErr(err))
		return
	}

	utils.Success(c, utils.WithData(resp))

}

type CreatePermissionsRequest struct {
	Tenant  string `json:"tenant" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Uri     string `json:"uri" binding:"required"`
	Method  string `json:"method" binding:"required"`
	Comment string `json:"comment"`
}

func (h *Handle) CreatePermission(c *gin.Context) {
	var req CreatePermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Failed(c, utils.WithErr(err))
		return
	}

	_, err := h.Service.AddPermission(c.Request.Context(), service.Permission{
		Sign:    service.AllPermission,
		Tenant:  req.Tenant,
		Uri:     req.Uri,
		Method:  req.Method,
		Name:    req.Name,
		Comment: req.Comment,
	})

	if err != nil {
		utils.Failed(c, utils.WithErr(err))
		return
	}

	utils.Success(c)
}

func (h *Handle) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	tenant := c.Param("tenant")

	err := h.Service.DeletePermission(c.Request.Context(), id, tenant)

	if err != nil {
		utils.Failed(c, utils.WithErr(err))
		return
	}

	utils.Success(c)
}
