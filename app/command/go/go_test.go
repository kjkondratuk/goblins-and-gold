package _go

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/app/command"
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

func Test_validateContext(t *testing.T) {
	nilStateContext := mock3.MockContext{}
	nilStateContext.On("State").Return(nil)

	nilCurrentRoomCtx := mock3.MockContext{}
	nilCurrentRoomCtx.On("State").Return(&state.State{})

	nilPathCtx := mock3.MockContext{}
	nilPathCtx.On("State").Return(&state.State{
		CurrRoom: &room.Definition{
			Paths: nil,
		},
	})

	validCtx := mock3.MockContext{}
	validCtx.On("State").Return(&state.State{
		CurrRoom: &room.Definition{
			Paths: []*path.Definition{
				{
					"test-room",
					"plain wooden door",
				},
			},
		},
	})

	type args struct {
		ctx command.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"should be invalid when state is nil",
			args{&nilStateContext},
			true,
		}, {
			"should be invalid when current room is nil",
			args{&nilCurrentRoomCtx},
			true,
		}, {
			"should be invalid when paths in the current room are nil",
			args{&nilPathCtx},
			true,
		}, {
			"should be valid when a state room and paths are non-nil",
			args{&validCtx},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateContext(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("action() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
