package platforms

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/itsbalamurali/heyasha/core/platforms/messenger"
	"net/http"
	"github.com/itsbalamurali/heyasha/models"
	"github.com/itsbalamurali/heyasha/core/engine"
	"github.com/jinzhu/gorm"
)

const (
	//old//token = "EAAGeBVsm2kQBALYaKjHZBVlMhf4nFx5LLztRiHMnpUjvb4gHAIzxqM6srWraxu2VtPWZAPEOtZCbZCha5MEiOQF5wcXojnQYgrTPTuoxV5YQZCAQ5qbx9mlfrKxv2TcG0e4m9xgAGbELW9uEoNChAsRFZCo0UOSbujn9OZArQNGXgZDZD"
	token = "EAAGeBVsm2kQBACeMa1y7xuwd9nD4oNO66oXa1jwOIhyeK4rnVuxochJ1aJGNvCw4WIDCUl4SlymmFkuIfqXf7423hHixjbZBvUh4VGswpfvyrQ96mfHVIUbCjc6TWsrgbg4bgdvC8YAOOGXhaYeOtDKdj8ZAjOltFQye4diwZDZD"
)

func MessengerBotVerify(c *gin.Context) {
	verify_token := "er7Wq4yREXBKpdRKjhAg"
	hub_mode := c.Query("hub.mode")
	hub_challenge := c.Query("hub.challenge")
	hub_verify_token := c.Query("hub.verify_token")
	if hub_verify_token == verify_token && hub_challenge != "" {
		c.Header("Hub Mode", hub_mode)
		c.String(http.StatusOK, hub_challenge)
	} else {
		c.String(http.StatusBadRequest,"Bad Request")
	}
}

func MessengerBotChat(c *gin.Context) {
	var msg = messenger.Receive{}

	err := c.BindJSON(&msg)
	if err != nil {
		c.Error(err)
		return
	}

	for _, entry := range msg.Entry {
		for _, info := range entry.Messaging {
			a := messenger.Classify(info, entry)
			if a == messenger.UnknownAction {
				fmt.Println("Unknown action:", info)
				continue
			}
			resp := &messenger.Response{
				token,
				messenger.Recipient{info.Sender.ID},
			}


			rep, err := engine.BotReply(info.Sender.ID,info.Message.Text)
			if err != nil || rep == "" {
				rep = "Whoops my brains not working!!!!"
				log.Println(err)
			}
			resp.Text(rep)
			mysqldb := c.MustGet("mysql").(*gorm.DB)
			convlog := &models.ConversationLog{
				Input:info.Message.Text,
				Response:rep,
				UserID: info.Sender.ID,
				ConvoID: info.Sender.ID,
			}
			mysqldb.Create(&convlog)
		}
	}
	c.String(http.StatusOK, "Webhook Success!!!")
}