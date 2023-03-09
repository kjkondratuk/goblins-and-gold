package command

import (
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/pterm/pterm"
)

type helpCommand struct {
	baseCommand
	usage string
}

func NewHelpCommand(p Command) Command {
	return &helpCommand{baseCommand{
		name:        "help",
		aliases:     []string{"h"},
		description: "Help information about the command",
		usage:       fmt.Sprintf("%s help", p.Name()),
	}, p.Usage()}
}

func (hc *helpCommand) Run(s state.State, args ...string) error {
	if len(hc.usage) > 0 {
		pterm.Info.Printfln("Usage: %s", hc.usage)
	}
	return nil
}
