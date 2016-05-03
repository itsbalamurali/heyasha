package platforms

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func KikBot(c *gin.Context) {
	c.String(http.StatusOK, "Kik Bot Response!\n")
}
