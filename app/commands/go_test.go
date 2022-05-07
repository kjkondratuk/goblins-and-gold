package commands

import (
	"flag"
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/kjkondratuk/goblins-and-gold/container"
	"github.com/kjkondratuk/goblins-and-gold/encounter"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	"github.com/kjkondratuk/goblins-and-gold/world"
	"github.com/kjkondratuk/goblins-and-gold/world/path"
	"github.com/kjkondratuk/goblins-and-gold/world/room"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"io/ioutil"
	"strings"
	"testing"
)

var (
	minimalApp   = &cli.App{Name: "test-app"}
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
			CurrRoom: &room.Definition{
				Name:                "",
				Description:         "",
				Paths:               nil,
				Containers:          nil,
				MandatoryEncounters: nil,
			},
			World: &world.Definition{
				Rooms: map[string]*room.Definition{
					"starting-room": {
						Name:                "",
						Description:         "",
						Paths:               []*path.Definition{},
						Containers:          []*container.Container{},
						MandatoryEncounters: []*encounter.Definition{},
					},
				},
				StartRoom: "starting-room",
			},
		},
	}}
)

func TestGoCommand_Action_NoPaths(t *testing.T) {
	cmd := basicCommand
	ux.Stdin = ioutil.NopCloser(strings.NewReader("go\n"))
	testContext := cli.NewContext(minimalApp, &flag.FlagSet{}, nil)
	err := cmd.Action(testContext)
	assert.NoError(t, err, fmt.Sprintf("Should not error invoking command with proper parameters: %v", err))
}

func TestGoCommand_Action_Navigates(t *testing.T) {
	cmd := basicCommand
	cmd._state.CurrRoom.Paths = []*path.Definition{
		{"new-room", "A splintered wooden door."},
	}
	ux.Stdin = ioutil.NopCloser(strings.NewReader("go\nj\n"))
	testContext := cli.NewContext(minimalApp, &flag.FlagSet{}, nil)
	err := cmd.Action(testContext)
	assert.NoError(t, err, fmt.Sprintf("Should not error navigating to a new room: %v", err))
}
