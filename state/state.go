package state

import (
	"context"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/model/world"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"time"
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
	//Apply(r interaction2.Result)
	Context() context.Context
	Player() actors.Player
	Prompter() ux.PromptLib
	SetPrompter(p ux.PromptLib)
	CurrentRoom() string
	UpdateCurrentRoom(r string)
	World() *world.WorldDefinition
}

func New(pr ux.PromptLib, p actors.Player, r string, w *world.WorldDefinition) State {
	s := &state{}
	// TODO : not sure if this is the best move, but it stops errors in unit tests
	if s._c == nil {
		s._c = context.Background()
	}
	s._c = context.WithValue(s._c, PromptLibKey, pr)
	s._c = context.WithValue(s._c, PlayerKey, p)
	s._c = context.WithValue(s._c, RoomKey, r)
	s._c = context.WithValue(s._c, WorldKey, w)
	return s
}

//func (s *state) Apply(r interaction2.Result) {
//	s.Player().Acquire(r.AcquiredItems...)
//}

func (s *state) Context() context.Context {
	return s._c
}

func (s *state) Player() actors.Player {
	return s._c.Value(PlayerKey).(actors.Player)
}

func (s *state) Prompter() ux.PromptLib {
	return s._c.Value(PromptLibKey).(ux.PromptLib)
}

func (s *state) SetPrompter(p ux.PromptLib) {
	s._c = context.WithValue(s._c, PromptLibKey, p)
}

func (s *state) CurrentRoom() string {
	return s._c.Value(RoomKey).(string)
}

func (s *state) UpdateCurrentRoom(r string) {
	s._c = context.WithValue(s._c, RoomKey, r)
}

func (s *state) World() *world.WorldDefinition {
	return s._c.Value(WorldKey).(*world.WorldDefinition)
}

func (s *state) Deadline() (deadline time.Time, ok bool) {
	return s._c.Deadline()
}

func (s *state) Done() <-chan struct{} {
	return s._c.Done()
}

func (s *state) Err() error {
	return s._c.Err()
}

func (s *state) Value(key any) any {
	return s._c.Value(key)
}
