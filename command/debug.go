package command

import (
	"github.com/kjkondratuk/goblins-and-gold/state"
)

type debugCommand struct {
	baseCommand
}

func NewDebugCommand(subcommands ...Command) Command {
	return &debugCommand{baseCommand{
		name:        "debug",
		description: "Debug commands",
		aliases:     []string{"d", "de"},
		subcommands: subcommands,
	}}
}

func (q *debugCommand) Run(s state.State, args ...string) error {
	return q.execSubcommand(s, args...)
}
