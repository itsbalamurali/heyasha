package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/itsbalamurali/heyasha/core/engine"
	"fmt"
)

func Chat(c *gin.Context) {
	//message := &models.Message{}
	chatmsg := c.Query("text")
	lang := c.Query("lang")
	rep, err := engine.BotReply("dummyid", chatmsg, lang)
	fmt.Printf(rep)
	if err != nil {
		c.Error(err)
		rep = "Something went wrong"
	}
	c.String(http.StatusOK, rep)
}