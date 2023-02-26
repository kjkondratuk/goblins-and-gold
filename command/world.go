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
	return &worldCommand{baseCommand{
		name:        "world",
		description: "Display world information",
		aliases:     []string{"w", "wo"},
	}}
}

func (q *worldCommand) Run(s state.State, args ...string) error {
	ws, _ := yaml.Marshal(s.World())
	pterm.Debug.Println(pterm.Green(string(ws)))
	return nil
}
