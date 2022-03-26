package navigator

import (
	"github.com/kjkondratuk/goblins-and-gold/src/player"
	"github.com/kjkondratuk/goblins-and-gold/src/room"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNavigatorFrom(t *testing.T) {
	n := NewNavigatorFrom(
		player.NewPlayer(player.WithHp(147)),
		room.NewRoom(room.WithDescription("Test Room")))

	assert.Equal(t, 147, n.Player().Hp(), "The player for the navigator should be the same one assigned.")
	assert.Equal(t, "Test Room", n.Room().Description(), "The room for the navigator should be the same one assigned.")
}

func TestNavigator_MoveTo(t *testing.T) {
	n := NewNavigatorFrom(
		player.NewPlayer(player.WithHp(147)),
		room.NewRoom(room.WithDescription("Test Room")))

	assert.Equal(t, "Test Room", n.Room().Description(), "The room for the navigator should be the same one assigned initially.")

	n.MoveTo(room.NewRoom(room.WithDescription("New Test Room")))

	assert.Equal(t, "New Test Room", n.Room().Description(), "The room for the navigator should be the same one moved to.")
}
