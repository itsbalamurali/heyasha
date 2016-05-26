package engine

import (
	"fmt"
	"github.com/aichaos/rivescript-go"
	"path/filepath"
)

//  Reply from Brain.
func BotReply(user_id string, usersay string) string {

	// Process a text message
	request := &wit.MessageRequest{}
	request.Query = "Hello world"
	result, err := client.Message(request)
	if err != nil {
		println(err)
		os.Exit(-1)
	}
	log.Println(result)
	data, _ := json.MarshalIndent(result, "", "    ")
	log.Println(string(data[:]))

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

	base := rivescript.New()
	base.UTF8 = true
	path, notfound_err := filepath.Abs("./data/brain")
	if notfound_err != nil {
		fmt.Println(notfound_err)
	}
	err := base.LoadDirectory(path)
	if err != nil {
		fmt.Printf("Error loading from directory: %s", err)
	}
	base.SortReplies()
	reply := base.Reply(user_id, usersay)
	return reply

}
