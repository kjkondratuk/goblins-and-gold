package _go

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/command"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

func New(s state.State) cli.Command {
	c := command.NewCommand(command.Params{
		Name:        "go",
		Aliases:     []string{"g"},
		Usage:       "Travel down a path",
		Description: "Travel down a path",
		ArgsUsage:   "[location number]",
		Category:    "Actions",
	}, s).Build(command.ValidatorHasArgs, validateContext, action)
	return c
}

func validateContext(s state.State) error {
	if s == nil || s.CurrentRoom() == nil || s.CurrentRoom().Paths == nil {
		return errors.New("invalid context for go command")
	}
	return nil
}

func action(s state.State) error {
	// TODO : add test coverage for this condition
	if len(s.CurrentRoom().Paths) <= 0 {
		pterm.Error.Println("Nowhere to go!")
		return nil
	}
	paths := make([]ux.Described, len(s.CurrentRoom().Paths))
	for i, p := range s.CurrentRoom().Paths {
		paths[i] = p
	}

	idx, _, err := s.Prompter().Select("Go", append(ux.DescribeToList(paths), "Stay Here"))
	if err != nil {
		return err
	}
	if idx == -1 {
		return nil
	}

	// Update the current room based on the selection, unless the user cancels navigation
	nr, _ := s.World().Room(s.CurrentRoom().Paths[idx].Room)
	s.UpdateCurrentRoom(&nr)
	s.CurrentRoom().RunEncounters(s)
	if s.Player().Unconscious() {
		_ = pterm.DefaultBigText.WithLetters(
			pterm.NewLettersFromStringWithStyle("You died.", pterm.NewStyle(pterm.FgRed)),
		).Render()
		_ = s.CliContext().App.Command("quit").Run(s.CliContext())
	}
	return nil
}
