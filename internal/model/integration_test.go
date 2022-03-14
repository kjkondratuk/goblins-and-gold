package model

import (
	"github.com/stretchr/testify/assert"
	"goblins-and-gold/internal/model/room"
	"testing"
)

func TestAll(t *testing.T) {
	r := room.NewRoom(room.WithDescription("This is a new room"))
	assert.Equal(t, "This is a new room", r.Description())
}
