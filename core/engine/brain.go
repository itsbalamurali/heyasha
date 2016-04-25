package engine

import (
	"github.com/aichaos/rivescript-go"
	"fmt"
)

type BotBrain struct {
	bot     *rivescript.RiveScript
}

// New returns a new Brain.
func Brain() *BotBrain {
	base := rivescript.New()
	base.UTF8 = true
	return &BotBrain{
		bot:  base,
	}
}

func (brain *BotBrain) Reply(user_id string, usersay string) string {

	// Load a directory full of RiveScript documents (.rive files)
	err := brain.bot.LoadDirectory("../../data/brain")
	if err != nil {
		fmt.Printf("Error loading from directory: %s", err)
	}
	brain.bot.SortReplies()
	reply := brain.bot.Reply(user_id,usersay)

	return reply
}
