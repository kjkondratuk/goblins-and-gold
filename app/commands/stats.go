package commands

import (
	"github.com/goccy/go-yaml"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

func Stats(s *state.State) cli.ActionFunc {
	return func(c *cli.Context) error {
		ps, _ := yaml.Marshal(s.Player.Definition())
		pterm.Success.Println(string(ps))
		return nil
	}
}
