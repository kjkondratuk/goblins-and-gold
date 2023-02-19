package _go

import (
	"context"
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/app/command"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

func New(s *state.State) cli.Command {
	c := command.NewCommand(command.Params{
		Name:        "go",
		Aliases:     []string{"g"},
		Usage:       "Travel down a path",
		Description: "Travel down a path",
		ArgsUsage:   "[location number]",
		Category:    "Actions",
	}, s).Build(command.ValidatorHasArgs, validateContext, action(s))
	return c
}

func validateContext(ctx command.Context) error {
	st := ctx.State()
	if st == nil || st.CurrRoom == nil || st.CurrRoom.Paths == nil {
		return errors.New("invalid context for go command")
	}
	return nil
}

func action(s *state.State) func(ctx command.Context) error {
	return func(ctx command.Context) error {
		// TODO : add test coverage for this condition
		if len(s.CurrRoom.Paths) <= 0 {
			pterm.Error.Println("Nowhere to go!")
			return nil
		}
		paths := make([]ux.Described, len(s.CurrRoom.Paths))
		for i, p := range s.CurrRoom.Paths {
			paths[i] = p
		}

		idx, _, err := s.PromptLib.Select("Go", append(ux.DescribeToList(paths), "Stay Here"))
		if err != nil {
			return err
		}
		if idx == -1 {
			return nil
		}

		// Update the current room based on the selection, unless the user cancels navigation
		nr, _ := s.World.Room(s.CurrRoom.Paths[idx].Room)
		s.CurrRoom = &nr
		s.CurrRoom.RunEncounters(s.AsContext(context.Background()))
		if s.Player.Unconscious() {
			_ = pterm.DefaultBigText.WithLetters(
				pterm.NewLettersFromStringWithStyle("You died.", pterm.NewStyle(pterm.FgRed)),
			).Render()
			_ = ctx.RunCommandByName("quit")
		}
		return nil
	}
}
