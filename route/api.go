package route

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/keepondream/RBAC_service/handle"
	"github.com/keepondream/RBAC_service/middleware"
	"github.com/keepondream/RBAC_service/utils"
)

func NewRoute(config *utils.Config, handle *handle.Handle) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CorsMiddleware(), middleware.TimeOut(config))

	r := router.Group("/v1/api")

	utils.AddRoute(r, "/:tenant/permissions", "GET", "权限列表", handle.ListPermissions)
	utils.AddRoute(r, "/permissions", "POST", "添加权限", handle.CreatePermission)
	utils.AddRoute(r, "/:tenant/permission/:id", "DELETE", "删除权限", handle.DeletePermission)

	fmt.Println()
	fmt.Println("route数量", utils.RouteNum)
	fmt.Println()
	res, _ := json.Marshal(utils.AllRoutes)
	fmt.Println("all route", string(res))
	fmt.Println()
	fmt.Println()

	return router
}
