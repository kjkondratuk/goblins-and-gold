package monster

import (
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/attack"
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/kjkondratuk/goblins-and-gold/item"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	"github.com/pterm/pterm"
	"time"
)

type monster struct {
	_name      string
	_dice      dice.Dice
	_hp        int // TODO: need to track current and max HP as well as temporary HP
	_baseStats stats.BaseStats
	_inventory []item.Item
	_attacks   attack.AttackSet
}

type Definition struct {
	Name      string           `yaml:"name"`
	HP        int              `yaml:"hp"`
	BaseStats stats.BaseStats  `yaml:"stats"`
	Inventory []item.Item      `yaml:"inventory"`
	Attacks   attack.AttackSet `yaml:"attacks"`
}

type Monster interface {
	actors.Combatant
	Acquire(item ...item.Item)
	Definition() Definition
}

func NewMonster(pd Definition) Monster {
	return &monster{
		_name:      pd.Name,
		_dice:      dice.NewDice(time.Now().UnixNano()),
		_hp:        pd.HP,
		_baseStats: pd.BaseStats,
		_inventory: pd.Inventory,
		_attacks:   pd.Attacks,
	}
}

func (m *monster) Name() string {
	return m._name
}

func (m *monster) Health() int {
	return m._hp
}

func (m *monster) Attack(c actors.Combatant) {
	r, _ := m._dice.Roll(fmt.Sprintf("1d%d", len(m._attacks)))
	keys := make([]string, len(m._attacks))
	i := 0
	for k := range m._attacks {
		keys[i] = k
	}

	// select randomly rolled attack and tabulate damage
	atk := keys[r]
	totalDamage := 0
	for _, v := range m._attacks[atk].Damage {
		dr, ok := m._dice.Roll(v.Roll)
		if !ok {
			pterm.Error.Printfln("Monster (%s) has invalid damage roll specifier: %s", m._name, v.Roll)
		}
		totalDamage += dr + v.Bonus
	}

	pterm.Info.Printfln("%s dealt %d damage to %s with a %s attack", m._name, totalDamage, c.Name(), atk)
	c.Dmg(totalDamage)
}

// Dmg : applies attack to a PlayerStruct's hitpoints.
// Parameters:
//   - hp - int - the number of hitpoints to remove
// Returns:
//   - bool - whether or not the attack could be applied
func (m *monster) Dmg(hp int) bool {
	if hp > 0 {
		m._hp -= hp
		if m._hp < 0 {
			m._hp = 0
		}
		return true
	}
	return false
}

func (m *monster) Acquire(item ...item.Item) {
	m._inventory = append(m._inventory, item...)
}

func (m *monster) Roll(rollExp string) int {
	r, _ := m._dice.Roll(rollExp)
	return r
}

func (m *monster) BaseStats() stats.BaseStats {
	return m._baseStats
}

func (m *monster) Definition() Definition {
	return Definition{
		Name:      m._name,
		HP:        m._hp,
		BaseStats: m._baseStats,
		Inventory: m._inventory,
	}
}
