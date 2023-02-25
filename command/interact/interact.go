package interact

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/command"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
	"reflect"
)

func New(s state.State) cli.Command {
	c := command.NewCommand(command.Params{
		Name:        "interact",
		Aliases:     []string{"i"},
		Usage:       "Interact with your surroundings",
		Description: "Interact with your surroundings",
		Category:    "Actions",
	}, s).Build(command.ValidatorHasArgs, validateState, action)
	return c
}

func validateState(s state.State) error {
	if s == nil || s.CurrentRoom() == nil || s.CurrentRoom().Containers == nil {
		return errors.New("invalid context for interact command")
	}
	return nil
}

func action(s state.State) error {
	// coerce to ux.Described
	ia := s.CurrentRoom().Containers
	if len(ia) <= 0 {
		return nil
	}
	dCon := make([]ux.Described, len(ia))
	for i, a := range ia {
		dCon[i] = a
	}

	// Prompt for selection of the interactable
	interactIdx, _, err := s.Prompter().Select("Interact with", append([]string{"None of these"}, ux.DescribeToList(dCon)...))
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

	_, actStr, err := s.Prompter().Select("Actions", append([]string{"Cancel"}, ux.DescribeToList(dAct)...))
	if err != nil {
		return err
	}
	if actStr == "" {
		return nil
	}

	a := interaction.Type(actStr)

	// get interactions available for this container
	result, err := ia[interactIdx].Do(s, a)
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
	return nil
}
