package platforms

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SmsBot(c *gin.Context) {
	c.String(http.StatusOK, "Sms Bot Response!\n")
}
