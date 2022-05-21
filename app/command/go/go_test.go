package _go

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	mock3 "github.com/kjkondratuk/goblins-and-gold/app/command/mock"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/app/ux/mock"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	"github.com/kjkondratuk/goblins-and-gold/world"
	"github.com/kjkondratuk/goblins-and-gold/world/path"
	"github.com/kjkondratuk/goblins-and-gold/world/room"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"testing"
)

func Test_action(t *testing.T) {
	t.Run("should not fail when there is nowhere to go", func(t *testing.T) {
		ctx := mock3.MockContext{}
		ctx.On("State").Return(&state.State{
			CurrRoom: &room.Definition{
				Paths: []*path.Definition{},
			},
		})
		err := action(&ctx)
		assert.NoError(t, err)
	})

	t.Run("should fail if there was an error building the select menu", func(t *testing.T) {
		ctx := mock3.MockContext{}
		sm := mock.SelectMock{}
		sm.On("Create", mock2.AnythingOfType("string"), mock2.AnythingOfType("string")).
			Return(&sm)
		sm.On("Run", mock2.AnythingOfType("[]ux.Described")).
			Return(0, "", errors.New("something went wrong"))
		ctx.On("State").Return(&state.State{
			Player: nil,
			CurrRoom: &room.Definition{
				Paths: []*path.Definition{
					{
						"another-room", "a creaky wooden door",
					},
				},
			},
			World:         nil,
			SelectBuilder: &sm,
		})

		err := action(&ctx)
		assert.Error(t, err)
	})

	t.Run("should complete successfully if the index corresponds to cancellation", func(t *testing.T) {
		ctx := mock3.MockContext{}
		sm := mock.SelectMock{}
		sm.On("Create", mock2.AnythingOfType("string"), mock2.AnythingOfType("string")).
			Return(&sm)
		sm.On("Run", mock2.AnythingOfType("[]ux.Described")).
			Return(-1, "", nil)
		ctx.On("State").Return(&state.State{
			Player: nil,
			CurrRoom: &room.Definition{
				Paths: []*path.Definition{
					{
						"another-room", "a creaky wooden door",
					},
				},
			},
			World:         nil,
			SelectBuilder: &sm,
		})

		err := action(&ctx)
		assert.NoError(t, err)
	})

	t.Run("should invoke quit command when a player is unconscious by encounter conclusion", func(t *testing.T) {
		ctx := mock3.MockContext{}
		sm := mock.SelectMock{}
		sm.On("Create", mock2.AnythingOfType("string"), mock2.AnythingOfType("string")).
			Return(&sm)
		sm.On("Run", mock2.AnythingOfType("[]ux.Described")).
			Return(0, "", nil)
		ctx.On("RunCommandByName", mock2.AnythingOfType("string")).Return(nil).Once()
		ctx.On("State").Return(&state.State{
			Player: actors.NewPlayer(actors.PlayerParams{CombatantParams: actors.CombatantParams{
				Name:      "",
				AC:        0,
				HP:        0,
				BaseStats: stats.BaseStats{},
				Inventory: nil,
				Attacks:   nil,
			}}),
			CurrRoom: &room.Definition{
				Paths: []*path.Definition{
					{
						"another-room", "a creaky wooden door",
					},
				},
			},
			World: &world.Definition{
				Rooms: map[string]*room.Definition{
					"another-room": {
						Name: "another-room",
					},
				},
			},
			SelectBuilder: &sm,
		}).Once()

		err := action(&ctx)
		assert.NoError(t, err)
	})

	t.Run("should NOT invoke quit command when a player ends with HP", func(t *testing.T) {
		ctx := mock3.MockContext{}
		sm := mock.SelectMock{}
		sm.On("Create", mock2.AnythingOfType("string"), mock2.AnythingOfType("string")).
			Return(&sm)
		sm.On("Run", mock2.AnythingOfType("[]ux.Described")).
			Return(0, "", nil)
		ctx.On("State").Return(&state.State{
			Player: actors.NewPlayer(actors.PlayerParams{CombatantParams: actors.CombatantParams{
				Name:      "",
				AC:        0,
				HP:        10,
				BaseStats: stats.BaseStats{},
				Inventory: nil,
				Attacks:   nil,
			}}),
			CurrRoom: &room.Definition{
				Paths: []*path.Definition{
					{
						"another-room", "a creaky wooden door",
					},
				},
			},
			World: &world.Definition{
				Rooms: map[string]*room.Definition{
					"another-room": {
						Name: "another-room",
					},
				},
			},
			SelectBuilder: &sm,
		}).Times(10)

		err := action(&ctx)
		assert.NoError(t, err)
	})
}
