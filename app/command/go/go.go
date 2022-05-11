package _go

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/app/command"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

func NewGoCommand(s *state.State) cli.Command {
	c := command.NewCommand(command.Params{
		Name:        "go",
		Aliases:     []string{"g"},
		Usage:       "Travel down a path",
		Description: "Travel down a path",
		ArgsUsage:   "[location number]",
		Category:    "Actions",
	}, s).Build(validate, action)
	return c
}

func validate(args []string) error {
	if len(args) != 0 {
		return errors.New("invalid number of arguments")
	}
	return nil
}

func action(ctx command.Context) error {
	if len(ctx.State().CurrRoom.Paths) <= 0 {
		pterm.Error.Println("Nowhere to go!")
		return nil
	}
	paths := make([]ux.Described, len(ctx.State().CurrRoom.Paths))
	for i, p := range ctx.State().CurrRoom.Paths {
		paths[i] = p
	}

	idx, _, err := ctx.State().SelectBuilder.Create("Stay here", "Go").Run(paths)
	if err != nil {
		return err
	}
	if idx == -1 {
		return nil
	}

	// Update the current room based on the selection, unless the user cancels navigation
	nr, _ := ctx.State().World.Room(ctx.State().CurrRoom.Paths[idx].Room)
	ctx.State().CurrRoom = &nr
	ctx.State().CurrRoom.RunEncounters(ctx.State().Player)
	if ctx.State().Player.Unconscious() {
		_ = pterm.DefaultBigText.WithLetters(
			pterm.NewLettersFromStringWithStyle("You died.", pterm.NewStyle(pterm.FgRed)),
		).Render()
		_ = ctx.RunCommandByName("quit")
	}
	return nil
}
