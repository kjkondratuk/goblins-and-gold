package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDebugCommand(t *testing.T) {
	c := NewDebugCommand()
	assert.NotNil(t, c)
	assert.Len(t, c.Subcommands(), 2)
	assert.Equal(t, c.Name(), "debug")
	assert.Equal(t, "world", c.Subcommands()[0].Name())
	assert.Equal(t, "help", c.Subcommands()[1].Name())
}

func TestDebugCommand_Run(t *testing.T) {
	c := NewDebugCommand()
	err := c.Run(nil, "help")
	assert.NoError(t, err)
}
