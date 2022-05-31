package interact

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/app/command"
	"github.com/kjkondratuk/goblins-and-gold/app/command/mock"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	mock2 "github.com/kjkondratuk/goblins-and-gold/app/ux/mock"
	"github.com/kjkondratuk/goblins-and-gold/container"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/item"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	"github.com/kjkondratuk/goblins-and-gold/world/room"
	mock3 "github.com/stretchr/testify/mock"
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

	singleContainerState := state.State{
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
		CurrRoom: &room.Definition{
			Name: "test-room",
			Containers: []*container.Container{
				{
					Type:        container.Chest,
					Description: "A small wooden crate",
					SupportedInteractions: []interaction.Type{
						container.InteractionTypeOpen,
						container.InteractionTypeUnlock,
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

	interactionCancelSelector := mock2.SelectMock{}
	interactionCancelSelector.On("Create", mock3.AnythingOfType("string"), mock3.AnythingOfType("string")).Return(&interactionCancelSelector)
	interactionCancelSelector.On("Run", mock3.AnythingOfType("[]ux.Described")).Return(-1, "some value", nil)

	// copy state and assign select builder appropriate to a cancelled interactable select
	interactionErrorState := singleContainerState
	interactionErrorState.SelectBuilder = &interactionCancelSelector
	singleContainerCancelCtx.On("State").Return(&interactionErrorState)

	interactionErrorSelector := mock2.SelectMock{}
	interactionErrorSelector.On("Create", mock3.AnythingOfType("string"), mock3.AnythingOfType("string")).Return(&interactionErrorSelector)
	interactionErrorSelector.On("Run", mock3.AnythingOfType("[]ux.Described")).Return(-1, "", errors.New("there was an error"))

	// copy state and assign select builder appropriate to an errored interactable select
	actionErrorState := singleContainerState
	actionErrorState.SelectBuilder = &interactionErrorSelector
	singleContainerErrorCtx := mock.MockContext{}
	singleContainerErrorCtx.On("State").Return(&actionErrorState)

	actionErrorSelector := mock2.SelectMock{}
	actionErrorSelector.On("Create", mock3.AnythingOfType("string"), mock3.AnythingOfType("string")).Return(&actionErrorSelector).Twice()
	actionErrorSelector.On("Run", mock3.AnythingOfType("[]ux.Described")).Return(0, "", nil).Once()
	actionErrorSelector.On("Run", mock3.AnythingOfType("[]ux.Described")).Return(-1, "", errors.New("there was an error")).Once()

	// copy state and assign select builder appropriate to an errored action select
	errorState := singleContainerState
	errorState.SelectBuilder = &actionErrorSelector
	actionErrorCtx := mock.MockContext{}
	actionErrorCtx.On("State").Return(&errorState)

	actionCancelSelector := mock2.SelectMock{}
	actionCancelSelector.On("Create", mock3.AnythingOfType("string"), mock3.AnythingOfType("string")).
		Return(&actionCancelSelector).Twice()
	actionCancelSelector.On("Run", mock3.AnythingOfType("[]ux.Described")).Return(0, "", nil).
		Once()
	actionCancelSelector.On("Run", mock3.AnythingOfType("[]ux.Described")).Return(-1, "", nil).
		Once()

	// copy state and assign select builder appropriate to a cancelled action select
	cancelState := singleContainerState
	cancelState.SelectBuilder = &actionCancelSelector
	actionCancelCtx := mock.MockContext{}
	actionCancelCtx.On("State").Return(&cancelState)

	actionCompleteSelector := mock2.SelectMock{}
	actionCompleteSelector.On("Create", mock3.AnythingOfType("string"), mock3.AnythingOfType("string")).
		Return(&actionCompleteSelector).Twice()
	actionCompleteSelector.On("Run", mock3.AnythingOfType("[]ux.Described")).Return(0, "", nil).
		Once()
	actionCompleteSelector.On("Run", mock3.AnythingOfType("[]ux.Described")).Return(0, "Open", nil).
		Once()

	completeState := singleContainerState
	completeState.SelectBuilder = &actionCompleteSelector
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

	nilContainerCtx := mock.MockContext{}
	nilContainerCtx.On("State").Return(&state.State{
		CurrRoom: &room.Definition{
			Containers: nil,
		},
	})

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
