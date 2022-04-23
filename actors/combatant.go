package actors

import (
	"github.com/kjkondratuk/goblins-and-gold/attack"
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/kjkondratuk/goblins-and-gold/item"
	"github.com/kjkondratuk/goblins-and-gold/stats"
)

type CombatantParams struct {
	Name      string           `yaml:"name"`
	AC        int              `yaml:"ac"`
	HP        int              `yaml:"hp"`
	BaseStats stats.BaseStats  `yaml:"stats"`
	Inventory []item.Item      `yaml:"inventory"`
	Attacks   attack.AttackSet `yaml:"attacks"`
}

type combatant struct {
	_name      string
	_dice      dice.Dice
	_ac        int
	_hp        int // TODO: need to track current and max HP as well as temporary HP
	_baseStats stats.BaseStats
	_inventory []item.Item
	_attacks   attack.AttackSet
}

type Combatant interface {
	Name() string
	Health() int
	AC() int
	Dmg(hp int) bool
	BaseStats() stats.BaseStats
	Roll(rollExp string) int
	Attack(c Combatant)
	Unconscious() bool
}

// Dmg : applies attack to a PlayerStruct's hitpoints.
// Parameters:
//   - hp - int - the number of hitpoints to remove
// Returns:
//   - bool - whether or not the attack could be applied
func (c *combatant) Dmg(hp int) bool {
	if hp > 0 {
		c._hp -= hp
		if c._hp < 0 {
			c._hp = 0
		}
		return true
	}
	return false
}

func (c *combatant) Roll(rollExp string) int {
	r, _ := c._dice.Roll(rollExp)
	return r
}

func (c *combatant) BaseStats() stats.BaseStats {
	return c._baseStats
}

func (c *combatant) Unconscious() bool {
	return c._hp == 0
}

func (c *combatant) Name() string {
	return c._name
}

func (c *combatant) AC() int {
	return c._ac
}

func (c *combatant) Health() int {
	return c._hp
}
