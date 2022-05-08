package commands

import (
	"flag"
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	selectmock "github.com/kjkondratuk/goblins-and-gold/app/ux/mock"
	"github.com/kjkondratuk/goblins-and-gold/container"
	"github.com/kjkondratuk/goblins-and-gold/encounter"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	"github.com/kjkondratuk/goblins-and-gold/world"
	"github.com/kjkondratuk/goblins-and-gold/world/path"
	"github.com/kjkondratuk/goblins-and-gold/world/room"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/urfave/cli"
	"testing"
)

var (
	minimalApp   = &cli.App{Name: "test-app"}
	startingRoom = room.Definition{
		Name:                "starting-room",
		Description:         "",
		Paths:               []*path.Definition{},
		Containers:          []*container.Container{},
		MandatoryEncounters: []*encounter.Definition{},
	}
	newRoom = room.Definition{
		Name:                "new-room",
		Description:         "",
		Paths:               []*path.Definition{},
		Containers:          []*container.Container{},
		MandatoryEncounters: []*encounter.Definition{},
	}
	basicCommand = goCommand{baseCommand: baseCommand{
		_cmd: cli.Command{},
		_state: &state.State{
			Player: actors.NewPlayer(actors.PlayerParams{
				CombatantParams: actors.CombatantParams{
					Name:      "testPlayer",
					AC:        0,
					HP:        10,
					BaseStats: stats.BaseStats{},
					Inventory: nil,
					Attacks:   nil,
				}},
			),
			CurrRoom: &startingRoom,
			World: &world.Definition{
				Rooms: map[string]*room.Definition{
					"starting-room": &startingRoom,
					"new-room":      &newRoom,
				},
				StartRoom: "starting-room",
			},
		},
	}}
)

func TestGoCommand_Action_NoPaths(t *testing.T) {
	cmd := basicCommand
	selector := selectmock.SelectMock{}
	selector.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&selector)
	selector.On("Run", mock.AnythingOfType("[]ux.Described")).Return(-1, "", nil)
	cmd._state.SelectBuilder = &selector

	testContext := cli.NewContext(minimalApp, &flag.FlagSet{}, nil)
	err := cmd.Action(testContext)
	assert.NoError(t, err, fmt.Sprintf("Should not error invoking command with proper parameters: %v", err))
	assert.Equal(t, "starting-room", cmd._state.CurrRoom.Name, "Should not have moved if there is nowhere to go")
}

func TestGoCommand_Action_Navigates(t *testing.T) {
	cmd := basicCommand
	selector := selectmock.SelectMock{}
	selector.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&selector)
	selector.On("Run", mock.AnythingOfType("[]ux.Described")).Return(0, "A splintered wooden door.", nil)
	cmd._state.SelectBuilder = &selector
	cmd._state.CurrRoom.Paths = []*path.Definition{
		{"new-room", "A splintered wooden door."},
	}
	testContext := cli.NewContext(minimalApp, &flag.FlagSet{}, nil)
	err := cmd.Action(testContext)
	assert.NoError(t, err, fmt.Sprintf("Should not error navigating to a new room: %v", err))
	assert.Equal(t, "new-room", cmd._state.CurrRoom.Name, "Should have navigated to the new room")
}

func TestGoCommand_Action_ErrorOnValidate(t *testing.T) {
	cmd := basicCommand
	selector := selectmock.SelectMock{}
	selector.On("Create", mock.AnythingOfType("string"), mock.AnythingOfType("string")).
		Return(&selector)
	selector.On("Run", mock.AnythingOfType("[]ux.Described")).Return(0, "A splintered wooden door.", nil)
	testContext := cli.NewContext(minimalApp, &flag.FlagSet{}, nil)
	err := cmd.Action(testContext)
	assert.NoError(t, err, fmt.Sprintf("Should not error navigating to a new room: %v", err))
	assert.Equal(t, "starting-room", cmd._state.CurrRoom.Name, "Should not have moved if there is nowhere to go")
}
