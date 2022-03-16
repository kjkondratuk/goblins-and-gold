package model

import (
	"github.com/kjkondratuk/goblins-and-gold/internal/model/room"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAll(t *testing.T) {
	r := room.NewRoom(room.WithDescription("This is a new room"))
	assert.Equal(t, "This is a new room", r.Description())
}
