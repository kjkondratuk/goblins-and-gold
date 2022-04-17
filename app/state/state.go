package state

import (
	interaction2 "github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/player"
	"github.com/kjkondratuk/goblins-and-gold/world"
)

type State struct {
	Player   *player.PlayerStruct
	CurrRoom *world.Room
	World    *world.World
}

func (s *State) Apply(r interaction2.Result) {
	s.Player.Acquire(r.AcquiredItems...)
}
