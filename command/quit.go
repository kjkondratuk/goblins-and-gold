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
	return &quitCommand{baseCommand{
		name:        "quit",
		description: "Travel between rooms",
		aliases:     []string{"q"},
	}, exit}
}

func (q *quitCommand) Run(s state.State, args ...string) error {
	pterm.Info.Println("Quitting...")
	q.exit <- syscall.SIGTERM
	return nil
}
