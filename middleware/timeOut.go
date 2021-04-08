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

		c.Next()
	}
}
