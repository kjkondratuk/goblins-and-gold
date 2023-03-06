package command

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/ux/mock"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"testing"
)

func TestNewInteractCommand(t *testing.T) {
	c := NewInteractCommand()
	assert.NotNil(t, c)
	assert.Equal(t, c.Name(), "interact")
	assert.Len(t, c.Subcommands(), 1)
}

func TestInteractCommand_Run(t *testing.T) {
	t.Run("should invoke help when args are supplied", func(t *testing.T) {
		c := NewInteractCommand()
		err := c.Run(nil, "some arg")
		assert.NoError(t, err)
	})

	t.Run("should handle nil containers to interact with", func(t *testing.T) {
		c := NewInteractCommand()
		err := c.Run(state.New(nil, nil, &state.RoomDefinition{
			Name: "some-room",
		}, nil))
		assert.NoError(t, err)
	})

	t.Run("should handle no containers to interact with", func(t *testing.T) {
		c := NewInteractCommand()
		err := c.Run(state.New(nil, nil, &state.RoomDefinition{
			Name:       "some-room",
			Containers: []*state.Container{},
		}, nil))
		assert.NoError(t, err)
	})

	t.Run("should error for interaction select error", func(t *testing.T) {
		pm := &mock.PromptMock{}

		pm.On("Select", mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(-1, "", errors.New("an error occurred"))

		c := NewInteractCommand()
		err := c.Run(state.New(pm, nil, &state.RoomDefinition{
			Name: "some-room",
			Containers: []*state.Container{
				{
					Type:        "Chest",
					Description: "some container",
				},
			},
		}, nil))
		assert.Error(t, err)
	})

	t.Run("should not error for interaction select cancel", func(t *testing.T) {
		pm := &mock.PromptMock{}

		pm.On("Select", mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(0, "", nil)

		c := NewInteractCommand()
		err := c.Run(state.New(pm, nil, &state.RoomDefinition{
			Name: "some-room",
			Containers: []*state.Container{
				{
					Type:        "Chest",
					Description: "some container",
				},
			},
		}, nil))
		assert.NoError(t, err)
	})

	t.Run("should error on action select failure", func(t *testing.T) {
		pm := &mock.PromptMock{}

		pm.On("Select", mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "some container", nil).Times(1)
		pm.On("Select", mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(-1, "", errors.New("something went wrong")).Times(1)

		c := NewInteractCommand()
		err := c.Run(state.New(pm, nil, &state.RoomDefinition{
			Name: "some-room",
			Containers: []*state.Container{
				{
					Type:        "Chest",
					Description: "some container",
					SupportedInteractions: []interaction.Type{
						"Open",
						"Loot",
					},
				},
			},
		}, nil))
		assert.Error(t, err)
	})

	t.Run("should succeed on action select cancel", func(t *testing.T) {
		pm := &mock.PromptMock{}

		pm.On("Select", mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "some container", nil).Times(1)
		pm.On("Select", mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(0, "", nil).Times(1)

		c := NewInteractCommand()
		err := c.Run(state.New(pm, nil, &state.RoomDefinition{
			Name: "some-room",
			Containers: []*state.Container{
				{
					Type:        "Chest",
					Description: "some container",
					SupportedInteractions: []interaction.Type{
						"Open",
						"Loot",
					},
				},
			},
		}, nil))
		assert.NoError(t, err)
	})

	t.Run("should error on failed interact action", func(t *testing.T) {
		pm := &mock.PromptMock{}

		pm.On("Select", mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "some container", nil).Times(1)
		pm.On("Select", mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "Open", nil).Times(1)

		c := NewInteractCommand()
		err := c.Run(state.New(pm, nil, &state.RoomDefinition{
			Name: "some-room",
			Containers: []*state.Container{
				{
					Type:        "Chest",
					Description: "some container",
					SupportedInteractions: []interaction.Type{
						"Open",
						"Loot",
					},
				},
			},
		}, nil))
		assert.Error(t, err)
	})
}
