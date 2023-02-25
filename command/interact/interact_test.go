package interact

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/item"
	"github.com/kjkondratuk/goblins-and-gold/state"
	state2 "github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	mock3 "github.com/kjkondratuk/goblins-and-gold/ux/mock"
	mock2 "github.com/stretchr/testify/mock"
	"testing"
)

func Test_action(t *testing.T) {
	noContainerState := state.New(nil, nil, &state.RoomDefinition{
		Name:                "test-room",
		Description:         "",
		Paths:               nil,
		Containers:          []*state2.Container{},
		MandatoryEncounters: nil,
	}, nil)

	singleContainerState := state.New(nil, actors.NewPlayer(actors.PlayerParams{
		CombatantParams: actors.CombatantParams{
			Name:      "Test Player",
			AC:        15,
			HP:        10,
			BaseStats: stats.BaseStats{},
			Inventory: []item.Item{},
			Attacks:   nil,
		},
	}), &state.RoomDefinition{
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
	}, nil)

	// copy state and assign select builder appropriate to a cancelled interactable select
	interactionCancelState := singleContainerState
	cancelInteractPromptMock := &mock3.PromptMock{}
	cancelInteractPromptMock.On("Select",
		mock2.AnythingOfType("string"),
		mock2.AnythingOfType("[]string")).
		Return(-1, "", nil).Times(1)
	interactionCancelState.SetPrompter(cancelInteractPromptMock)

	// copy state and assign select builder appropriate to an errored interactable select
	interactErrorState := singleContainerState
	errorInteractPromptMock := &mock3.PromptMock{}
	errorInteractPromptMock.On("Select",
		mock2.AnythingOfType("string"),
		mock2.AnythingOfType("[]string")).
		Return(0, "", errors.New("something bad happened")).Times(1)
	interactErrorState.SetPrompter(errorInteractPromptMock)

	// copy state and assign select builder appropriate to an errored action select
	errorActionState := singleContainerState
	errorActionPromptMock := &mock3.PromptMock{}
	errorActionPromptMock.On("Select",
		mock2.AnythingOfType("string"),
		mock2.AnythingOfType("[]string")).
		Return(0, "", nil).Times(1).
		Return(0, "", errors.New("something bad happened")).Times(1)
	errorActionState.SetPrompter(errorActionPromptMock)

	// copy state and assign select builder appropriate to a cancelled action select
	cancelActionState := singleContainerState
	cancelActionPromptMock := &mock3.PromptMock{}
	cancelActionPromptMock.On("Select",
		mock2.AnythingOfType("string"),
		mock2.AnythingOfType("[]string")).
		Return(0, "", nil).Times(1).
		Return(-1, "", nil).Times(1)
	cancelActionState.SetPrompter(cancelActionPromptMock)

	completeState := singleContainerState
	completeActionPromptMock := &mock3.PromptMock{}
	completeActionPromptMock.On("Select",
		mock2.AnythingOfType("string"),
		mock2.AnythingOfType("[]string")).
		Return(0, "", nil).Times(1).
		Return(0, "", nil).Times(1)
	completeState.SetPrompter(completeActionPromptMock)

	type args struct {
		s state.State
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"should complete successfully if there are no interactable items",
			args{noContainerState},
			false,
		}, {
			"should complete successfully when interactable selection cancelled",
			args{interactionCancelState},
			false,
		}, {
			"should fail when there is an error in the action selector",
			args{errorActionState},
			true,
		}, {
			"should do nothing when the action selection is cancelled",
			args{cancelActionState},
			false,
		}, { // TODO: finish this test
			"should perform interaction when one is selected",
			args{completeState},
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

	nilCurrentRoomState := state.New(nil, nil, nil, nil)

	nilContainerState := state.New(nil, nil, &state.RoomDefinition{
		Containers: nil,
	}, nil)

	validState := state.New(nil, nil, &state.RoomDefinition{
		Containers: []*state2.Container{},
	}, nil)

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
			args{nilCurrentRoomState},
			true,
		}, {
			"should be invalid when containers in the current room are nil",
			args{nilContainerState},
			true,
		}, {
			"should be valid when a state room and containers are non-nil",
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
