package model

import (
	"testing"

	"github.com/kjkondratuk/goblins-and-gold/internal/model/player"
	"github.com/kjkondratuk/goblins-and-gold/internal/model/room"
	"github.com/kjkondratuk/goblins-and-gold/internal/model/stats"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	r := room.NewRoom(room.WithDescription("This is a new room"))
	assert.Equal(t, "This is a new room", r.Description())

	p := player.NewPlayer(player.WithHp(12), player.WithBaseStats(stats.NewBaseStats(
		stats.WithLvl(1),
		stats.
	)))
}
