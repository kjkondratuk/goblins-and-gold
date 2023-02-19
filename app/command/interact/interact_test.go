package interact

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/app/command"
	"github.com/kjkondratuk/goblins-and-gold/app/command/mock"
	mock3 "github.com/kjkondratuk/goblins-and-gold/app/ux/mock"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/item"
	"github.com/kjkondratuk/goblins-and-gold/state"
	state2 "github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	mock2 "github.com/stretchr/testify/mock"
	"testing"
)

func Test_action(t *testing.T) {
	noContainerCtx := mock.MockContext{}
	noContainerCtx.On("State").Return(&state.State{
		Player: nil,
		CurrRoom: &state.RoomDefinition{
			Name:                "test-room",
			Description:         "",
			Paths:               nil,
			Containers:          []*state2.Container{},
			MandatoryEncounters: nil,
		},
		World: nil,
	})

	promptMock := &mock3.PromptMock{}
	promptMock.On("Select",
		mock2.AnythingOfType("string"),
		mock2.AnythingOfType("[]string")).
		Return(0, "", nil)

	singleContainerState := state.State{
		PromptLib: promptMock,
		Player: actors.NewPlayer(actors.PlayerParams{
			CombatantParams: actors.CombatantParams{
				Name:      "Test Player",
				AC:        15,
				HP:        10,
				BaseStats: stats.BaseStats{},
				Inventory: []item.Item{},
				Attacks:   nil,
			},
		}),
		CurrRoom: &state.RoomDefinition{
			Name: "test-room",
			Containers: []*state2.Container{
				{
					Type:        state2.Chest,
					Description: "A small wooden crate",
					SupportedInteractions: []interaction.Type{
						state2.InteractionTypeOpen,
						state2.InteractionTypeUnlock,
					},
					Items: []item.Item{
						{
							Type:        "Gold",
							Description: "A single gold coin",
							Quantity:    1,
							Unit:        "coins",
						},
					},
				},
			},
		},
	}

	singleContainerCancelCtx := mock.MockContext{}

	// copy state and assign select builder appropriate to a cancelled interactable select
	interactionErrorState := singleContainerState
	singleContainerCancelCtx.On("State").Return(&interactionErrorState)

	// copy state and assign select builder appropriate to an errored interactable select
	actionErrorState := singleContainerState
	errorInteractPromptMock := &mock3.PromptMock{}
	errorInteractPromptMock.On("Select",
		mock2.AnythingOfType("string"),
		mock2.AnythingOfType("[]string")).
		Return(0, "", errors.New("something bad happened"))
	actionErrorState.PromptLib = errorInteractPromptMock

	singleContainerErrorCtx := mock.MockContext{}
	singleContainerErrorCtx.On("State").Return(&actionErrorState)

	// copy state and assign select builder appropriate to an errored action select
	errorState := singleContainerState
	errorActionPromptMock := &mock3.PromptMock{}
	errorActionPromptMock.On("Select",
		mock2.AnythingOfType("string"),
		mock2.AnythingOfType("[]string")).
		Return(0, "", nil).
		Return(0, "", errors.New("something bad happened"))
	errorState.PromptLib = errorActionPromptMock
	actionErrorCtx := mock.MockContext{}
	actionErrorCtx.On("State").Return(&errorState)

	// copy state and assign select builder appropriate to a cancelled action select
	cancelState := singleContainerState
	actionCancelCtx := mock.MockContext{}
	actionCancelCtx.On("State").Return(&cancelState)

	completeState := singleContainerState
	actionCompleteCtx := mock.MockContext{}
	actionCompleteCtx.On("State").Return(&completeState)

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
			"should complete successfully when interactable selection cancelled",
			args{&singleContainerCancelCtx},
			false,
		}, {
			"should error when the interactable selector errors",
			args{&singleContainerErrorCtx},
			true,
		}, {
			"should fail when there is an error in the action selector",
			args{&actionErrorCtx},
			true,
		}, {
			"should do nothing when the action selection is cancelled",
			args{&actionCancelCtx},
			false,
		}, { // TODO: finish this test
			"should perform interaction when one is selected",
			args{&actionCompleteCtx},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := action(tt.args.c.State())(tt.args.c); (err != nil) != tt.wantErr {
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

	nilContainerCtx := mock.MockContext{}
	nilContainerCtx.On("State").Return(&state.State{
		CurrRoom: &state.RoomDefinition{
			Containers: nil,
		},
	})

	validCtx := mock.MockContext{}
	validCtx.On("State").Return(&state.State{
		CurrRoom: &state.RoomDefinition{
			Containers: []*state2.Container{},
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
			"should be invalid when containers in the current room are nil",
			args{&nilContainerCtx},
			true,
		}, {
			"should be valid when a state room and containers are non-nil",
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
