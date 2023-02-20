package state

import (
	"context"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	interaction2 "github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"github.com/urfave/cli"
)

const (
	PlayerKey = iota
	CliContextKey
	RoomKey
	WorldKey
	PromptLibKey
)

type state struct {
	_c context.Context
}

type State interface {
	SetCliContext(c *cli.Context)
	CliContext() *cli.Context
	Apply(r interaction2.Result)
	Context() context.Context
	SetContext(c context.Context)
	Player() actors.Player
	Prompter() ux.PromptLib
	CurrentRoom() *RoomDefinition
	UpdateCurrentRoom(r *RoomDefinition)
	World() *WorldDefinition
}

type InteractionFunc func(c State) (interaction2.Result, error)

func New(pr ux.PromptLib, p actors.Player, r *RoomDefinition, w *WorldDefinition) State {
	s := &state{}
	s._c = context.WithValue(s._c, PromptLibKey, pr)
	s._c = context.WithValue(s._c, PlayerKey, p)
	s._c = context.WithValue(s._c, RoomKey, r)
	s._c = context.WithValue(s._c, WorldKey, w)
	return s
}

func (s *state) Apply(r interaction2.Result) {
	s.Player().Acquire(r.AcquiredItems...)
}

func (s *state) Context() context.Context {
	return s._c
}

func (s *state) SetContext(c context.Context) {
	s._c = c
}

func (s *state) Player() actors.Player {
	return s._c.Value(PlayerKey).(actors.Player)
}

func (s *state) Prompter() ux.PromptLib {
	return s._c.Value(PromptLibKey).(ux.PromptLib)
}

func (s *state) CurrentRoom() *RoomDefinition {
	return s._c.Value(RoomKey).(*RoomDefinition)
}

func (s *state) UpdateCurrentRoom(r *RoomDefinition) {
	s._c = context.WithValue(s._c, RoomKey, r)
}

func (s *state) World() *WorldDefinition {
	return s._c.Value(WorldKey).(*WorldDefinition)
}

func (s *state) SetCliContext(c *cli.Context) {
	s._c = context.WithValue(s._c, CliContextKey, c)
}

func (s *state) CliContext() *cli.Context {
	return s._c.Value(CliContextKey).(*cli.Context)
}
