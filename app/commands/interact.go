package commands

import (
	"context"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
	"reflect"
)

func Interact(s *state.State) cli.ActionFunc {
	return func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			// coerce to ux.Described
			ia := s.CurrRoom.Containers
			if len(ia) <= 0 {
				return nil
			}
			dCon := make([]ux.Described, len(ia))
			for i, a := range ia {
				dCon[i] = a
			}

			// Pass the available containers on context to the interactionData selector
			//c := context.WithValue(context.Background(), interaction.InteractionDataKey, s.CurrRoom.Containers)
			//p := context.WithValue(context.Background(), interaction.PlayerDataKey, s.Player)

			// Prompt for selection of the interactable
			interactIdx, _, err := s.SelectBuilder.Create("None of these", "Interact with").Run(dCon)
			// handle errors with creating the selector for item ia
			if err != nil {
				return err
			}
			if interactIdx < 0 {
				return nil
			}

			// Prompt for selection of the action
			dAct := make([]ux.Described, len(ia[interactIdx].SupportedInteractions))
			for i, a := range ia[interactIdx].SupportedInteractions {
				dAct[i] = a
			}
			_, actStr, err := s.SelectBuilder.Create("Cancel", "Actions").Run(dAct)
			if actStr == "" {
				return nil
			}

			a := interaction.Type(actStr)

			// get interactions available for this container
			result, err := ia[interactIdx].Do(context.WithValue(context.Background(), interaction.PlayerDataKey, s.Player), a)
			if err != nil {
				return err
			}

			// Apply to game state, if there was an applicable result
			emptyResult := interaction.Result{}
			if !reflect.DeepEqual(result, emptyResult) {
				//r := result.(interaction.Result)
				s.Apply(result)
				switch result.Type {
				case interaction.RT_Success:
					pterm.Success.Println(result.Message)
				case interaction.RT_Failure:
					pterm.Error.Println(result.Message)
				default:
					if result.Message != "" {
						pterm.Info.Println(result.Message)
					}
				}
			}
		}
		return nil
	}
}
