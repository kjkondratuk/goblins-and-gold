package command

import (
	"github.com/kjkondratuk/goblins-and-gold/state"
)

type app struct {
	baseCommand
}

type App interface {
	Commands() []Command
	Run(s state.State, args ...string) error
}

func NewApp(n string, d string, c ...Command) App {
	a := &app{baseCommand{
		name:        n,
		description: d,
		subcommands: c,
	}}

	return a
}

func (a *app) Commands() []Command {
	return a.subcommands
}

func (a *app) Run(s state.State, args ...string) error {
	return a.execSubcommand(s, args...)
}
