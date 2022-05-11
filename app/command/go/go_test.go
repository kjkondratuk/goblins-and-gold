package _go

import (
	"github.com/kjkondratuk/goblins-and-gold/app/command"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/world/room"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_validate(t *testing.T) {
	t.Run("should pass validation when there are no arguments", func(t *testing.T) {
		err := validate([]string{})
		assert.NoError(t, err)
	})

	t.Run("should not pass validation when there are arguments", func(t *testing.T) {
		err := validate([]string{"some", "arguments"})
		assert.Error(t, err)
	})
}

func Test_action(t *testing.T) {
	t.Run("should not fail when there is nowhere to go", func(t *testing.T) {
		ctx := command.MockContext{}
		ctx.On("State").Return(&state.State{
			CurrRoom: &room.Definition{
				Paths: nil,
			},
		})
		err := action(&ctx)
		assert.NoError(t, err)
	})
}
