package command

import (
	"github.com/goccy/go-yaml"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/pterm/pterm"
)

type statsCommand struct {
	baseCommand
}

func NewStatsCommand() Command {
	c := &statsCommand{baseCommand{
		name:        "stats",
		description: "Display player stats",
		aliases:     []string{"s", "st"},
		usage:       `stats [help]`,
	}}

	c.subcommands = append(c.subcommands, NewHelpCommand(c))

	return c
}

func (sc *statsCommand) Run(s state.State, args ...string) error {
	if len(args) > 0 {
		err := sc.execSubcommand(s, args...)
		if err != nil {
			return err
		}
		return nil
	}
	if s != nil && s.Player() != nil {
		ps, _ := yaml.Marshal(s.Player().Summary())
		pterm.Success.Println(string(ps))
	}
	return nil
}
