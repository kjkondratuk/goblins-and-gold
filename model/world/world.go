package world

import "github.com/kjkondratuk/goblins-and-gold/model/room"

type WorldDefinition struct {
	Rooms     map[string]*room.RoomDefinition `yaml:"rooms"`
	StartRoom string                          `yaml:"startingRoom"`
}

// Room : retrieves a room by its key from the rooms present in this world
// Parameters:
//   - key - string - the key of the room
//
// Return:
//   - room.PathDefinition - the room
//   - bool - whether the room could be located
//func (w WorldDefinition) Room(key string) (string, room.RoomDefinition, bool) {
//	if r, ok := w.Rooms[key]; ok {
//		return key, *r, true
//	} else {
//		return "", room.RoomDefinition{}, false
//	}
//}
