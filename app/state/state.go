package state

import (
	"github.com/kjkondratuk/goblins-and-gold/actors/player"
	interaction2 "github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/world"
	"github.com/kjkondratuk/goblins-and-gold/world/room"
)

type State struct {
	Player   player.Player
	CurrRoom *room.Definition
	World    *world.Definition
}

func (s *State) Apply(r interaction2.Result) {
	s.Player.Acquire(r.AcquiredItems...)
}
