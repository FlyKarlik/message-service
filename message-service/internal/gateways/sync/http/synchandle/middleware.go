package synchandle

import "github.com/gin-gonic/gin"

func (h *SyncHandler) SetJSON(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Next()
}
