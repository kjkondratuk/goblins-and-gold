package commands

import (
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/world"
	"github.com/urfave/cli"
	"strconv"
)

func Interact(s *state.GameState, w *world.World) cli.ActionFunc {
	return func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			fmt.Println("Interactions:")
			for i, c := range s.CurrRoom.Containers {
				fmt.Println(strconv.Itoa(i) + " - " + c.Description)
			}
		}
		return nil
	}
}
