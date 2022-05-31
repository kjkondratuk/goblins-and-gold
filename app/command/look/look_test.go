package look

import (
	"github.com/kjkondratuk/goblins-and-gold/app/command"
	"github.com/kjkondratuk/goblins-and-gold/app/command/mock"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/container"
	"github.com/kjkondratuk/goblins-and-gold/world/room"
	"testing"
)

func Test_action(t *testing.T) {
	ctx := mock.MockContext{}
	ctx.On("State").Return(&state.State{
		Player: nil,
		CurrRoom: &room.Definition{
			Name:        "test-room",
			Description: "Some room description",
		},
		World:         nil,
		SelectBuilder: nil,
	})

	errorCtx := mock.MockContext{}
	errorCtx.On("State").Return(&state.State{
		Player:        nil,
		CurrRoom:      nil,
		World:         nil,
		SelectBuilder: nil,
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
			"should not error when looking with a valid context",
			args{c: &ctx},
			false,
		}, {
			"should not error when looking with an invalid context",
			args{c: &errorCtx},
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

	nilCurrentRoomCtx := mock.MockContext{}
	nilCurrentRoomCtx.On("State").Return(&state.State{})

	validCtx := mock.MockContext{}
	validCtx.On("State").Return(&state.State{
		CurrRoom: &room.Definition{
			Containers: []*container.Container{},
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
			"should be valid when a state and current room are non-nil",
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
