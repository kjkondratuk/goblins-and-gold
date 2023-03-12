package command

import (
	"github.com/kjkondratuk/goblins-and-gold/container"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"github.com/pterm/pterm"
	"reflect"
)

type interactCommand struct {
	baseCommand
	cc container.ContainerController
	sa state.Applier[interaction.Result]
}

func NewInteractCommand(cc container.ContainerController, sa state.Applier[interaction.Result]) Command {
	c := &interactCommand{baseCommand{
		name:        "interact",
		description: "Travel between rooms",
		aliases:     []string{"i", "in"},
		usage:       `interact [help]`,
	}, cc, sa}

	c.subcommands = append(c.subcommands, NewHelpCommand(c))

	return c
}

func (ic *interactCommand) Run(s state.State, args ...string) error {
	if len(args) > 0 {
		return ic.execSubcommand(s, "help")
	}
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
	interactIdx, _, err := s.Prompter().Select("Interact with",
		append([]string{"None of these"}, ux.DescribeToList(dCon)...))
	// handle errors with creating the selector for item ia
	if err != nil {
		return err
	}
	if interactIdx <= 0 || ia[interactIdx-1] == nil {
		return nil
	}

	// Prompt for selection of the action
	dAct := make([]ux.Described, len(ia[interactIdx-1].SupportedInteractions))
	for i, a := range ia[interactIdx-1].SupportedInteractions {
		dAct[i] = a
	}
	actIdx, actStr, err := s.Prompter().Select("Actions", append([]string{"Cancel"},
		ux.DescribeToList(dAct)...))
	if err != nil {
		return err
	}
	if actIdx <= 0 {
		return nil
	}

	a := interaction.Type(actStr)

	// get interactions available for this container
	result, err := ic.cc.Do(s, ia[interactIdx-1], a)
	if err != nil {
		return err
	}

	// Apply to game state, if there was an applicable result
	emptyResult := interaction.Result{}
	if !reflect.DeepEqual(result, emptyResult) {
		// TODO : should probably refactor this to an interaction result state applicator
		ic.sa.Apply(s, result)
		switch result.Type {
		case interaction.RtSuccess:
			pterm.Success.Println(result.Message)
		case interaction.RtFailure:
			pterm.Error.Println(result.Message)
		default:
			if result.Message != "" {
				pterm.Info.Println(result.Message)
			}
		}
	}
	return nil
}
