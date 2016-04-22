package engine

import (
	"github.com/dghubble/sling"
	"log"
	"net/http"
)

type APIClient struct {
	sling     *sling.Sling
}

type ChatParams struct {
	UserSay string `url:"say,omitempty"`
	ConvoId string `url:"convo_id,omitempty"`
}

type ChatResponse struct  {
	ConvoId  string `json:"convo_id"`
	UserSay   string `json:"usersay"`
	BotSay   string `json:"botsay"`
}

// NewClient returns a new Client.
func NewClient() *APIClient {
	base := sling.New().Base("https://asha-ai-api.heyasha.com/chatbot/conversation_start.php")
	return &APIClient{
		sling:  base,
	}
}

func (ApiC *APIClient) ChatConvoAI(convo_id string, usersay string) string {

	botresponse := &ChatResponse{}
	params := &ChatParams{ConvoId: convo_id, UserSay: usersay}
	//fmt.Println("Convo id: " + convo_id)
	//fmt.Println("User Say: " + usersay)
	res, err := ApiC.sling.New().QueryStruct(params).ReceiveSuccess(botresponse)
	if err != nil{
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal("AI API Status code error")
	}
	return botresponse.BotSay
}
