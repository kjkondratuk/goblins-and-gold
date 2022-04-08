package world

import "github.com/kjkondratuk/goblins-and-gold/room"

type World struct {
	Rooms     map[string]room.Room `yaml:"rooms"`
	StartRoom string               `yaml:"startingRoom"`
}

func (w World) Room(key string) (room.Room, bool) {
	if r, ok := w.Rooms[key]; ok {
		return r, true
	} else {
		return room.Room{}, false
	}
}
