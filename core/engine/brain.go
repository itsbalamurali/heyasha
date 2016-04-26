package engine

import (
	"github.com/aichaos/rivescript-go"
	"fmt"
	"path/filepath"
)

//  Reply from Brain.
func BotReply(user_id string, usersay string) string {
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
	reply := base.Reply(user_id,usersay)
	return reply
}
