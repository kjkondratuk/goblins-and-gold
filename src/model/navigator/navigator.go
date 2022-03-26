package navigator

import (
	"github.com/kjkondratuk/goblins-and-gold/src/model/player"
	"github.com/kjkondratuk/goblins-and-gold/src/model/room"
)

type navigator struct {
	_p player.Player
	_r room.Room
}

type Navigator interface {
}

// NewNavigatorFrom
func NewNavigatorFrom(p player.Player, r room.Room) Navigator {
	return &navigator{
		_p: p,
		_r: r,
	}
}

// MoveTo : blindly updates the room navigation pointer to the specified room
func (n *navigator) MoveTo(r room.Room) {
	n._r = r
}
