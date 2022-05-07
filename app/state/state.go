package state

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	interaction2 "github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/world"
	"github.com/kjkondratuk/goblins-and-gold/world/room"
)

type State struct {
	Player        actors.Player
	CurrRoom      *room.Definition
	World         *world.Definition
	SelectBuilder ux.SelectBuilder
}

func (s *State) Apply(r interaction2.Result) {
	s.Player.Acquire(r.AcquiredItems...)
}
