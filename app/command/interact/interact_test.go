package interact

import (
	"github.com/kjkondratuk/goblins-and-gold/app/command"
	"github.com/kjkondratuk/goblins-and-gold/app/command/mock"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/container"
	"github.com/kjkondratuk/goblins-and-gold/world/room"
	"testing"
)

func Test_action(t *testing.T) {
	noContainerCtx := mock.MockContext{}
	noContainerCtx.On("State").Return(&state.State{
		Player: nil,
		CurrRoom: &room.Definition{
			Name:                "test-room",
			Description:         "",
			Paths:               nil,
			Containers:          []*container.Container{},
			MandatoryEncounters: nil,
		},
		World:         nil,
		SelectBuilder: nil,
	})

	singleContainerCtx := mock.MockContext{}
	singleContainerCtx.On("State").Return(&state.State{
		CurrRoom: &room.Definition{
			Name: "test-room",
			Containers: []*container.Container{
				{
					Type:
				},
			},
		}
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
			"should complete successfully if there are no interactable items",
			args{&noContainerCtx},
			false,
		}, {
			"should complete successfully when cancelled",
			args{},
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
