package world

import "github.com/kjkondratuk/goblins-and-gold/world/room"

type Definition struct {
	Rooms     map[string]*room.Definition `yaml:"rooms"`
	StartRoom string                      `yaml:"startingRoom"`
}

// Room : retrieves a room by its key from the rooms present in this world
// Parameters:
//   - key - string - the key of the room
// Return:
//   - room.Definition - the room
//   - bool - whether or not the room could be located
func (w Definition) Room(key string) (room.Definition, bool) {
	if r, ok := w.Rooms[key]; ok {
		return *r, true
	} else {
		return room.Definition{}, false
	}
}
