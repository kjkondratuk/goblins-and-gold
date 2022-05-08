package commands

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/pterm/pterm"
	"github.com/urfave/cli"
)

type goCommand struct {
	baseCommand
}

type GoCommand interface {
	BaseCommand
}

func NewGoCommand(s *state.State) GoCommand {
	return &goCommand{baseCommand: newBaseCommand(s, CommandParams{
		Name:        "go",
		Aliases:     []string{"g"},
		Usage:       "Travel down a path",
		Description: "Travel down a path",
		ArgsUsage:   "[location number]",
		Category:    "Actions",
	})}
}

func (gc *goCommand) Command() cli.Command {
	return gc.commandWithAction(gc.Action)
}

func (gc *goCommand) Validate(args ArgProvider) error {
	if len(args.Args()) != 0 {
		return errors.New("invalid number of arguments")
	}
	return nil
}

func (gc *goCommand) Action(ctx *cli.Context) error {
	err := gc.Validate(ctx)
	if err != nil {
		return err
	}
	if len(gc._state.CurrRoom.Paths) <= 0 {
		pterm.Error.Println("Nowhere to go!")
		return nil
	}
	paths := make([]ux.Described, len(gc._state.CurrRoom.Paths))
	for i, p := range gc._state.CurrRoom.Paths {
		paths[i] = p
	}

	idx, _, err := gc._state.SelectBuilder.Create("Stay here", "Go").Run(paths)
	if err != nil {
		return err
	}
	if idx == -1 {
		return nil
	}

	// Update the current room based on the selection, unless the user cancels navigation
	nr, _ := gc._state.World.Room(gc._state.CurrRoom.Paths[idx].Room)
	gc._state.CurrRoom = &nr
	gc._state.CurrRoom.RunEncounters(gc._state.Player)
	if gc._state.Player.Unconscious() {
		_ = pterm.DefaultBigText.WithLetters(
			pterm.NewLettersFromStringWithStyle("You died.", pterm.NewStyle(pterm.FgRed)),
		).Render()
		_ = ctx.App.Command("quit").Run(ctx)
	}
	return nil
}
