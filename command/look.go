package command

import (
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/state"
)

type lookCommand struct {
	baseCommand
}

func NewLookCommand() Command {
	return &lookCommand{baseCommand{
		name:        "look",
		description: "Look at your surroundings",
		aliases:     []string{"l", "lo"},
	}}
}

func (lc *lookCommand) Run(s state.State, args ...string) error {
	if s != nil && s.CurrentRoom() != nil {
		fmt.Println(s.CurrentRoom().Description)
	}
	return nil
}
