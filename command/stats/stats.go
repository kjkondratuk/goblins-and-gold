package stats

import (
	"errors"
	"github.com/goccy/go-yaml"
	"github.com/kjkondratuk/goblins-and-gold/command"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

func New(s state.State) cli.Command {
	c := command.NewCommand(command.Params{
		Name:        "stats",
		Aliases:     []string{"s"},
		Usage:       "Interrogate your player stats",
		Description: "Interrogate your player stats",
		Category:    "Info",
	}, s).Build(nil, validateContext, action)
	return c
}

func validateContext(s state.State) error {
	if s == nil || s.Player() == nil {
		return errors.New("invalid context for stats command")
	}
	return nil
}

func action(s state.State) error {
	if s != nil && s.Player() != nil {
		ps, _ := yaml.Marshal(s.Player().Summary())
		pterm.Success.Println(string(ps))
	}
	return nil
}
