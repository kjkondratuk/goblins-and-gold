package _go

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/app/command"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
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
	}, s).Build(command.ValidatorHasArgs, action)
	return c
}

func action(ctx command.Context) error {
	st := ctx.State()
	// TODO : should probably move this to a context validator (separate from argument validator) and return an error
	if st == nil || st.CurrRoom == nil || st.CurrRoom.Paths == nil {
		return errors.New("invalid context for go command")
	}
	// TODO : add test coverage for this condition
	if len(st.CurrRoom.Paths) <= 0 {
		pterm.Error.Println("Nowhere to go!")
		return nil
	}
	paths := make([]ux.Described, len(st.CurrRoom.Paths))
	for i, p := range st.CurrRoom.Paths {
		paths[i] = p
	}

	idx, _, err := st.SelectBuilder.Create("Stay here", "Go").Run(paths)
	if err != nil {
		return err
	}
	if idx == -1 {
		return nil
	}

	// Update the current room based on the selection, unless the user cancels navigation
	nr, _ := st.World.Room(st.CurrRoom.Paths[idx].Room)
	st.CurrRoom = &nr
	st.CurrRoom.RunEncounters(st.Player)
	if st.Player.Unconscious() {
		_ = pterm.DefaultBigText.WithLetters(
			pterm.NewLettersFromStringWithStyle("You died.", pterm.NewStyle(pterm.FgRed)),
		).Render()
		_ = ctx.RunCommandByName("quit")
	}
	return nil
}
