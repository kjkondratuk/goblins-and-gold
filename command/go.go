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

	rm, ok := s.World().Rooms[s.CurrentRoom()]
	if !ok {
		return errors.New(fmt.Sprintf("could not locate room [%s]", s.CurrentRoom()))
	}
	paths := make([]ux.Described, len(rm.Paths))
	if len(rm.Paths) == 0 {
		pterm.Error.Println("Nowhere to go!")
		return nil
	}
	for i, p := range rm.Paths {
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
	s.UpdateCurrentRoom(rm.Paths[idx-1].Room)
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
	r := s.World().Rooms[s.CurrentRoom()]
	for i, e := range r.MandatoryEncounters {
		// TODO : returns an outcome.  do we need it?
		_ = encounter2.NewRunner().Run(s, encounter.NewEncounter(*e))

		// exit if the player is knocked unconscious -- otherwise remove the encounter
		if s.Player().Unconscious() {
			break
		} else {
			if len(r.MandatoryEncounters) > 2 {
				s.World().Rooms[s.CurrentRoom()].MandatoryEncounters = append(r.MandatoryEncounters[:i], r.MandatoryEncounters[i+1:]...)
			} else {
				if i == 0 {
					//r.MandatoryEncounters = r.MandatoryEncounters[i+1:]
					s.World().Rooms[s.CurrentRoom()].MandatoryEncounters = r.MandatoryEncounters[i+1:]
				} else {
					//r.MandatoryEncounters = r.MandatoryEncounters[:i]
					s.World().Rooms[s.CurrentRoom()].MandatoryEncounters = r.MandatoryEncounters[:i]
				}
			}
		}
	}

}
