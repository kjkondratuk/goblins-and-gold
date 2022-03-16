package player

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPlayer(t *testing.T) {
	p := NewPlayer()
	assert.Equal(t, p.Hp(), 1, "HP should be equal to 1 by default")
	assert.Equal(t, p.Lvl(), 1, "Level should be equal to 1 by default")
	assert.Equal(t, p.BaseStats().Str, 10, "Stats should default to 10")

	p = NewPlayer(WithHp(3473))
	assert.Equal(t, p.Hp(), 3473, "HP should equal the HP option specified")
}

func TestPlayer_Damage(t *testing.T) {
	p := NewPlayer(WithHp(100))
	if ok := p.Damage(72); ok {
		assert.Equal(t, 28, p.Hp(), "HP should decrement with damage")
	} else {
		t.Errorf("Expected HP to decrement properly with damage")
	}

	if ok := p.Damage(45); ok {
		t.Errorf("Expected HP to fail decrementing below 0")
	} else {
		assert.Equal(t, 0, p.Hp(), "HP should be zero if decremented below zero")
	}
}

func TestPlayer_LevelUp(t *testing.T) {

}
