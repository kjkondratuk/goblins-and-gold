package commands

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

func Stats(s *state.GameState) cli.ActionFunc {
	return func(c *cli.Context) error {
		ps, _ := yaml.Marshal(s.Player)
		fmt.Println(pterm.Green(string(ps)))
		return nil
	}
}
