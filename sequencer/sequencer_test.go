package sequencer

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/state/hasher"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	// These values need to be read-only as they are not thread-safe
	testPlayer = actors.NewPlayer(
		actors.PlayerParams{
			CombatantParams: actors.CombatantParams{
				Name:      "player",
				AC:        0,
				HP:        0,
				BaseStats: stats.BaseStats{},
				Inventory: nil,
				Attacks:   nil,
			},
		},
	)

	healthyPlayer = actors.NewPlayer(
		actors.PlayerParams{
			CombatantParams: actors.CombatantParams{
				Name:      "player",
				AC:        0,
				HP:        1,
				BaseStats: stats.BaseStats{},
				Inventory: nil,
				Attacks:   nil,
			},
		},
	)

	testMonster1 = actors.NewMonster(
		actors.MonsterParams{
			CombatantParams: actors.CombatantParams{
				Name:      "monster",
				AC:        0,
				HP:        0,
				BaseStats: stats.BaseStats{},
				Inventory: nil,
				Attacks:   nil,
			},
		},
	)

	singleTestMonster = []actors.Monster{
		testMonster1,
	}
)

func TestNewCombatSequencer_Success(t *testing.T) {
	seq := NewCombatSequencer(testPlayer, singleTestMonster, WithHasher(hasher.InitiativeDeclared)).(*sequencer)

	assert.NotNil(t, seq)
	assert.NotEqual(t, &sequencer{}, seq)
	assert.Equal(t, seq._turnHasher, hasher.InitiativeDeclared)
	assert.Equal(t, 0, seq._turn)
	assert.Equal(t, 2, seq._turnOrder.Len())
}

func TestNewCombatSequencer_NoPlayer(t *testing.T) {
	seq := NewCombatSequencer(nil, singleTestMonster, WithHasher(hasher.InitiativeDeclared))

	assert.Nil(t, seq)
}

func TestNewCombatSequencer_NilMonsters(t *testing.T) {
	var singleTestMonster []actors.Monster
	seq := NewCombatSequencer(testPlayer, singleTestMonster, WithHasher(hasher.InitiativeDeclared))

	assert.Nil(t, seq)
}

func TestSequencer_DoTurn(t *testing.T) {
	seq := NewCombatSequencer(testPlayer, singleTestMonster, WithHasher(hasher.InitiativeDeclared))

	seq.DoTurn(func(c actors.Combatant) {
		assert.Equal(t, "player", c.Name())
	})

	seq.DoTurn(func(c actors.Combatant) {
		assert.Equal(t, "monster", c.Name())
	})

	seq.DoTurn(func(c actors.Combatant) {
		assert.Equal(t, "player", c.Name())
	})

	seq.DoTurn(func(c actors.Combatant) {
		assert.Equal(t, "monster", c.Name())
	})
}

func TestSequencer_IsDone_NoPlayerHealth(t *testing.T) {
	seq := NewCombatSequencer(testPlayer, singleTestMonster, WithHasher(hasher.InitiativeDeclared))

	assert.True(t, seq.IsDone())
}

func TestSequencer_IsDone_PlayerAlive(t *testing.T) {
	seq := NewCombatSequencer(healthyPlayer, singleTestMonster, WithHasher(hasher.InitiativeDeclared))

	assert.False(t, seq.IsDone())
}

func TestSequencer_IsDone_NoEnemies(t *testing.T) {
	seq := NewCombatSequencer(healthyPlayer, []actors.Monster{}, WithHasher(hasher.InitiativeDeclared))

	assert.True(t, seq.IsDone())
}
