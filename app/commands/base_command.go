package commands

import (
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/urfave/cli"
)

type CommandParams struct {
	Name        string
	Aliases     []string
	Usage       string
	Description string
	ArgsUsage   string
	Category    string
}

type baseCommand struct {
	_cmd   cli.Command
	_state *state.State
}

type BaseCommand interface {
	// Action : A function which should handle actions for the command.  It should call
	// the validate function prior to doing anything meaningful.
	Action(*cli.Context) error

	// Validate : validates the cli.Context prior to processing the command action.
	Validate(args ArgProvider) error

	// Command : returns the urfave/cli command representation of this command
	Command() cli.Command
}

type ArgProvider interface {
	Args() cli.Args
}

func newBaseCommand(s *state.State, p CommandParams) baseCommand {
	return baseCommand{
		_cmd: cli.Command{
			Name:        p.Name,
			Aliases:     p.Aliases,
			Usage:       p.Usage,
			Description: p.Description,
			ArgsUsage:   p.ArgsUsage,
			Category:    p.Category,
		},
		_state: s,
	}
}

func (bc *baseCommand) Action(*cli.Context) error {
	panic("Action() not implemented for BaseCommand")
}

func (bc *baseCommand) Command() cli.Command {
	panic("Command() not implemented for BaseCommand")
}

func (bc *baseCommand) commandWithAction(action cli.ActionFunc) cli.Command {
	return cli.Command{
		Name:        bc._cmd.Name,
		Aliases:     bc._cmd.Aliases,
		Usage:       bc._cmd.Usage,
		Description: bc._cmd.Description,
		ArgsUsage:   bc._cmd.ArgsUsage,
		Category:    bc._cmd.Category,
		Action:      action,
	}
}

func (bc *baseCommand) Validate(ctx *cli.Context) error {
	panic("Validate() not implemented for BaseCommand")
}
