package command

import (
	"errors"
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/pterm/pterm"
)

type baseCommand struct {
	name        string
	description string
	usage       string
	aliases     []string
	subcommands []Command
}

type Command interface {
	Name() string
	Usage() string
	Description() string
	Aliases() []string
	Subcommands() []Command
	Run(s state.State, args ...string) error
}

func (b *baseCommand) Name() string {
	return b.name
}

func (b *baseCommand) Usage() string {
	return b.usage
}

func (b *baseCommand) Description() string {
	return b.description
}

func (b *baseCommand) Aliases() []string {
	return b.aliases
}
func (b *baseCommand) Subcommands() []Command {
	return b.subcommands
}

func (b *baseCommand) Run(s state.State, args ...string) error {
	return nil
}

func (b *baseCommand) execSubcommand(s state.State, args ...string) error {
	if len(args) > 0 {
		for _, c := range b.subcommands {
			isAlias := false
			for _, alias := range c.Aliases() {
				if alias == args[0] {
					isAlias = true
				}
			}
			if c.Name() == args[0] || isAlias {
				return c.Run(s, args[1:]...)
			}
		}
		return errors.New(fmt.Sprintf("no subcommand of %s found for %s", b.name, args[0]))
	}

	// command has no action itself, so print usage
	if len(b.Usage()) > 0 {
		pterm.Info.Println(b.Usage())
	}
	return nil
}
