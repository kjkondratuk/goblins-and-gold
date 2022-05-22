package look

import (
	"errors"
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
	}, s).Build(nil, validateContext, action)
	return c
}

func validateContext(ctx command.Context) error {
	if ctx.State() == nil || ctx.State().CurrRoom == nil {
		return errors.New("invalid context for look command")
	}
	return nil
}

func action(c command.Context) error {
	if c.State() != nil && c.State().CurrRoom != nil {
		fmt.Println(c.State().CurrRoom.Description)
	}
	return nil
}
