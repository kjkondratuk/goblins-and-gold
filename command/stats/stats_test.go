package stats

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/command"
	"github.com/kjkondratuk/goblins-and-gold/command/mock"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	"testing"
)

func Test_action(t *testing.T) {
	nilCtx := mock.MockContext{}
	nilCtx.On("State").Return(nil)

	nilPlayer := mock.MockContext{}
	nilPlayer.On("State").Return(&state.State{
		Player:   nil,
		CurrRoom: nil,
		World:    nil,
	})

	nnPlayer := mock.MockContext{}
	nnPlayer.On("State").Return(&state.State{
		Player: actors.NewPlayer(actors.PlayerParams{CombatantParams: actors.CombatantParams{
			Name:      "test combatant",
			AC:        0,
			HP:        0,
			BaseStats: stats.BaseStats{},
			Inventory: nil,
			Attacks:   nil,
		}}),
		CurrRoom: nil,
		World:    nil,
	})

	type args struct {
		c command.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"should not error if state is nil",
			args{&nilCtx},
			false,
		}, {
			"should not error when player is nil",
			args{&nilPlayer},
			false,
		}, {
			"should not error when player is non-nil",
			args{&nnPlayer},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := action(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("action() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_validateContext(t *testing.T) {
	nilStateContext := mock.MockContext{}
	nilStateContext.On("State").Return(nil)

	nilPlayerCtx := mock.MockContext{}
	nilPlayerCtx.On("State").Return(&state.State{})

	validCtx := mock.MockContext{}
	validCtx.On("State").Return(&state.State{
		Player: actors.NewPlayer(actors.PlayerParams{CombatantParams: actors.CombatantParams{
			Name:      "Test Player",
			AC:        15,
			HP:        10,
			BaseStats: stats.BaseStats{},
		}}),
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
			"should be invalid when player is nil",
			args{&nilPlayerCtx},
			true,
		}, {
			"should be valid when a state and player are non-nil",
			args{&validCtx},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateContext(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("validateContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
