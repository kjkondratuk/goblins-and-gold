package actors

import (
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"github.com/pterm/pterm"
	"time"
)

type monster struct {
	combatant
	_desc string
}

type MonsterParams struct {
	CombatantParams `yaml:",inline"`
	Description     string `yaml:"description"`
}

type Monster interface {
	Combatant
	ux.Described
}

type MonsterOption func(c *monster)

func WithMonsterDice(d dice.Dice) MonsterOption {
	return func(m *monster) {
		m._dice = d
	}
}

func NewMonster(pd MonsterParams, opts ...MonsterOption) Monster {
	m := &monster{
		combatant{
			_name:      pd.Name,
			_dice:      dice.NewDice(time.Now().UnixNano()),
			_hp:        pd.HP,
			_ac:        pd.AC,
			_baseStats: pd.BaseStats,
			_inventory: pd.Inventory,
			_attacks:   pd.Attacks,
		},
		pd.Description,
	}

	for _, o := range opts {
		o(m)
	}

	return m
}

func (m *monster) Attack(c Combatant) {
	r, _ := m._dice.Roll(fmt.Sprintf("1d%d", len(m._attacks)))
	keys := make([]string, len(m._attacks))
	i := 0
	for k := range m._attacks {
		keys[i] = k
		i++
	}

	// select randomly rolled attack and tabulate damage
	ak := keys[r-1]
	atk := m._attacks[ak]

	// Figure out if the attack hits or not
	dr, _ := m._dice.Roll(dice.D20)
	dr += atk.Bonus

	if dr >= c.AC() {
		// Calculate total damage for the strike
		totalDamage := 0
		for _, v := range atk.Damage {
			dr, ok := m._dice.Roll(v.Roll)
			if !ok {
				pterm.Error.Printfln("Monster (%s) has invalid damage roll specifier: %s", m.Name(), v.Roll)
			}
			totalDamage += dr + v.Bonus
		}

		c.Dmg(totalDamage)
		pterm.Info.Printfln("%s dealt %d damage with a %s attack.  You have %d HP remaining.", m.Name(), totalDamage, ak, c.Health())
	} else {
		pterm.Info.Printfln("The %s strikes at you and misses. (%d)", m.Name(), dr)
	}
}

func (m *monster) Describe() string {
	return m._desc
}
