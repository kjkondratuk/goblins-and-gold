package look

import (
	"errors"
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/command"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/urfave/cli"
)

func New(s state.State) cli.Command {
	c := command.NewCommand(command.Params{
		Name:        "look",
		Aliases:     []string{"l"},
		Usage:       "Look at your surroundings",
		Description: "Look at your surroundings",
		Category:    "Info",
	}, s).Build(nil, validateState, action)
	return c
}

func validateState(s state.State) error {
	if s == nil || s.CurrentRoom() == nil {
		return errors.New("invalid context for look command")
	}
	return nil
}

func action(s state.State) error {
	if s != nil && s.CurrentRoom() != nil {
		fmt.Println(s.CurrentRoom().Description)
	}
	return nil
}
