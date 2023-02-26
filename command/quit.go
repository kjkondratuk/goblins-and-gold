package command

import (
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/pterm/pterm"
	"os"
	"syscall"
)

type quitCommand struct {
	baseCommand
	exit chan os.Signal
}

func NewQuitCommand(exit chan os.Signal) Command {
	c := &quitCommand{baseCommand{
		name:        "quit",
		description: "Travel between rooms",
		aliases:     []string{"q"},
		usage:       `quit [help]`,
	}, exit}

	c.subcommands = append(c.subcommands, NewHelpCommand(c))

	return c
}

func (q *quitCommand) Run(s state.State, args ...string) error {
	if len(args) > 0 {
		err := q.execSubcommand(s, args...)
		if err != nil {
			return err
		}
		return nil
	}
	pterm.Info.Println("Quitting...")
	q.exit <- syscall.SIGTERM
	return nil
}
