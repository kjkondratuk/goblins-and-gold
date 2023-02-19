package command

import (
	"errors"
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/urfave/cli"
)

var (
	ValidatorHasArgs = func(args []string) error {
		if len(args) != 0 {
			return errors.New("invalid number of arguments")
		}
		return nil
	}
)

type ArgValidator func(args []string) error

type ContextValidator func(ctx Context) error

type Action func(ctx Context) error

type ctx struct {
	_c *cli.Context
	_s *state.State
}

type Context interface {
	State() *state.State
	Context() *cli.Context
	RunCommandByName(name string) error
}

type command struct {
	*state.State
	*cli.Command
}

type Command interface {
	Build(argValidator ArgValidator, ctxValidator ContextValidator, action Action) cli.Command
}

type Params struct {
	Name        string
	Aliases     []string
	Usage       string
	ArgsUsage   string
	Description string
	Category    string
}

func NewCommand(c Params, s *state.State) Command {
	return &command{
		State: s,
		Command: &cli.Command{
			Name:        c.Name,
			Aliases:     c.Aliases,
			Usage:       c.Usage,
			ArgsUsage:   c.ArgsUsage,
			Description: c.Description,
			Category:    c.Category,
		},
	}
}

func (c *command) Build(argValidator ArgValidator, ctxValidator ContextValidator, action Action) cli.Command {
	c.Action = cli.ActionFunc(func(con *cli.Context) error {
		if argValidator != nil {
			err := argValidator(con.Args())
			if err != nil {
				return err
			}
		}
		if action != nil {
			cx := ctx{con, c.State}
			// TODO : add test coverage for context validator
			if ctxValidator != nil {
				err := ctxValidator(&cx)
				if err != nil {
					return err
				}
			}
			err := action(&cx)
			if err != nil {
				return err
			}
		} else {
			panic(fmt.Sprintf("must specify an action to build command: "))
		}
		return nil
	})
	return *c.Command
}

func (c *ctx) State() *state.State {
	return c._s
}

func (c *ctx) Context() *cli.Context {
	return c._c
}

func (c *ctx) RunCommandByName(name string) error {
	return c._c.App.Command(name).Run(c._c)
}
