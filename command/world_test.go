package command

import (
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWorldCommand(t *testing.T) {
	c := NewWorldCommand()
	assert.NotNil(t, c)
	assert.Len(t, c.Subcommands(), 1)
}

func TestWorldCommand_Run(t *testing.T) {
	t.Run("should execute valid subcommands", func(t *testing.T) {
		c := NewWorldCommand()
		err := c.Run(nil, "help")
		assert.NoError(t, err)
	})

	t.Run("should not execute invalid subcommands", func(t *testing.T) {
		c := NewWorldCommand()
		err := c.Run(nil, "other")
		assert.Error(t, err)
	})

	t.Run("should not error if the world is invalid", func(t *testing.T) {
		c := NewWorldCommand()
		err := c.Run(state.New(nil, nil, nil, nil))
		assert.NoError(t, err)
	})
}
