package platforms

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SlackBot(c *gin.Context) {
	c.String(http.StatusOK, "Slack Bot Response!\n")
}
