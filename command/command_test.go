package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBaseCommand_Name(t *testing.T) {
	c := baseCommand{name: "some name"}
	assert.Equal(t, "some name", c.Name())
}

func TestBaseCommand_Usage(t *testing.T) {
	c := baseCommand{usage: "some usage string"}
	assert.Equal(t, "some usage string", c.Usage())
}

func TestBaseCommand_Description(t *testing.T) {
	c := baseCommand{description: "some description"}
	assert.Equal(t, "some description", c.Description())
}

func TestBaseCommand_Aliases(t *testing.T) {
	c := baseCommand{aliases: []string{"a1", "a2"}}
	assert.Equal(t, []string{"a1", "a2"}, c.Aliases())
}

func TestBaseCommand_Subcommands(t *testing.T) {
	c := baseCommand{subcommands: []Command{
		&baseCommand{name: "cmd1"},
		&baseCommand{name: "cmd2"},
	}}
	assert.Len(t, c.Subcommands(), 2)
	assert.Equal(t, "cmd1", c.Subcommands()[0].Name())
	assert.Equal(t, "cmd2", c.Subcommands()[1].Name())
}

func TestBaseCommand_Run(t *testing.T) {
	c := baseCommand{}
	assert.NoError(t, c.Run(nil))
}

func TestBaseCommand_execSubcommand(t *testing.T) {
	c := baseCommand{
		name: "cmd1",
		subcommands: []Command{
			&baseCommand{
				name:    "subcmd1",
				aliases: []string{"sc1"},
			},
		},
	}

	t.Run("should run successfully when the command is a subcommand", func(t *testing.T) {
		err := c.execSubcommand(nil, []string{"subcmd1"}...)
		assert.NoError(t, err)
	})

	t.Run("should error if the command is not a subcommand", func(t *testing.T) {
		err := c.execSubcommand(nil, []string{"xyz"}...)
		assert.Error(t, err)
	})

	t.Run("should execute command based on valid alias", func(t *testing.T) {
		err := c.execSubcommand(nil, []string{"sc1"}...)
		assert.NoError(t, err)
	})
}
