package commands

import (
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/urfave/cli"
)

func Look(s *state.State) cli.ActionFunc {
	return func(c *cli.Context) error {
		fmt.Println(s.CurrRoom.Description)
		return nil
	}
}
