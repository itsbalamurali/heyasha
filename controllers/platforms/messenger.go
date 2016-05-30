package platforms

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/iron-io/iron_go3/mq"
	"github.com/itsbalamurali/heyasha/core/platforms/messenger"
	"net/http"
	"github.com/itsbalamurali/heyasha/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/itsbalamurali/heyasha/core/engine"
	"github.com/jinzhu/gorm"
	"strconv"
)

const (
	token = "EAAGeBVsm2kQBALYaKjHZBVlMhf4nFx5LLztRiHMnpUjvb4gHAIzxqM6srWraxu2VtPWZAPEOtZCbZCha5MEiOQF5wcXojnQYgrTPTuoxV5YQZCAQ5qbx9mlfrKxv2TcG0e4m9xgAGbELW9uEoNChAsRFZCo0UOSbujn9OZArQNGXgZDZD"
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
	var queue = mq.New("messages");

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
			//fmt.Println("Message: " +info.Message.Text)
			//fmt.Println("Sender: " + info.Sender.ID)
			//ai_msg := engine.BotReply(strconv.FormatInt(info.Message.Sender.ID, 10), info.Message.Text)
			profile, fberr := messenger.ProfileByID(info.Sender.ID,token)
			if fberr != nil {
				log.Errorln(err.Error())
			}

			user := &models.User{
				Pid: "fb"+info.Sender.ID,
				FirstName:profile.FirstName,
				LastName:profile.LastName,
				ProfilePicURL:profile.ProfilePicURL,
				Platforms:[]models.Platform{
					{
					PlatformID:info.Sender.ID,
					Name:"facebook",
					},
				},
			}

			db := c.MustGet("db").(*mgo.Database)

			count, err := db.C("users").Find(bson.M{"pid": "fb"+info.Sender.ID}).Limit(1).Count()
			if err != nil {
				c.Error(err)
			}
			if count == 0 {
				//Document doesnt exist
				//Insert Document
				err = db.C("users").Insert(user)
				if err != nil {
					c.Error(err)
				}
			}
			_, qerr := queue.PushString(user.Pid+":----:"+info.Message.Text)
			if qerr != nil {
				c.Error(qerr)
			}

			rep, err := engine.BotReply(user.Pid,info.Message.Text)
			if err != nil || rep == "" {
				rep = "Whoops my brains not working!!!!"
				log.Println(err)
			}
			resp.Text(rep)

			mysqldb := c.MustGet("mysql").(*gorm.DB)
			convlog := &models.ConversationLog{
				Input:info.Message.Text,
				Response:rep,
				UserID: strconv.Atoi(user.Pid),
				ConvoID:user.Pid,
			}
			mysqldb.Create(&convlog)
		}
	}
	c.String(http.StatusOK, "Webhook Success!!!")
}