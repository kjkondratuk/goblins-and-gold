package navigator

import (
	"github.com/kjkondratuk/goblins-and-gold/src/player"
	"github.com/kjkondratuk/goblins-and-gold/src/room"
)

type navigator struct {
	_p player.Player
	_r room.Room
}

type Navigator interface {
	Player() player.Player
	Room() room.Room
	MoveTo(r room.Room)
}

// NewNavigatorFrom : creates a new navigator with the given player starting in the given room.
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

func (n *navigator) Player() player.Player {
	return n._p
}

func (n *navigator) Room() room.Room {
	return n._r
}
