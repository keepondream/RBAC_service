package utils

import (
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

// GinRender gin框架统一输出格式
func GinRender(c *gin.Context, code int, modResp ...ModResponse) {
	resp := Response{}
	resp.Msg = http.StatusText(code)

	for _, fn := range modResp {
		fn(&resp)
	}

	if resp.Data != nil {
		c.AbortWithStatusJSON(code, resp)
		return
	}

	if resp.Err != nil {
		log.Printf("render response err: %v \n", resp.Err)
		resp.Msg = resp.Err.Error()
		if e, ok := resp.Err.(net.Error); ok && e.Timeout() {
			code = http.StatusGatewayTimeout
		} else if resp.Err == sql.ErrNoRows {
			code = http.StatusNotFound
		} else if e, ok := resp.Err.(pg.Error); ok && e.IntegrityViolation() {
			code = http.StatusBadRequest
		}
	}

	c.AbortWithStatusJSON(code, resp)
}
