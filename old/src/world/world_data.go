package world

import "github.com/kjkondratuk/goblins-and-gold/src/room"

type WorldData struct {
	StartRoom room.RoomData `yaml:"startingRoom"`
}
