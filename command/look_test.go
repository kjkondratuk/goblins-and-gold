package command

import (
	"github.com/kjkondratuk/goblins-and-gold/model/room"
	"github.com/kjkondratuk/goblins-and-gold/model/world"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLookCommand(t *testing.T) {
	c := NewLookCommand()
	assert.NotNil(t, c)
	assert.Len(t, c.Subcommands(), 1)
}

func TestLookCommand_Run(t *testing.T) {
	t.Run("should execute subcommands", func(t *testing.T) {
		c := NewLookCommand()

		err := c.Run(nil, "help")
		assert.NoError(t, err)
	})

	t.Run("should error on subcommand error", func(t *testing.T) {
		c := NewLookCommand()
		err := c.Run(nil, "other")
		assert.Error(t, err)
	})

	t.Run("should run successfully when state and room are populated", func(t *testing.T) {
		c := NewLookCommand()
		err := c.Run(state.New(nil, nil, "start-room", &world.WorldDefinition{Rooms: map[string]*room.RoomDefinition{
			"start-room": {
				Name:                "",
				Description:         "some description",
				Paths:               nil,
				Containers:          nil,
				MandatoryEncounters: nil,
			},
		}}))
		assert.NoError(t, err)
	})
}
