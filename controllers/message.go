package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/itsbalamurali/heyasha/models"
	"net/http"
	"github.com/itsbalamurali/heyasha/core/engine"
)

func Chat(c *gin.Context) {
	message := &models.Message{}
	chatmsg := c.Query("text")
	lang := c.Query("lang")
	engine.BotReply("dummyid", chatmsg, lang)
	c.JSON(http.StatusOK,&message)
}
