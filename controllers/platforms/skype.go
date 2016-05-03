package platforms

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SkypeBot(c *gin.Context) {
	c.String(http.StatusOK, "Skype Bot Response!\n")
}
