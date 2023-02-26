package command

import (
	"github.com/kjkondratuk/goblins-and-gold/state"
)

type debugCommand struct {
	baseCommand
}

func NewDebugCommand(subcommands ...Command) Command {
	c := &debugCommand{baseCommand{
		name:        "debug",
		description: "Debug commands",
		aliases:     []string{"d", "de"},
		subcommands: subcommands,
		usage:       `debug [help|world]`,
	}}

	c.subcommands = append(c.subcommands, NewHelpCommand(c))

	return c
}

func (q *debugCommand) Run(s state.State, args ...string) error {
	return q.execSubcommand(s, args...)
}
