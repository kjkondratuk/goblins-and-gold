package commands

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

func Go(s *state.State) cli.ActionFunc {
	return func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			paths := make([]ux.Described, len(s.CurrRoom.Paths))
			for i, p := range s.CurrRoom.Paths {
				paths[i] = p
			}

			idx, _, err := ux.NewSelector("Stay here", "Go").Run(paths)
			if err != nil {
				return err
			}

			if idx != 0 {
				// Update the current room based on the selection, unless the user cancels navigation
				nr, _ := s.World.Room(s.CurrRoom.Paths[idx].Room)
				s.CurrRoom = &nr
				s.CurrRoom.RunEncounters(s.Player)
				if s.Player.Unconscious() {
					_ = pterm.DefaultBigText.WithLetters(
						pterm.NewLettersFromStringWithStyle("You died.", pterm.NewStyle(pterm.FgRed)),
					).Render()
					_ = c.App.Command("quit").Run(c)
				}
			}
		} else {
			return errors.New("invalid number of arguments")
		}
		return nil
	}
}
