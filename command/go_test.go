package command

import (
	"errors"
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/model/container"
	"github.com/kjkondratuk/goblins-and-gold/model/encounter"
	"github.com/kjkondratuk/goblins-and-gold/model/path"
	"github.com/kjkondratuk/goblins-and-gold/model/room"
	"github.com/kjkondratuk/goblins-and-gold/model/stats"
	"github.com/kjkondratuk/goblins-and-gold/model/world"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"github.com/stretchr/testify/assert"
	mock2 "github.com/stretchr/testify/mock"
	"os"
	"testing"
)

func TestNewGoCommand(t *testing.T) {
	exit := make(chan os.Signal)
	qc := NewQuitCommand(exit)
	lc := NewLookCommand()
	c := NewGoCommand(qc, lc)
	assert.NotNil(t, c)
	assert.Equal(t, "go", c.Name())
	assert.Len(t, c.Subcommands(), 1)
	assert.Equal(t, qc, c.(*goCommand).qc)
}

func Test_GoCommand_Run(t *testing.T) {
	sameRoomReturningSelectPrompt := &ux.MockPromptLib{}
	sameRoomReturningSelectPrompt.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
		Return(0, "", nil)

	errorReturningSelectPrompt := &ux.MockPromptLib{}
	errorReturningSelectPrompt.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
		Return(0, "", errors.New("something bad happened"))

	nextRoomReturningSelectPrompt := &ux.MockPromptLib{}
	nextRoomReturningSelectPrompt.EXPECT().Select(mock2.AnythingOfType("string"), mock2.AnythingOfType("[]string")).
		Return(1, "another-room", nil)

	type fields struct {
		qc Command
		lc Command
	}
	type args struct {
		s    state.State
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"should do nothing when nowhere to go",
			fields{
				qc: nil,
			},
			args{
				s: state.New(nil, nil, "start-room", &world.WorldDefinition{Rooms: map[string]*room.RoomDefinition{
					"start-room": {
						Description: "some room",
					},
				}}),
				args: []string{},
			},
			func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		}, {
			"should stay in the current room when selected",
			fields{qc: nil},
			args{
				s: state.New(sameRoomReturningSelectPrompt, nil, "next-room", &world.WorldDefinition{Rooms: map[string]*room.RoomDefinition{
					"next-room": {
						Description: "some room 1 description",
					}, "another-room": {
						Description: "some other room description",
					},
				}}),
				args: nil,
			},
			func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		}, {
			"should return an error if there was an error prompting the user for a choice",
			fields{qc: nil},
			args{
				s: state.New(errorReturningSelectPrompt, nil, "start-room", &world.WorldDefinition{
					Rooms: map[string]*room.RoomDefinition{
						"start-room": {
							Name:        "start-room",
							Description: "",
							Paths: []*path.PathDefinition{
								{
									Room:        "next-room",
									Description: "A dark hallway",
								},
							},
							Containers:          nil,
							MandatoryEncounters: nil,
						},
						"next-room": {
							Name:                "next-room",
							Description:         "",
							Paths:               nil,
							Containers:          nil,
							MandatoryEncounters: nil,
						},
					},
					StartRoom: "start-room",
				}),
				args: nil,
			},
			func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		}, {
			"should return an error if the room being traversed to does not exist",
			fields{qc: nil},
			args{
				s: state.New(nextRoomReturningSelectPrompt, nil, "next-room", &world.WorldDefinition{
					Rooms: map[string]*room.RoomDefinition{
						"start-room": {
							"start-room",
							"some description",
							[]*path.PathDefinition{},
							[]*container.Container{},
							[]*encounter.EncounterDefinition{},
						},
					},
					StartRoom: "start-room",
				}),
				args: nil,
			},
			func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		}, {
			name:   "should exit successfully when the new room exists",
			fields: fields{qc: nil, lc: NewLookCommand()},
			args: args{
				s: state.New(nextRoomReturningSelectPrompt,
					actors.NewPlayer(actors.PlayerParams{CombatantParams: actors.CombatantParams{
						"player",
						10,
						10,
						stats.BaseStats{},
						nil,
						nil,
					}}),
					"next-room", &world.WorldDefinition{
						Rooms: map[string]*room.RoomDefinition{
							"start-room": {
								"start-room",
								"some description",
								[]*path.PathDefinition{},
								[]*container.Container{},
								[]*encounter.EncounterDefinition{},
							}, "next-room": {
								"next-room",
								"some other description",
								[]*path.PathDefinition{},
								[]*container.Container{},
								[]*encounter.EncounterDefinition{},
							},
						},
						StartRoom: "start-room",
					}),
				args: nil,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		}, {
			name:   "should exit successfully when player is unconscious",
			fields: fields{qc: &baseCommand{}, lc: NewLookCommand()},
			args: args{
				s: state.New(nextRoomReturningSelectPrompt,
					actors.NewPlayer(actors.PlayerParams{CombatantParams: actors.CombatantParams{
						"player",
						10,
						0,
						stats.BaseStats{},
						nil,
						nil,
					}}),
					"next-room", &world.WorldDefinition{
						Rooms: map[string]*room.RoomDefinition{
							"start-room": {
								"start-room",
								"some description",
								[]*path.PathDefinition{},
								[]*container.Container{},
								[]*encounter.EncounterDefinition{},
							}, "next-room": {
								"next-room",
								"some other description",
								[]*path.PathDefinition{},
								[]*container.Container{},
								[]*encounter.EncounterDefinition{},
							},
						},
						StartRoom: "start-room",
					}),
				args: nil,
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewGoCommand(tt.fields.qc, tt.fields.lc)
			tt.wantErr(t, g.Run(tt.args.s, tt.args.args...), fmt.Sprintf("Run(%v, %v)", tt.args.s, tt.args.args))
		})
	}
}
