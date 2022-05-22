package stats

import (
	"errors"
	"github.com/goccy/go-yaml"
	"github.com/kjkondratuk/goblins-and-gold/app/command"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

func New(s *state.State) cli.Command {
	c := command.NewCommand(command.Params{
		Name:        "stats",
		Aliases:     []string{"s"},
		Usage:       "Interrogate your player stats",
		Description: "Interrogate your player stats",
		Category:    "Info",
	}, s).Build(nil, validateContext, action)
	return c
}

func validateContext(ctx command.Context) error {
	if ctx.State() == nil || ctx.State().Player == nil {
		return errors.New("invalid context for stats command")
	}
	return nil
}

func action(c command.Context) error {
	if c.State() != nil && c.State().Player != nil {
		ps, _ := yaml.Marshal(c.State().Player.Summary())
		pterm.Success.Println(string(ps))
	}
	return nil
}
