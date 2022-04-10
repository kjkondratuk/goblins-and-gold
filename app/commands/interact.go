package commands

import (
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/kjkondratuk/goblins-and-gold/world"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

func Interact(s *state.GameState, w *world.World) cli.ActionFunc {
	return func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			interactions := s.CurrRoom.Containers
			d := make([]ux.Described, len(interactions))
			for i, a := range interactions {
				d[i] = a
			}

			// Prompt for selection of the interactable
			err := ux.NewSelector("None of these", "Interact with", func(i int, v string, err error) error {
				pterm.Success.Printf("Selected %s\n", v)
				// coerce
				interactions := s.CurrRoom.Containers[i-1].SupportedInteractions
				d := make([]ux.Described, len(interactions))
				for i, a := range interactions {
					d[i] = &a
				}

				// Prompt for selection of the action
				err = ux.NewSelector("Cancel", "Actions", func(idx int, val string, err error) error {
					pterm.Success.Printf("%sed %s\n", val, v)
					return nil
				}).Run(d)
				if err != nil {
					return err
				}
				return nil
			}).Run(d)
			if err != nil {
				return err
			}

			//var options []string
			//options = append(options, "None of these")
			//for _, c := range s.CurrRoom.Containers {
			//	options = append(options, c.Description)
			//}
			//p := promptui.Select{Label: "Interact With", Items: options}
			//i, v, _ := p.Run()
			//if i > 0 {
			//
			//}
		}
		return nil
	}
}
