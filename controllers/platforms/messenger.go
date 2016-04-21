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

			resp := &messenger.Response{
				"CAAQCVhc9mrYBAFxgONRKzps2FdCYaUrtIPRVtZBEi9yN8skZB5RyySwDWgMhyRtFZCyVcOVGMTQXUQ4uPEvkNy0ZBLjpyEZA6xhTzP7Np9N4UCDnBIGG9XKMVEXkiZAZBhICIem0m5P06iv1k24Fpbpxnh3fONAO8DAbKjRrRn2V9SvXS2iyCTj3IFxPkfoUu5vZC4c0aa1zOgZDZD",
				messenger.Recipient{info.Sender.ID},
			}

			resp.Text(info.Message.Text)

			fmt.Printf("Recieved Text: %s\n", info.Message.Text)


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

