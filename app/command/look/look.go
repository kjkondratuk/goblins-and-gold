package look

import (
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/app/command"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/urfave/cli"
)

func New(s *state.State) cli.Command {
	c := command.NewCommand(command.Params{
		Name:        "look",
		Aliases:     []string{"l"},
		Usage:       "Look at your surroundings",
		Description: "Look at your surroundings",
		Category:    "Info",
	}, s).Build(nil, action)
	return c
}

func action(c command.Context) error {
	if c.State() != nil && c.State().CurrRoom != nil {
		fmt.Println(c.State().CurrRoom.Description)
	}
	return nil
}
