package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHelpCommand(t *testing.T) {
	c := NewHelpCommand(&baseCommand{
		name:        "test",
		description: "test command",
		usage:       "some usage",
		aliases:     nil,
		subcommands: nil,
	})
	assert.NotNil(t, c)
	assert.Equal(t, "help", c.Name())
	assert.Equal(t, "some usage", c.(*helpCommand).usage)
}

func TestHelpCommand_Run(t *testing.T) {
	t.Run("should run without error when usage is populated", func(t *testing.T) {
		c := NewHelpCommand(&baseCommand{
			name:        "test",
			description: "test command",
			usage:       "some usage",
			aliases:     nil,
			subcommands: nil,
		})

		err := c.Run(nil)
		if err != nil {
			return
		}
	})

	t.Run("should run without error when usage is not populated", func(t *testing.T) {
		c := NewHelpCommand(&baseCommand{
			name:        "test",
			description: "test command",
			usage:       "",
			aliases:     nil,
			subcommands: nil,
		})

		err := c.Run(nil)
		if err != nil {
			return
		}
	})
}
