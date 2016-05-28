package platforms

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/itsbalamurali/heyasha/core/platforms/telegram"
	"log"
	"net/http"
	"github.com/itsbalamurali/heyasha/core/engine"
	"github.com/iron-io/iron_go3/mq"
	"strconv"
)

func TelegramBot(c *gin.Context) {

	update := &telegram.Update{}
	var queue = mq.New("messages");

	apiToken := "213239467:AAGWDAvFMfdfXuwlMkC2dSwKWEaW-NVl4bo"
	//Decode incoming json
	c.BindJSON(&update)
	switch update.Type() {
	case telegram.MessageUpdate:
		msg := update.Message
		/*typ := msg.Type()
		  if typ == telegram.TextMessage {
				//fmt.Printf("<-%d, From:\t%s, Text: %s \n", msg.ID, msg.Chat, *msg.Text)
		  }*/
		api, err := telegram.New(apiToken)
		if err != nil {
			log.Fatal(err)
		}

		_, qerr := queue.PushString(strconv.Itoa(msg.Chat.ID)+":----:"+*msg.Text)
		if qerr != nil {
			c.Error(qerr)
		}

		rep, err := engine.BotReply(strconv.Itoa(msg.Chat.ID),*msg.Text)
		if err != nil {
			rep = "Whoops my brains not working!!!!"
			log.Println(err)
		}

		outMsg, err := api.NewOutgoingMessage(telegram.NewRecipientFromChat(msg.Chat), rep).Send()

		if err != nil {
			fmt.Printf("Error sending: %s\n", err)
		}

		fmt.Printf("->%d, To:\t%s, Text: %s\n", outMsg.Message.ID, outMsg.Message.Chat, *outMsg.Message.Text)
	case telegram.InlineQueryUpdate:
		fmt.Println("Ignoring received inline query: ", update.InlineQuery.Query)
	case telegram.ChosenInlineResultUpdate:
		fmt.Println("Ignoring chosen inline query result (ID): ", update.ChosenInlineResult.ID)
	default:
		fmt.Printf("Ignoring unknown Update type.")
	}
	//Webhook received and is success
	c.String(http.StatusOK, "Webhook Success!!")
}
