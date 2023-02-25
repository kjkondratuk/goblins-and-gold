package stats

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	"testing"
)

func Test_action(t *testing.T) {
	var nilState state.State

	nilPlayer := state.New(nil, nil, &state.RoomDefinition{}, &state.WorldDefinition{})

	nnPlayer := state.New(
		nil,
		actors.NewPlayer(actors.PlayerParams{CombatantParams: actors.CombatantParams{
			Name:      "test combatant",
			AC:        0,
			HP:        0,
			BaseStats: stats.BaseStats{},
			Inventory: nil,
			Attacks:   nil,
		}}),
		nil,
		nil,
	)

	type args struct {
		s state.State
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"should not error if state is nil",
			args{nilState},
			false,
		}, {
			"should not error when player is nil",
			args{nilPlayer},
			false,
		}, {
			"should not error when player is non-nil",
			args{nnPlayer},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := action(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("action() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateContext(t *testing.T) {
	var nilState state.State

	nilPlayerState := state.New(nil, nil, &state.RoomDefinition{}, &state.WorldDefinition{})

	validState := state.New(nil, actors.NewPlayer(actors.PlayerParams{}),
		&state.RoomDefinition{}, &state.WorldDefinition{})

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
			"should be invalid when player is nil",
			args{nilPlayerState},
			true,
		}, {
			"should be valid when a state and player are non-nil",
			args{validState},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateState(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("validateContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
