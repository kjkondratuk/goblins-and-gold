package player

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPlayer_Default(t *testing.T) {
	p := NewPlayer()
	assert.Equal(t, p.Hp, 1, "HP should be equal to 1 by default")
	assert.Equal(t, p.Lvl, 1, "Level should be equal to 1 by default")
	assert.Equal(t, p.BaseStats.Str, 10, "Stats should default to 10")
}
