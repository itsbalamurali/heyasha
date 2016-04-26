package engine

import (
	"github.com/aichaos/rivescript-go"
	"fmt"
	"go/build"
)

//  Reply from Brain.
func BotReply(user_id string, usersay string) string {
	base := rivescript.New()
	base.UTF8 = true
	p, err := build.Default.Import("github.com/itsbalamurali/heyasha/data/brain", "", build.FindOnly)
	if err != nil {
		// handle error
		fmt.Printf("Unable to locate package")
	}
	loaderr := base.LoadDirectory(p.Dir)
	if loaderr != nil {
		fmt.Printf("Error loading from directory: %s", err)
	}
	base.SortReplies()
	reply := base.Reply(user_id,usersay)
	return reply
}
