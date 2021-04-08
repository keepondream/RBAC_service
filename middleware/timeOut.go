package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/keepondream/RBAC_service/utils"
	"github.com/spf13/cast"
)

// TimeOut 超时中间件
func TimeOut(config *utils.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), cast.ToDuration(config.TIME_OUT)*time.Millisecond)

		defer func() {
			cancel()
		}()

		c.Request = c.Request.WithContext(ctx)

		// 上下文传递是否启用debug, 控制接口输出 developMsg
		if config.LOG_LEVEL == "debug" {
			c.Set("debug", true)
		} else {
			c.Set("debug", false)
		}

		c.Next()
	}
}
