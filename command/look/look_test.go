package look

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"testing"
)

func Test_action(t *testing.T) {
	s := state.New(
		nil,
		nil,
		&state.RoomDefinition{
			Name:        "test-room",
			Description: "Some room description",
		},
		nil,
	)

	bs := state.New(
		nil,
		nil,
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
			"should not error when looking with a valid context",
			args{s: s},
			false,
		}, {
			"should not error when looking with an invalid context",
			args{s: bs},
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
	var nilContext state.State

	nilRoomState := state.New(nil, actors.NewPlayer(actors.PlayerParams{}), nil, &state.WorldDefinition{})

	validState := state.New(nil, actors.NewPlayer(actors.PlayerParams{}), &state.RoomDefinition{}, &state.WorldDefinition{})

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
			args{nilContext},
			true,
		}, {
			"should be invalid when current room is nil",
			args{nilRoomState},
			true,
		}, {
			"should be valid when a state and current room are non-nil",
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
