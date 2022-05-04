package commands

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	"github.com/kjkondratuk/goblins-and-gold/world"
	"github.com/kjkondratuk/goblins-and-gold/world/path"
	"github.com/kjkondratuk/goblins-and-gold/world/room"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var (
	emptyRoom = room.Definition{
		Name: "empty-room",
	}
	onePathRoom = room.Definition{
		Name:        "empty-room",
		Description: "",
		Paths: []*path.Definition{
			{"another-room", "A plain wooden door."},
		},
		Containers:          nil,
		MandatoryEncounters: nil,
	}
	emptyPlayer = actors.NewPlayer(actors.PlayerParams{CombatantParams: actors.CombatantParams{
		Name:      "",
		AC:        0,
		HP:        0,
		BaseStats: stats.BaseStats{},
		Inventory: nil,
		Attacks:   nil,
	}})
)

func StateWithRoom(r room.Definition) state.State {
	return state.State{
		Player:   emptyPlayer,
		CurrRoom: &r,
		World: &world.Definition{
			Rooms: map[string]*room.Definition{
				"empty-room": &r,
			},
			StartRoom: "empty-room",
		},
	}
}

func TestGo_NoPaths(t *testing.T) {
	ux.Stdin = ioutil.NopCloser(strings.NewReader("go\n"))
	s := StateWithRoom(emptyRoom)
	app := cli.App{
		Commands: []cli.Command{
			NewGoCommand(&s),
		},
	}
	os.Args = []string{""}
	err := app.Run(os.Args)
	assert.NoError(t, err, "Should not error running go command when there are no paths.")
	assert.Equal(t, "empty-room", s.CurrRoom.Name, "Should be in the same room.")
}

func TestGo_NavigatesToNewRoom(t *testing.T) {
	args := []string{"", "go"}
	ux.Stdin = ioutil.NopCloser(strings.NewReader("j\n"))
	s := StateWithRoom(onePathRoom)
	app := &cli.App{
		Commands: []cli.Command{
			NewGoCommand(&s),
		},
	}
	err := app.Run(args)
	assert.NoError(t, err, "Should not error running go command with a single path.")
	assert.Equal(t, "another-room", s.CurrRoom.Name, "Should be in the next room.")
}
