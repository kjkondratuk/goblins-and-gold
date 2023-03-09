package state

type WorldDefinition struct {
	Rooms     map[string]*RoomDefinition `yaml:"rooms"`
	StartRoom string                     `yaml:"startingRoom"`
}

// Room : retrieves a room by its key from the rooms present in this world
// Parameters:
//   - key - string - the key of the room
//
// Return:
//   - room.PathDefinition - the room
//   - bool - whether or not the room could be located
func (w WorldDefinition) Room(key string) (RoomDefinition, bool) {
	if r, ok := w.Rooms[key]; ok {
		return *r, true
	} else {
		return RoomDefinition{}, false
	}
}
