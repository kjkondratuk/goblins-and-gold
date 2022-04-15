package commands

import (
	"context"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/kjkondratuk/goblins-and-gold/container"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

const (
	interaction = iota
	action
)

func Interact(s *state.GameState) cli.ActionFunc {
	return func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			// coerce to ux.Described
			ia := s.CurrRoom.Containers
			d := make([]ux.Described, len(ia))
			for i, a := range ia {
				d[i] = a
			}
			//convert(s.CurrRoom.Containers)

			// Pass the available containers on context to the interaction selector
			c := context.WithValue(context.Background(), interaction, s.CurrRoom.Containers)

			// Prompt for selection of the interactable
			err := ux.NewSelector("None of these", "Interact with", interactionSelector).Run(c, d)

			// handle errors with creating the selector for item ia
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func actionSelector(ctx context.Context, idx int, val string, err error) error {
	pterm.Success.Printf("%sed %s\n", val, ctx.Value(action))
	return nil
}

func interactionSelector(ctx context.Context, i int, v string, err error) error {
	pterm.Success.Printf("Selected %s\n", v)

	// coerce to ux.Described
	ia := ctx.Value(interaction).([]container.Container)[i-1].SupportedInteractions
	d := make([]ux.Described, len(ia))
	for i, a := range ia {
		d[i] = a
	}

	// Prompt for selection of the action
	err = ux.NewSelector("Cancel", "Actions", actionSelector).Run(context.WithValue(ctx, action, v), d)

	// handle errors with creating the action selector
	if err != nil {
		return err
	}
	return nil
}
