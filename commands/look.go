package commands

import (
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/urfave/cli"
)

func Look(s *state.GameState) cli.ActionFunc {
	return func(c *cli.Context) error {
		fmt.Println(s.CurrRoom.Description)
		return nil
	}
}
