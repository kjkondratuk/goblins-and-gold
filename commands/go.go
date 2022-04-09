package commands

import (
	"errors"
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/world"
	"github.com/urfave/cli"
	"strconv"
)

func Go(s *state.GameState, w *world.World) cli.ActionFunc {
	return func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			fmt.Println("Paths:")
			for i, p := range s.CurrRoom.Paths {
				fmt.Println(strconv.Itoa(i) + " - " + p.Description)
			}
		} else if len(c.Args()) == 1 {
			pi, err := strconv.Atoi(c.Args()[0])
			if err != nil {
				return errors.New("argument should be numeric")
			}
			s.CurrRoom, _ = w.Room(s.CurrRoom.Paths[pi].Room)
		} else {
			return errors.New("invalid number of arguments")
		}
		return nil
	}
}
