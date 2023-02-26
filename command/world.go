package command

import (
	"github.com/goccy/go-yaml"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/pterm/pterm"
)

type worldCommand struct {
	baseCommand
}

func NewWorldCommand() Command {
	c := &worldCommand{baseCommand{
		name:        "world",
		description: "Display world information",
		aliases:     []string{"w", "wo"},
		usage:       `debug world [help]`,
	}}

	c.subcommands = append(c.subcommands, NewHelpCommand(c))

	return c
}

func (q *worldCommand) Run(s state.State, args ...string) error {
	if len(args) > 0 {
		err := q.execSubcommand(s, args...)
		if err != nil {
			return err
		}
		return nil
	}
	ws, _ := yaml.Marshal(s.World())
	pterm.Debug.Println(pterm.Green(string(ws)))
	return nil
}
