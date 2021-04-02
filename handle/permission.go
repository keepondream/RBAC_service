package handle

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/keepondream/RBAC_service/utils"
)

func (h *Handle) ListPermissions(c *gin.Context) {

}

type CreatePermissionsRequest struct {
	Name    string `json:"name" binding:"required"`
	Uri     string `json:"uri" binding:"required"`
	Method  string `json:"method" binding:"required"`
	Comment string `json:"comment" `
}

func (h *Handle) CreatePermissions(c *gin.Context) {
	var req CreatePermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.Failed(c, utils.WithErr(utils.Err_HTTP_Json.WithErr(err)))
		return
	}

	_ = req

	fmt.Println(req)
}

func (h *Handle) EditPermissions(c *gin.Context) {
	fmt.Println(h.Config)
}

func (h *Handle) DeletePermissions(c *gin.Context) {
	fmt.Println(h.Config)
}
