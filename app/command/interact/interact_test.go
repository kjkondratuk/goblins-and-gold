package interact

import (
	"errors"
	"github.com/kjkondratuk/goblins-and-gold/app/command"
	"github.com/kjkondratuk/goblins-and-gold/app/command/mock"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	mock2 "github.com/kjkondratuk/goblins-and-gold/app/ux/mock"
	"github.com/kjkondratuk/goblins-and-gold/container"
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/item"
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
	sccs := singleContainerState
	sccs.SelectBuilder = &interactionCancelSelector
	singleContainerCancelCtx.On("State").Return(&sccs)

	interactionErrorSelector := mock2.SelectMock{}
	interactionErrorSelector.On("Create", mock3.AnythingOfType("string"), mock3.AnythingOfType("string")).Return(&interactionErrorSelector)
	interactionErrorSelector.On("Run", mock3.AnythingOfType("[]ux.Described")).Return(-1, "", errors.New("there was an error"))

	// copy state and assign select builder appropriate to an errored interactable select
	sces := singleContainerState
	sces.SelectBuilder = &interactionErrorSelector
	singleContainerErrorCtx := mock.MockContext{}
	singleContainerErrorCtx.On("State").Return(&sces)

	actionErrorSelector := mock2.SelectMock{}
	actionErrorSelector.On("Create", mock3.AnythingOfType("string"), mock3.AnythingOfType("string")).Return(&actionErrorSelector).Twice()
	actionErrorSelector.On("Run", mock3.AnythingOfType("[]ux.Described")).Return(0, "", nil).Once()
	actionErrorSelector.On("Run", mock3.AnythingOfType("[]ux.Described")).Return(-1, "", errors.New("there was an error")).Once()

	// copy state and assign select builder appropriate to an errored action select
	scesa := singleContainerState
	scesa.SelectBuilder = &actionErrorSelector
	actionErrorCtx := mock.MockContext{}
	actionErrorCtx.On("State").Return(&scesa)

	actionCancelSelector := mock2.SelectMock{}
	actionCancelSelector.On("Create", mock3.AnythingOfType("string"), mock3.AnythingOfType("string")).Return(&actionCancelSelector).Twice()
	actionCancelSelector.On("Run", mock3.AnythingOfType("[]ux.Described")).Return(0, "", nil).Once()
	actionCancelSelector.On("Run", mock3.AnythingOfType("[]ux.Described")).Return(-1, "", nil).Once()

	// copy state and assign select builder appropriate to a cancelled action select
	scesc := singleContainerState
	scesc.SelectBuilder = &actionCancelSelector
	actionCancelCtx := mock.MockContext{}
	actionCancelCtx.On("State").Return(&scesc)

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
			args{},
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
