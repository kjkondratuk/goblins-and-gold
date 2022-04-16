package world

type World struct {
	Rooms     map[string]*Room `yaml:"rooms"`
	StartRoom string           `yaml:"startingRoom"`
}

// Room : retrieves a room by its key from the rooms present in this world
// Parameters:
//   - key - string - the key of the room
// Return:
//   - room.Room - the room
//   - bool - whether or not the room could be located
func (w World) Room(key string) (Room, bool) {
	if r, ok := w.Rooms[key]; ok {
		return *r, true
	} else {
		return Room{}, false
	}
}
