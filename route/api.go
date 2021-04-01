package route

import (
	"github.com/gin-gonic/gin"
	"github.com/keepondream/RBAC_service/handle"
	"github.com/keepondream/RBAC_service/middleware"
	"github.com/keepondream/RBAC_service/utils"
)

func NewRoute(config *utils.Config, handle *handle.Handle) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CorsMiddleware())

	r := router.Group("/v1/api")

	r.GET("permissions", handle.ListPermissions)

	return router
}
