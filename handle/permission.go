package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/keepondream/RBAC_service/utils"
)

func (h *Handle) ListPermissions(c *gin.Context) {

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
		utils.Failed(c, utils.WithErr(utils.Err_HTTP_Json.WithErr(err)))
		return
	}

	_, err := h.Service.AddPermission(c.Request.Context(), AllPermission, req.Tenant, req.Uri, req.Method, req.Name, req.Comment)

	if err != nil {
		utils.Failed(c, utils.WithErr(utils.Err_Failed.WithErr(err)))
		return
	}

	utils.Success(c)
}

type EidtPermissionsRequest struct {
	Tenant  string `json:"tenant" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Uri     string `json:"uri" binding:"required"`
	Method  string `json:"method" binding:"required"`
	Comment string `json:"comment"`
}

func (h *Handle) EditPermission(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		utils.Failed(c, utils.WithErr(utils.Err_HTTP_Param))
		return
	}

	var req EidtPermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Failed(c, utils.WithErr(utils.Err_HTTP_Json.WithErr(err)))
		return
	}

	err := h.Service.EditPermission(c.Request.Context(), id, req.Tenant, req.Uri, req.Method, req.Name, req.Comment)

	if err != nil {
		utils.Failed(c, utils.WithErr(utils.Err_Failed.WithErr(err)))
		return
	}

	utils.Success(c)
}

func (h *Handle) DeletePermission(c *gin.Context) {
	id := c.Param("id")
	tenant := c.Param("tenant")
	if id == "" || tenant == "" {
		utils.Failed(c, utils.WithErr(utils.Err_HTTP_Param))
		return
	}

	err := h.Service.DeletePermission(c.Request.Context(), id, tenant)

	if err != nil {
		utils.Failed(c, utils.WithErr(utils.Err_Failed.WithErr(err)))
		return
	}

	utils.Success(c)
}
