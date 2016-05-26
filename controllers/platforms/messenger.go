package platforms

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/itsbalamurali/heyasha/core/platforms/messenger"
	"net/http"
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
		log.Errorln("Something wrong: %s", err.Error())
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
				"EAAGeBVsm2kQBALYaKjHZBVlMhf4nFx5LLztRiHMnpUjvb4gHAIzxqM6srWraxu2VtPWZAPEOtZCbZCha5MEiOQF5wcXojnQYgrTPTuoxV5YQZCAQ5qbx9mlfrKxv2TcG0e4m9xgAGbELW9uEoNChAsRFZCo0UOSbujn9OZArQNGXgZDZD",
				messenger.Recipient{info.Sender.ID},
			}
			//fmt.Println("Message: " +info.Message.Text)
			//fmt.Println("Sender: " + info.Sender.ID)
			//ai_msg := engine.BotReply(strconv.FormatInt(info.Message.Sender.ID, 10), info.Message.Text)
			resp.Text(info.Message.Text)
		}
	}
	c.String(http.StatusOK, "Webhook Success!!!")
}