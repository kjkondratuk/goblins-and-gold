package command

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/model/stats"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStatsCommand(t *testing.T) {
	c := NewStatsCommand()
	assert.NotNil(t, c)
	assert.Len(t, c.Subcommands(), 1)
}

func TestStatsCommand_Run(t *testing.T) {
	t.Run("should execute valid subcommands successfully", func(t *testing.T) {
		c := NewStatsCommand()
		err := c.Run(nil, "help")
		assert.NoError(t, err)
	})

	t.Run("should error for invalid subcommands", func(t *testing.T) {
		c := NewStatsCommand()
		err := c.Run(nil, "other")
		assert.Error(t, err)
	})

	t.Run("should print nothing when there is no player to print stats for", func(t *testing.T) {
		c := NewStatsCommand()
		err := c.Run(nil)
		assert.NoError(t, err)
	})

	t.Run("should print player stats when player is present", func(t *testing.T) {
		c := NewStatsCommand()
		err := c.Run(state.New(nil, actors.NewPlayer(actors.PlayerParams{
			CombatantParams: actors.CombatantParams{
				Name: "a player",
				AC:   15,
				HP:   50,
				BaseStats: stats.BaseStats{
					Lvl: 5,
					Str: 10,
					Dex: 14,
					Con: 13,
					Int: 11,
					Wis: 12,
					Cha: 16,
				},
			},
		}), "", nil))
		assert.NoError(t, err)
	})
}
