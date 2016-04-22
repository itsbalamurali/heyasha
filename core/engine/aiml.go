package engine

import (
	"github.com/dghubble/sling"
	"log"
	"encoding/json"
)

type ChatParams struct {
	ConvoId string `url:"convo_id,omitempty"`
	UserSay string `url:"usersay,omitempty"`
}

type ChatResponse struct  {
	ConvoId  string `json:"convo_id"`
	UserSay   string `json:"usersay"`
	BotSay   string `json:"botsay"`

}

func ChatConvoAI(convo_id string, usersay string)  string {

	botresponse := &ChatResponse{}
	params := &ChatParams{ConvoId: convo_id, UserSay: usersay}
	req, err := sling.New().Get("https://asha-ai-api.heyasha.com/chatbot/conversation_start.php").QueryStruct(params).Request()
	if err != nil{
		log.Fatal(err)
	}

	json.NewDecoder(req.Body).Decode(botresponse)

	return botresponse.BotSay
}
