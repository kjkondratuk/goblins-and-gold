package command

import (
	"github.com/stretchr/testify/assert"
	"os"
	"syscall"
	"testing"
)

func TestNewQuitCommand(t *testing.T) {
	quitchan := make(chan os.Signal)
	c := NewQuitCommand(quitchan)
	assert.NotNil(t, c)
	assert.Len(t, c.Subcommands(), 1)
}

func TestQuitCommand_Run(t *testing.T) {
	t.Run("should run subcommands", func(t *testing.T) {
		quitchan := make(chan os.Signal)
		c := NewQuitCommand(quitchan)
		err := c.Run(nil, "help")
		assert.NoError(t, err)
	})

	t.Run("should error if subcommand is invalid", func(t *testing.T) {
		quitchan := make(chan os.Signal)
		c := NewQuitCommand(quitchan)
		err := c.Run(nil, "other")
		assert.Error(t, err)
	})

	t.Run("should quit", func(t *testing.T) {
		quitchan := make(chan os.Signal)
		defer close(quitchan)
		c := NewQuitCommand(quitchan)
		go func() {
			err := c.Run(nil)
			assert.NoError(t, err)
		}()
		val := <-quitchan
		assert.Equal(t, syscall.SIGTERM, val)
	})
}
