package state

import (
	"github.com/kjkondratuk/goblins-and-gold/player"
	"github.com/kjkondratuk/goblins-and-gold/world/room"
)

type GameState struct {
	Player   *player.Player
	CurrRoom *room.Room
}
