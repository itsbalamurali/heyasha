package platforms

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/itsbalamurali/bot/core/platforms/telegram"
	"fmt"
	"encoding/json"
	"log"
)


func TelegramBot(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	update := &telegram.Update{}

	apiToken := "213239467:AAGWDAvFMfdfXuwlMkC2dSwKWEaW-NVl4bo"

	if err := json.NewDecoder(r.Body).Decode(update); err != nil {
		return  err
	}

	switch update.Type() {
	case telegram.MessageUpdate:
		msg := update.Message
		typ := msg.Type()
		if typ == telegram.TextMessage {
			fmt.Printf("<-%d, From:\t%s, Text: %s \n", msg.ID, msg.Chat, *msg.Text)

		}

		api, err := telegram.New(apiToken)

		if err != nil {
			log.Fatal(err)
		}

		// just to show its working
		fmt.Printf("User ID: %d\n", api.ID)
		fmt.Printf("Bot Name: %s\n", api.Name)
		fmt.Printf("Bot Username: %s\n", api.Username)

		// now simply echo that back
		outMsg, err := api.NewOutgoingMessage(telegram.NewRecipientFromChat(msg.Chat), *msg.Text).Send()

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

	//log.Printf("Json Received: %s\n", update)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", "Webhook Success!!")

}
