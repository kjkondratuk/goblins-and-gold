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
	return &statsCommand{baseCommand{
		name:        "stats",
		description: "Display player stats",
		aliases:     []string{"s", "st"},
		subcommands: nil,
	}}
}

func (sc *statsCommand) Run(s state.State, args ...string) error {
	if s != nil && s.Player() != nil {
		ps, _ := yaml.Marshal(s.Player().Summary())
		pterm.Success.Println(string(ps))
	}
	return nil
}
