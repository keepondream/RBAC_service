package handle

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handle) ListPermissions(c *gin.Context) {
	fmt.Println(h.Config)
}
