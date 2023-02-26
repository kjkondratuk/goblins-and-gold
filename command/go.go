package command

import (
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"github.com/pterm/pterm"
)

type goCommand struct {
	baseCommand
	qc Command
}

func NewGoCommand(quit Command) Command {
	return &goCommand{baseCommand{
		name:        "go",
		description: "Travel between rooms",
		aliases:     []string{"g"},
	}, quit}
}

func (g *goCommand) Run(s state.State, args ...string) error {
	// TODO : add test coverage for this condition
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
	nr, _ := s.World().Room(s.CurrentRoom().Paths[idx-1].Room)
	s.UpdateCurrentRoom(&nr)
	s.CurrentRoom().RunEncounters(s)
	if s.Player().Unconscious() {
		_ = pterm.DefaultBigText.WithLetters(
			pterm.NewLettersFromStringWithStyle("You died.", pterm.NewStyle(pterm.FgRed)),
		).Render()
		_ = g.qc.Run(s)
	}
	return nil
}
