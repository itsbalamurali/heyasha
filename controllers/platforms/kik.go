package platforms

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/itsbalamurali/heyasha/models"
	"log"
	"github.com/itsbalamurali/heyasha/core/engine"
	"github.com/itsbalamurali/heyasha/core/platforms/kik"
)


func KikBot(c *gin.Context) {
	var msg = kik.KikWebhook{}

	err := c.BindJSON(&msg)
	if err != nil {
		c.Error(err)
		return
	}
	for _, message := range msg.Messages {
		a := kik.Classify(message)
		if a == kik.UnknownType {
			fmt.Println("Unknown action:", message.Type)
			continue
		}
		if a == kik.TextType {

		rep, err := engine.BotReply(message.From, message.Body)
		if err != nil || rep == "" {
			rep = "Whoops my brains not working!!!!"
			log.Println(err)
		}
		kik.Text(message.From,message.ChatID,rep)
		mysqldb := c.MustGet("mysql").(*gorm.DB)
		convlog := &models.ConversationLog{
			Input:message.Body,
			Response:rep,
			UserID: message.From,
			ConvoID: message.ChatID,
		}
		mysqldb.Create(&convlog)
		}
	}
	c.String(http.StatusOK, "Kik Bot Response!\n")
}