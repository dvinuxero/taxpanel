package status

import "net/http"
import "github.com/gin-gonic/gin"

//Health check
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

//Legacy health check
func Ok(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ok": "i'm alive"})
}
