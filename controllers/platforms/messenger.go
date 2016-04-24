package platforms

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"encoding/json"
	"github.com/itsbalamurali/heyasha/core/platforms/messenger"
	"strconv"
	"github.com/itsbalamurali/heyasha/core/engine"
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
				"CAAQCVhc9mrYBAGYZCX9c0EQlIwecNPE5Th4b8T1zVvIumZA3verOV1RZC5ZBvQIqJJRZAHD4TYCHVeZBVOL1lpwn9qxEGLih0qUCR3ZC57kEY8O5BlflKNEmKyxDKpPpgambCWgWLizhkmsfob1d5OHDHGUILeaZBqbuyuk5ix2J9FJ7GjwCE0zwxLF76KidSQMDA5OMz7vlLgZDZD",
				messenger.Recipient{info.Sender.ID},
			}

			ai_msg := engine.NewClient().ChatConvoAI(strconv.FormatInt(info.Message.Sender.ID,10), info.Message.Text)
			resp.Text(ai_msg)
		}
	}
}

