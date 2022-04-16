package state

import (
	interaction2 "github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/player"
	"github.com/kjkondratuk/goblins-and-gold/world/room"
)

type GameState struct {
	Player   *player.Player
	CurrRoom *room.Room
}

func (s *GameState) Apply(r interaction2.Result) {
	for _, i := range r.AcquiredItems {
		s.Player.Inventory = append(s.Player.Inventory, i)
	}
}
