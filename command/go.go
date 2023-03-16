package command

import (
	"errors"
	"fmt"
	encounter2 "github.com/kjkondratuk/goblins-and-gold/encounter"
	"github.com/kjkondratuk/goblins-and-gold/model/encounter"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"github.com/pterm/pterm"
)

type goCommand struct {
	baseCommand
	qc Command
	lc Command
}

func NewGoCommand(quit Command, look Command) Command {
	c := &goCommand{baseCommand{
		name:        "go",
		description: "Travel between rooms",
		aliases:     []string{"g"},
		usage:       `go [help]`,
	}, quit, look}

	c.subcommands = append(c.subcommands, NewHelpCommand(c))

	return c
}

func (g *goCommand) Run(s state.State, args ...string) error {
	// TODO : need to refactor this so it isn't in every command
	if len(args) > 0 {
		return g.execSubcommand(s, args...)
	}
	if len(s.CurrentRoom().Paths) <= 0 {
		pterm.Error.Println("Nowhere to go!")
		return nil
	}
	paths := make([]ux.Described, len(s.CurrentRoom().Paths))
	for i, p := range s.CurrentRoom().Paths {
		paths[i] = p
	}

	idx, _, err := s.Prompter().Select("Go", append([]string{"Stay Here"}, ux.DescribeToList(paths)...))
	if err != nil {
		return err
	}
	if idx == 0 {
		return nil
	}

	// Update the current room based on the selection, unless the user cancels navigation
	rm := s.CurrentRoom().Paths[idx-1].Room
	nr, ok := s.World().Room(rm)
	if !ok {
		return errors.New(fmt.Sprintf("could not locate room [%s]", rm))
	}
	s.UpdateCurrentRoom(&nr)
	g.runEncounters(s)
	if s.Player().Unconscious() {
		_ = pterm.DefaultBigText.WithLetters(
			pterm.NewLettersFromStringWithStyle("You died.", pterm.NewStyle(pterm.FgRed)),
		).Render()
		_ = g.qc.Run(s)
	}
	return g.lc.Run(s)
}

func (g *goCommand) runEncounters(s state.State) {
	for _, e := range s.CurrentRoom().MandatoryEncounters {
		// TODO : returns an outcome.  do we need it?
		_ = encounter2.NewRunner().Run(s, encounter.NewEncounter(*e))

		p := s.Player()
		if p.Unconscious() {
			break
		}
	}

}
