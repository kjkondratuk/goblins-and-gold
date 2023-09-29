package command

import (
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/pterm/pterm"
)

type lookCommand struct {
	baseCommand
}

func NewLookCommand() Command {
	c := &lookCommand{baseCommand{
		name:        "look",
		description: "Look at your surroundings",
		aliases:     []string{"l", "lo"},
		usage:       `look [help]`,
	}}

	c.subcommands = append(c.Subcommands(), NewHelpCommand(c))

	return c
}

func (lc *lookCommand) Run(s state.State, args ...string) error {
	if len(args) > 0 {
		err := lc.execSubcommand(s, args...)
		if err != nil {
			return err
		}
		return nil
	}
	rm := s.World().Rooms[s.CurrentRoom()]
	if s != nil && s.CurrentRoom() != "" {
		pterm.Info.Println(rm.Description)
	}
	return nil
}
