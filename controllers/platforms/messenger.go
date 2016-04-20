package platforms

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"encoding/json"
	"github.com/itsbalamurali/bot/core/platforms/messenger"
)

func MessengerBot(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	verify_token := "er7Wq4yREXBKpdRKjhAg"
	hub_mode := ps.ByName("hub.mode")
	hub_challenge := ps.ByName("hub.challenge")
	hub_verify_token := ps.ByName("hub.verify_token")
	if hub_verify_token == verify_token && hub_challenge != "" {
		w.Header().Set("Hub Mode",hub_mode)
		w.WriteHeader(200)
		fmt.Fprintln(w, hub_challenge)
		return
	}

	var msg = messenger.Receive{}
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		fmt.Printf("Something wrong: %s\n", err.Error())
		return
	}

	for _, entry := range msg.Entry {
		for _, info := range entry.Messaging {
			a := messenger.Classify(info, entry)
			if a == messenger.UnknownAction {
				fmt.Println("Unknown action:", info)
				continue
			}
			fmt.Printf("Recieved Text: %s\n", info.Message.Text)
			/*
			resp := &messenger.Response{
				to: messenger.Recipient{info.Sender.ID},
				token: "",
			}*/

		}
		/*
		sender := event.Sender.ID
		if event.Message != nil {
			fmt.Printf("Recieved Text: %s\n", event.Message.Text)
			err := messenger.SendMessage{}.SendTextMessage(sender, event.Message.Text)
			if err != nil {
				fmt.Printf("Something wrong: %s\n", err.Error())
			}
		}
		*/
	}
}

