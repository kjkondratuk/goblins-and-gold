package commands

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/manifoldco/promptui"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

func Go(s *state.State) cli.ActionFunc {
	return func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			var options []string
			options = append(options, "Stay here")
			for _, p := range s.CurrRoom.Paths {
				options = append(options, p.Description)
			}
			p := promptui.Select{
				Label: "Go",
				Items: options,
			}
			i, _, _ := p.Run()
			if i != 0 {
				// Update the current room based on the selection, unless the user cancels navigation
				nr, _ := s.World.Room(s.CurrRoom.Paths[i-1].Room)
				s.CurrRoom = &nr
				s.CurrRoom.RunEncounters(s.Player)
				if s.Player.Unconscious() {
					_ = pterm.DefaultBigText.WithLetters(
							pterm.NewLettersFromStringWithStyle("You died.", pterm.NewStyle(pterm.FgRed))
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
