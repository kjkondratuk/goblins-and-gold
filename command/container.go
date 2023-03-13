package command

import (
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/state"
)

type containerCommand struct {
	baseCommand
}

func NewContainerCommand() Command {
	c := &containerCommand{baseCommand{
		name:        "container",
		description: "Interact with a container",
		aliases:     []string{"c"},
		usage:       `interact container [help]`,
	}}

	c.subcommands = append(c.subcommands, NewHelpCommand(c))

	return c
}

func (c *containerCommand) Run(s state.State, args ...string) error {
	fmt.Println(args)
	return nil
}
