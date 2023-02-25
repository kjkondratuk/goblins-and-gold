package _go

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	"github.com/kjkondratuk/goblins-and-gold/ux/mock"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"testing"
)

type mockPromptLib struct {
}

func Test_action(t *testing.T) {
	t.Run("should not fail when there is nowhere to go", func(t *testing.T) {
		s := state.New(
			nil,
			nil,
			&state.RoomDefinition{
				Paths: []*state.PathDefinition{},
			},
			nil,
		)

		err := action(s)
		assert.NoError(t, err)
	})

	t.Run("should fail if there was an error building the select menu", func(t *testing.T) {
		promptMock := &mock.PromptMock{}
		promptMock.On("Select",
			mock2.AnythingOfType("string"),
			mock2.AnythingOfType("[]string")).
			Return(0, "", errors.New("something went wrong"))
		s := state.New(
			promptMock,
			nil,
			&state.RoomDefinition{
				Paths: []*state.PathDefinition{
					{
						"another-room", "a creaky wooden door",
					},
				},
			},
			nil,
		)

		err := action(s)
		assert.Error(t, err)
	})

	t.Run("should complete successfully if the index corresponds to cancellation", func(t *testing.T) {
		promptMock := &mock.PromptMock{}
		promptMock.On("Select",
			mock2.AnythingOfType("string"),
			mock2.AnythingOfType("[]string")).
			Return(-1, "", nil)
		s := state.New(
			promptMock,
			nil,
			&state.RoomDefinition{
				Paths: []*state.PathDefinition{
					{
						"another-room", "a creaky wooden door",
					},
				},
			},
			nil,
		)

		err := action(s)
		assert.NoError(t, err)
	})

	t.Run("should invoke quit command when a player is unconscious by encounter conclusion", func(t *testing.T) {
		promptMock := &mock.PromptMock{}
		promptMock.On("Select",
			mock2.AnythingOfType("string"),
			mock2.AnythingOfType("[]string")).
			Return(0, "", nil)
		s := state.New(
			promptMock,
			actors.NewPlayer(actors.PlayerParams{CombatantParams: actors.CombatantParams{
				Name:      "",
				AC:        0,
				HP:        0,
				BaseStats: stats.BaseStats{},
				Inventory: nil,
				Attacks:   nil,
			}}),
			&state.RoomDefinition{
				Paths: []*state.PathDefinition{
					{
						"another-room", "a creaky wooden door",
					},
				},
			},
			&state.WorldDefinition{
				Rooms: map[string]*state.RoomDefinition{
					"another-room": {
						Name: "another-room",
					},
				},
			},
		)

		err := action(s)
		assert.NoError(t, err)
	})

	t.Run("should NOT invoke quit command when a player ends with HP", func(t *testing.T) {
		promptMock := &mock.PromptMock{}
		promptMock.On("Select",
			mock2.AnythingOfType("string"),
			mock2.AnythingOfType("[]string")).
			Return(0, "", nil)
		s := state.New(
			promptMock,
			actors.NewPlayer(actors.PlayerParams{CombatantParams: actors.CombatantParams{
				Name:      "",
				AC:        0,
				HP:        10,
				BaseStats: stats.BaseStats{},
				Inventory: nil,
				Attacks:   nil,
			}}),
			&state.RoomDefinition{
				Paths: []*state.PathDefinition{
					{
						"another-room", "a creaky wooden door",
					},
				},
			},
			&state.WorldDefinition{
				Rooms: map[string]*state.RoomDefinition{
					"another-room": {
						Name: "another-room",
					},
				},
			},
		)

		err := action(s)
		assert.NoError(t, err)
	})
}

func Test_validateState(t *testing.T) {
	var nilState state.State

	emptyState := state.New(nil, nil, nil, nil)

	nilRoomState := state.New(nil, actors.NewPlayer(actors.PlayerParams{}),
		&state.RoomDefinition{Paths: []*state.PathDefinition{}}, &state.WorldDefinition{})

	nilPathState := state.New(nil, actors.NewPlayer(actors.PlayerParams{}),
		&state.RoomDefinition{Paths: nil}, &state.WorldDefinition{})

	validState := state.New(nil, actors.NewPlayer(actors.PlayerParams{}),
		&state.RoomDefinition{Paths: []*state.PathDefinition{}}, &state.WorldDefinition{})

	type args struct {
		s state.State
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"should be invalid when state is nil",
			args{nilState},
			true,
		}, {
			"should be invalid when current room is nil",
			args{nilRoomState},
			true,
		}, {
			"should be invalid when paths in the current room are nil",
			args{nilPathState},
			true,
		}, {
			"should be valid when a state room and paths are non-nil",
			args{validState},
			false,
		}, {
			"should be invalid when state is empty",
			args{emptyState},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateState(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("action() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
