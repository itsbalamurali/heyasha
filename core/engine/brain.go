package engine

import (
	//"github.com/jsgoecke/go-wit"
	//"os"
	"encoding/json"
	"log"
	"net/http"
	"github.com/itsbalamurali/heyasha/models"
	"github.com/itsbalamurali/heyasha/core/database"
	"fmt"
)

type BotResp struct {
	ConvoID string `json:"convo_id"`
	UserSay	string `json:"usersay"`
	BotSay	string `json:"botsay"`
}

// newMsg builds a message struct with Tokens, Stems, and a Structured Input.
func newMsg(u *models.User, msg string) *models.Message {
	tokens := TokenizeSentence(msg)
	stems := StemTokens(tokens)
	//si := ClassifyTokens(tokens)
	si := models.StructuredInput{}
	// Get the intents as determined by each plugin
	for domainID, c := range bClassifiers {
		scores, idx, _ := c.ProbScores(stems)
		log.Println("intent score", domainIntents[domainID][idx],
			scores[idx])
		if scores[idx] > 0.7 {
			si.Intents = append(si.Intents, string(domainIntents[domainID][idx]))
		}
	}

	m := &models.Message{
		User:            u,
		Sentence:        msg,
		Tokens:          tokens,
		Stems:           stems,
		//StructuredInput: si,
	}
	/*
		m, err = addContext(db, m)
		if err != nil {
			log.Debug(err)
		}
	*/
	return m
}

//BotReply  Reply from Brain.
func BotReply(user_id string, usersay string, lang ...string) (string, error) {
	db := database.MysqlCon()
	user := &models.User{}
	db.Where("email = ?","balamurali@live.com").First(&user)
	msg := newMsg(user,usersay)
	fmt.Printf("%+v",msg)
	//Call(msg.StructuredInput.Intents.)
	/*
	client := wit.NewClient("OBU6TR5J7EOJ7RR6HA7LER6W7NP5XRLX")
	// Process a text message
	request := &wit.MessageRequest{}
	request.Query = usersay
	result, err := client.Message(request)
	if err != nil {
		println(err)
		os.Exit(-1)
	}
	log.Println(result)
	data, _ := json.MarshalIndent(result, "", "    ")
	log.Println(string(data[:]))
	*/
	/*
	// Process an audio/wav message
	request = &wit.MessageRequest{}
	request.File = "../audio_sample/helloWorld.wav"
	request.ContentType = "audio/wav;rate=8000"
	result, err = client.AudioMessage(request)
	if err != nil {
		println(err)
		os.Exit(-1)
	}
	log.Println(result)
	data, _ = json.MarshalIndent(result, "", "    ")
	log.Println(string(data[:]))
        */

	r := BotResp{}
	url := "https://asha-ai-api.heyasha.com/chatbot/conversation_start.php"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}
	req.URL.RawQuery = "convo_id="+ user_id +"&say=" + usersay
	api := &http.Client{}
	resp, err := api.Do(req)
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&r)
	return r.BotSay, err
}

func EntityParser()  {
	
}