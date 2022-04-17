package commands

import (
	"context"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/kjkondratuk/goblins-and-gold/container"
	interaction2 "github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

func Interact(s *state.State) cli.ActionFunc {
	return func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			// coerce to ux.Described
			ia := s.CurrRoom.Containers
			d := make([]ux.Described, len(ia))
			for i, a := range ia {
				d[i] = a
			}

			// Pass the available containers on context to the interactionData selector
			c := context.WithValue(context.Background(), interaction2.InteractionDataKey, s.CurrRoom.Containers)
			c = context.WithValue(c, interaction2.PlayerDataKey, s.Player)

			// Prompt for selection of the interactable
			result, err := ux.NewSelector("None of these", "Interact with", interactionSelector).Run(c, d)
			// handle errors with creating the selector for item ia
			if err != nil {
				return err
			}

			// Apply to game state, if there was an applicable result
			if result != nil {
				r := result.(interaction2.Result)
				s.Apply(r)
				switch r.Type {
				case interaction2.RT_Success:
					pterm.Success.Println(r.Message)
				case interaction2.RT_Failure:
					pterm.Error.Println(r.Message)
				default:
					pterm.Info.Println(r.Message)
				}
			}
		}
		return nil
	}
}

// actionSelector : Selects an actionItemData to take upon a given item provided a context containing both the
func actionSelector(ctx context.Context, idx int, val string, err error) (interface{}, error) {
	a := interaction2.Type(val)

	// get interactions available for this container
	containerInteractions := ctx.Value(interaction2.ContainerDataKey).(*container.Container)

	result, err := containerInteractions.Do(ctx, a)
	if err != nil {
		return interaction2.Result{}, err
	}

	return result, nil
}

// interactionSelector : Selects an object to interact with provided a context with interactables.  Further
// prompts a user for the actionItemData they'd like to take on the item based on its supported interactions.
func interactionSelector(ctx context.Context, idx int, val string, err error) (interface{}, error) {
	pterm.Success.Printf("Selected: %s\n", val)
	selection := ctx.Value(interaction2.InteractionDataKey).([]*container.Container)[idx-1]

	// coerce to ux.Described
	ia := selection.SupportedInteractions
	d := make([]ux.Described, len(ia))
	for i, a := range ia {
		d[i] = a
	}

	// Prompt for selection of the actionItemData
	return ux.NewSelector("Cancel", "Actions", actionSelector).Run(context.WithValue(ctx, interaction2.ContainerDataKey, selection), d)
}
