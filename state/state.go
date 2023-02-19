package state

import (
	"context"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	interaction2 "github.com/kjkondratuk/goblins-and-gold/interaction"
)

const (
	StateKey = iota
)

type State struct {
	Player    actors.Player
	CurrRoom  *RoomDefinition
	World     *WorldDefinition
	PromptLib ux.PromptLib
}

func (s *State) Apply(r interaction2.Result) {
	s.Player.Acquire(r.AcquiredItems...)
}

func (s *State) AsContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, StateKey, s)
}
