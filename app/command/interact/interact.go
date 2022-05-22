package interact

import (
	"context"
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/app/command"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
	"reflect"
)

func New(s *state.State) cli.Command {
	c := command.NewCommand(command.Params{
		Name:        "interact",
		Aliases:     []string{"i"},
		Usage:       "Interact with your surroundings",
		Description: "Interact with your surroundings",
		Category:    "Actions",
	}, s).Build(command.ValidatorHasArgs, validateContext, action)
	return c
}

func validateContext(ctx command.Context) error {
	st := ctx.State()
	if st == nil || st.CurrRoom == nil || st.CurrRoom.Containers == nil {
		return errors.New("invalid context for interact command")
	}
	return nil
}

func action(c command.Context) error {
	st := c.State()
	// coerce to ux.Described
	ia := st.CurrRoom.Containers
	if len(ia) <= 0 {
		return nil
	}
	dCon := make([]ux.Described, len(ia))
	for i, a := range ia {
		dCon[i] = a
	}

	// Prompt for selection of the interactable
	interactIdx, _, err := st.SelectBuilder.Create("None of these", "Interact with").Run(dCon)
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
	_, actStr, err := st.SelectBuilder.Create("Cancel", "Actions").Run(dAct)
	if actStr == "" {
		return nil
	}

	a := interaction.Type(actStr)

	// get interactions available for this container
	result, err := ia[interactIdx].Do(context.WithValue(context.Background(), interaction.PlayerDataKey, st.Player), a)
	if err != nil {
		return err
	}

	// Apply to game state, if there was an applicable result
	emptyResult := interaction.Result{}
	if !reflect.DeepEqual(result, emptyResult) {
		//r := result.(interaction.Result)
		st.Apply(result)
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
	return nil
}
