package command

import (
	"context"
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"testing"
)

func Test_validate(t *testing.T) {
	t.Run("should pass validation when there are no arguments", func(t *testing.T) {
		err := ValidatorHasArgs([]string{})
		assert.NoError(t, err)
	})

	t.Run("should not pass validation when there are arguments", func(t *testing.T) {
		err := ValidatorHasArgs([]string{"some", "arguments"})
		assert.Error(t, err)
	})
}

func Test_command_Build_and_Exec(t *testing.T) {
	s := state.New(nil, nil, nil, nil)
	s.SetContext(context.Background())
	t.Run("should error when an action fails", func(t *testing.T) {
		e := errors.New("there was an error")
		c := NewCommand(Params{
			Name: "command",
		}, s).Build(nil, nil, func(s state.State) error {
			return e
		})

		app := cli.App{
			Commands: []cli.Command{
				c,
			},
		}

		err := app.Run([]string{"", "command"})
		assert.Error(t, err)
		assert.Equal(t, e, err)
		assert.NotNil(t, c.Action)
	})

	t.Run("should fail if validation doesnt pass", func(t *testing.T) {
		e := errors.New("there was an error")
		c := NewCommand(Params{
			Name: "command",
		}, state.New(nil, nil, nil, nil)).Build(func(args []string) error {
			return e
		}, nil, func(s state.State) error {
			t.Logf("executing action handler")
			return nil
		})

		app := cli.App{
			Commands: []cli.Command{
				c,
			},
		}

		err := app.Run([]string{"", "command"})
		assert.Error(t, err)
		assert.Equal(t, e, err)
		assert.NotNil(t, c.Action)
	})

	t.Run("should panic if an action is not specified", func(t *testing.T) {
		c := NewCommand(Params{
			Name: "command",
		}, state.New(nil, nil, nil, nil)).Build(nil, nil, nil)

		app := cli.App{
			Commands: []cli.Command{
				c,
			},
		}

		defer func() {
			if r := recover(); r == nil {
				t.Errorf("should panic when no action is specified")
			}
		}()
		_ = app.Run([]string{"", "command"})
	})

	t.Run("should complete successfully if validation and command action pass", func(t *testing.T) {
		c := NewCommand(Params{
			Name: "command",
		}, state.New(nil, nil, nil, nil)).Build(func(args []string) error {
			return nil
		}, nil, func(s state.State) error {
			t.Logf("executing action handler")
			return nil
		})

		app := cli.App{
			Commands: []cli.Command{
				c,
			},
		}

		err := app.Run([]string{"", "command"})
		assert.NoError(t, err)
		assert.NotNil(t, c.Action)
	})
}
