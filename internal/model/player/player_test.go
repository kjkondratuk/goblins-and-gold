package player

import (
	"github.com/kjkondratuk/goblins-and-gold/internal/model/stats"
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
	p := NewPlayer()
	p.LevelUp()
	assert.Equal(t, 2, p.Lvl(), "Level should be incremented")

	p = &player{
		_baseStats:    stats.NewBaseStats(),
		_derivedStats: stats.DerivedStats{},
		_visionType:   0,
		_spd:          0,
		_hp:           0,
		_currHp:       0,
		_tempHp:       0,
	}

	p.LevelUp()
	assert.Equal(t, 20, p.Lvl(), "Should not level above max level")
}

func TestPlayer_Heal(t *testing.T) {
	p := NewPlayer(WithHp(20))
	p.Damage(5)
	p.Heal(3)
	assert.Equal(t, 18, p.Hp(), "Should do partial heal")

	if ok := p.Heal(10); ok {
		t.Errorf("Should not heal above max")
	} else {
		assert.Equal(t, 20, p.Hp(), "Should be at max when over-healed")
	}
}
