package actors

import (
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/kjkondratuk/goblins-and-gold/model/attack"
	"github.com/kjkondratuk/goblins-and-gold/model/item"
	"github.com/kjkondratuk/goblins-and-gold/model/stats"
	"github.com/pterm/pterm"
)

const (
	MonsterType = "monster"
	PlayerType  = "player"
)

type Type string

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
	_type      Type
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
	Attacks() attack.AttackSet
	Inventory() []item.Item
	Attack(c Combatant, s AttackSelector) bool
	Unconscious() bool
}

// Dmg : applies attack to a PlayerStruct's hitpoints.
// Parameters:
//   - hp - int - the number of hitpoints to remove
//
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

func (c *combatant) Attacks() attack.AttackSet {
	return c._attacks
}

func (c *combatant) Inventory() []item.Item {
	return c._inventory
}

func (c *combatant) Attack(t Combatant, s AttackSelector) bool {
	ak, atk := s.Select(c)

	// Figure out if the attack hits or not
	dr, _ := c._dice.Roll(dice.D20)
	dr += atk.Bonus

	if dr >= t.AC() {
		// Calculate total damage for the strike
		totalDamage := 0
		for _, v := range atk.Damage {
			dr, ok := c._dice.Roll(v.Roll)
			if !ok {
				pterm.Error.Printfln("%s has invalid damage roll specifier: %s", c.Name(), v.Roll)
			}
			totalDamage += dr + v.Bonus
		}

		t.Dmg(totalDamage)
		var pr pterm.PrefixPrinter
		if c._type == MonsterType {
			pr = pterm.Warning
		} else {
			pr = pterm.Success
		}
		pr.Printfln("%s dealt %d damage with a %s attack.  %s has %d HP remaining.", c.Name(), totalDamage, ak, t.Name(), t.Health())
		if t.Health() <= 0 {
			return true
		}
	} else {
		var pr pterm.PrefixPrinter
		if c._type == PlayerType {
			pr = pterm.Warning
		} else {
			pr = pterm.Info
		}
		pr.Printfln("%s strikes at %s with a %s attack and misses. (%d)", c.Name(), t.Name(), ak, dr)
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
