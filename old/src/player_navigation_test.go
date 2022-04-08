package src

import (
	"github.com/kjkondratuk/goblins-and-gold/src/navigator"
	"github.com/kjkondratuk/goblins-and-gold/src/player"
	"github.com/kjkondratuk/goblins-and-gold/src/room"
	"github.com/kjkondratuk/goblins-and-gold/src/stats"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerNavigates_RoomToRoom(t *testing.T) {
	r := room.NewRoom(room.WithDescription("This is a new room"))
	assert.Equal(t, "This is a new room", r.Description())

	p := player.NewPlayer(player.WithHp(12), player.WithBaseStats(stats.NewBaseStats(
		stats.WithLvl(1),
	)))

	n := navigator.NewNavigatorFrom(p, r)
	n.MoveTo(room.NewRoom(room.WithDescription("New room")))
	assert.Equal(t, "New room", n.Room().Description())
}
