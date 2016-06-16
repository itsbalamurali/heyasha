package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/itsbalamurali/heyasha/core/engine"
	"fmt"
	"time"
)

func Chat(c *gin.Context) {
	chatmsg := c.Query("text")
	lang := c.Query("lang")
	rep, err := engine.BotReply("dummyid", chatmsg, lang)
	fmt.Printf(rep)
	if err != nil {
		c.Error(err)
		rep = "Something went wrong"
	}
	if rep == ""{
		rep = "Something went wrong"

	}
	c.JSON(http.StatusOK, gin.H{"success":true,"timestamp":time.Now().UnixNano(),"userSay":chatmsg,"ashaSay":rep,"body":nil})
}