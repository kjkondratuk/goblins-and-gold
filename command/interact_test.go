package command

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/container"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	container2 "github.com/kjkondratuk/goblins-and-gold/model/container"
	"github.com/kjkondratuk/goblins-and-gold/model/room"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"testing"
)

func TestNewInteractCommand(t *testing.T) {
	mc := &container.MockContainerController{}
	mc.EXPECT().Do(
		mock2.AnythingOfType("state.State"),
		mock2.AnythingOfType("container.Container"),
		mock2.AnythingOfType("interaction.Type")).Times(0)

	ma := &state.MockApplier[interaction.Result]{}
	ma.EXPECT().Apply(mock2.Anything, mock2.Anything).Times(0)

	c := NewInteractCommand(mc, ma)
	assert.NotNil(t, c)
	assert.Equal(t, c.Name(), "interact")
	assert.Len(t, c.Subcommands(), 1)
}

func TestInteractCommand_Run(t *testing.T) {
	t.Run("should invoke help when args are supplied", func(t *testing.T) {
		mc := &container.MockContainerController{}
		mc.EXPECT().Do(
			mock2.AnythingOfType("state.State"),
			mock2.AnythingOfType("container.Container"),
			mock2.AnythingOfType("interaction.Type")).Times(0)

		ma := &state.MockApplier[interaction.Result]{}
		ma.EXPECT().Apply(mock2.Anything, mock2.Anything).Times(0)

		c := NewInteractCommand(mc, ma)
		err := c.Run(nil, "some arg")
		assert.NoError(t, err)
	})

	t.Run("should handle nil containers to interact with", func(t *testing.T) {
		mc := &container.MockContainerController{}
		mc.EXPECT().Do(
			mock2.AnythingOfType("state.State"),
			mock2.AnythingOfType("container.Container"),
			mock2.AnythingOfType("interaction.Type")).Times(0)

		ma := &state.MockApplier[interaction.Result]{}
		ma.EXPECT().Apply(mock2.Anything, mock2.Anything).Times(0)

		c := NewInteractCommand(mc, ma)
		err := c.Run(state.New(nil, nil, &room.RoomDefinition{
			Name: "some-room",
		}, nil))
		assert.NoError(t, err)
	})

	t.Run("should handle no containers to interact with", func(t *testing.T) {
		mc := &container.MockContainerController{}
		mc.EXPECT().Do(
			mock2.AnythingOfType("state.State"),
			mock2.AnythingOfType("container.Container"),
			mock2.AnythingOfType("interaction.Type")).Times(0)

		ma := &state.MockApplier[interaction.Result]{}
		ma.EXPECT().Apply(mock2.Anything, mock2.Anything).Times(0)

		c := NewInteractCommand(mc, ma)
		err := c.Run(state.New(nil, nil, &room.RoomDefinition{
			Name:       "some-room",
			Containers: []*container2.Container{},
		}, nil))
		assert.NoError(t, err)
	})

	t.Run("should error for interaction select error", func(t *testing.T) {
		pm := &ux.MockPromptLib{}
		pm.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(-1, "", errors.New("an error occurred"))

		mc := &container.MockContainerController{}
		mc.EXPECT().Do(
			mock2.AnythingOfType("state.State"),
			mock2.AnythingOfType("container.Container"),
			mock2.AnythingOfType("interaction.Type")).Times(0)

		ma := &state.MockApplier[interaction.Result]{}
		ma.EXPECT().Apply(mock2.Anything, mock2.Anything).Times(0)

		c := NewInteractCommand(mc, ma)
		err := c.Run(state.New(pm, nil, &room.RoomDefinition{
			Name: "some-room",
			Containers: []*container2.Container{
				{
					Type:        "Chest",
					Description: "some container",
				},
			},
		}, nil))
		assert.Error(t, err)
	})

	t.Run("should not error for interaction select cancel", func(t *testing.T) {
		pm := &ux.MockPromptLib{}
		pm.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(0, "", nil)

		mc := &container.MockContainerController{}
		mc.EXPECT().Do(
			mock2.AnythingOfType("state.State"),
			mock2.AnythingOfType("container.Container"),
			mock2.AnythingOfType("interaction.Type")).Times(0)

		ma := &state.MockApplier[interaction.Result]{}
		ma.EXPECT().Apply(mock2.Anything, mock2.Anything).Times(0)

		c := NewInteractCommand(mc, ma)
		err := c.Run(state.New(pm, nil, &room.RoomDefinition{
			Name: "some-room",
			Containers: []*container2.Container{
				{
					Type:        "Chest",
					Description: "some container",
				},
			},
		}, nil))
		assert.NoError(t, err)
	})

	t.Run("should error on action select failure", func(t *testing.T) {
		pm := &ux.MockPromptLib{}
		pm.EXPECT().
			Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "some container", nil).Times(1)
		pm.EXPECT().
			Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(-1, "", errors.New("something went wrong")).Times(1)

		mc := &container.MockContainerController{}
		mc.EXPECT().Do(
			mock2.AnythingOfType("state.State"),
			mock2.AnythingOfType("container.Container"),
			mock2.AnythingOfType("interaction.Type")).Times(0)

		ma := &state.MockApplier[interaction.Result]{}
		ma.EXPECT().Apply(mock2.Anything, mock2.Anything).Times(0)

		c := NewInteractCommand(mc, ma)
		err := c.Run(state.New(pm, nil, &room.RoomDefinition{
			Name: "some-room",
			Containers: []*container2.Container{
				{
					Type:        "Chest",
					Description: "some container",
					SupportedInteractions: []container2.Type{
						"Open",
						"Loot",
					},
				},
			},
		}, nil))
		assert.Error(t, err)
	})

	t.Run("should succeed on action select cancel", func(t *testing.T) {
		pm := &ux.MockPromptLib{}
		pm.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "some container", nil).Times(1)
		pm.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(0, "", nil).Times(1)

		mc := &container.MockContainerController{}
		mc.EXPECT().Do(
			mock2.AnythingOfType("state.State"),
			mock2.AnythingOfType("container.Container"),
			mock2.AnythingOfType("interaction.Type")).Times(0)

		ma := &state.MockApplier[interaction.Result]{}
		ma.EXPECT().Apply(mock2.Anything, mock2.Anything).Times(0)

		c := NewInteractCommand(mc, ma)
		err := c.Run(state.New(pm, nil, &room.RoomDefinition{
			Name: "some-room",
			Containers: []*container2.Container{
				{
					Type:        "Chest",
					Description: "some container",
					SupportedInteractions: []container2.Type{
						"Open",
						"Loot",
					},
				},
			},
		}, nil))
		assert.NoError(t, err)
	})

	t.Run("should error on failed interact action", func(t *testing.T) {
		pm := &ux.MockPromptLib{}
		pm.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "some container", nil)
		pm.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "Open", nil)

		mc := &container.MockContainerController{}
		mc.EXPECT().Do(
			mock2.Anything,
			mock2.Anything,
			mock2.Anything).
			Return(interaction.Result{}, errors.New("something bad happened")).
			Times(1)

		ma := &state.MockApplier[interaction.Result]{}
		ma.EXPECT().Apply(mock2.Anything, mock2.Anything).Times(0)

		c := NewInteractCommand(mc, ma)
		err := c.Run(state.New(pm, nil, &room.RoomDefinition{
			Name: "some-room",
			Containers: []*container2.Container{
				{
					Type:        "Chest",
					Description: "some container",
					SupportedInteractions: []container2.Type{
						"Open",
						"Loot",
					},
				},
			},
		}, nil))
		assert.Error(t, err)
	})

	t.Run("should not error on non-applicable result", func(t *testing.T) {
		pm := &ux.MockPromptLib{}
		pm.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "some container", nil)
		pm.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "Open", nil)

		mc := &container.MockContainerController{}
		mc.EXPECT().Do(
			mock2.Anything,
			mock2.Anything,
			mock2.Anything).
			Return(interaction.Result{}, nil).
			Times(1)

		ma := &state.MockApplier[interaction.Result]{}
		ma.EXPECT().Apply(mock2.Anything, mock2.Anything).Times(0)

		c := NewInteractCommand(mc, ma)
		err := c.Run(state.New(pm, nil, &room.RoomDefinition{
			Name: "some-room",
			Containers: []*container2.Container{
				{
					Type:        "Chest",
					Description: "some container",
					SupportedInteractions: []container2.Type{
						"Open",
						"Loot",
					},
				},
			},
		}, nil))
		assert.NoError(t, err)
	})

	t.Run("should not error on success result", func(t *testing.T) {
		pm := &ux.MockPromptLib{}
		pm.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "some container", nil)
		pm.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "Open", nil)

		mc := &container.MockContainerController{}
		mc.EXPECT().Do(
			mock2.Anything,
			mock2.Anything,
			mock2.Anything).
			Return(interaction.Result{Type: interaction.RtSuccess}, nil).
			Times(1)

		ma := &state.MockApplier[interaction.Result]{}
		ma.EXPECT().Apply(mock2.Anything, mock2.Anything).Times(1)

		c := NewInteractCommand(mc, ma)
		err := c.Run(state.New(pm, nil, &room.RoomDefinition{
			Name: "some-room",
			Containers: []*container2.Container{
				{
					Type:        "Chest",
					Description: "some container",
					SupportedInteractions: []container2.Type{
						"Open",
						"Loot",
					},
				},
			},
		}, nil))
		assert.NoError(t, err)
	})

	t.Run("should not error on failure result", func(t *testing.T) {
		pm := &ux.MockPromptLib{}
		pm.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "some container", nil)
		pm.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "Open", nil)

		mc := &container.MockContainerController{}
		mc.EXPECT().Do(
			mock2.Anything,
			mock2.Anything,
			mock2.Anything).
			Return(interaction.Result{Type: interaction.RtFailure}, nil).
			Times(1)

		ma := &state.MockApplier[interaction.Result]{}
		ma.EXPECT().Apply(mock2.Anything, mock2.Anything).Times(1)

		c := NewInteractCommand(mc, ma)
		err := c.Run(state.New(pm, nil, &room.RoomDefinition{
			Name: "some-room",
			Containers: []*container2.Container{
				{
					Type:        "Chest",
					Description: "some container",
					SupportedInteractions: []container2.Type{
						"Open",
						"Loot",
					},
				},
			},
		}, nil))
		assert.NoError(t, err)
	})

	t.Run("should not error on other result type", func(t *testing.T) {
		pm := &ux.MockPromptLib{}
		pm.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "some container", nil)
		pm.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
			Return(1, "Open", nil)

		mc := &container.MockContainerController{}
		mc.EXPECT().Do(
			mock2.Anything,
			mock2.Anything,
			mock2.Anything).
			Return(interaction.Result{Type: "other", Message: "some message"}, nil).
			Times(1)

		ma := &state.MockApplier[interaction.Result]{}
		ma.EXPECT().Apply(mock2.Anything, mock2.Anything).Times(1)

		c := NewInteractCommand(mc, ma)
		err := c.Run(state.New(pm, nil, &room.RoomDefinition{
			Name: "some-room",
			Containers: []*container2.Container{
				{
					Type:        "Chest",
					Description: "some container",
					SupportedInteractions: []container2.Type{
						"Open",
						"Loot",
					},
				},
			},
		}, nil))
		assert.NoError(t, err)
	})
}
