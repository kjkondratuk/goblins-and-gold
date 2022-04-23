package monster

import (
	"github.com/kjkondratuk/goblins-and-gold/attack"
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/kjkondratuk/goblins-and-gold/item"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	"time"
)

type monster struct {
	_name      string
	_dice      dice.Dice
	_hp        int // TODO: need to track current and max HP as well as temporary HP
	_baseStats stats.BaseStats
	_inventory []item.Item
	_attacks   []attack.AttackSet
}

type Definition struct {
	Name      string          `yaml:"name"`
	HP        int             `yaml:"hp"`
	BaseStats stats.BaseStats `yaml:"stats"`
	Inventory []item.Item     `yaml:"inventory"`
}

type Monster interface {
	Dmg(hp int) bool
	BaseStats() stats.BaseStats
	Roll(rollExp string) int
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
	}
}

// Dmg : applies attack to a PlayerStruct's hitpoints.
// Parameters:
//   - hp - int - the number of hitpoints to remove
// Returns:
//   - bool - whether or not the attack could be applied
func (p *monster) Dmg(hp int) bool {
	if hp > 0 {
		p._hp -= hp
		if p._hp < 0 {
			p._hp = 0
		}
		return true
	}
	return false
}

func (p *monster) Acquire(item ...item.Item) {
	p._inventory = append(p._inventory, item...)
}

func (p *monster) Roll(rollExp string) int {
	r, _ := p._dice.Roll(rollExp)
	return r
}

func (p *monster) BaseStats() stats.BaseStats {
	return p._baseStats
}

func (p *monster) Definition() Definition {
	return Definition{
		Name:      p._name,
		HP:        p._hp,
		BaseStats: p._baseStats,
		Inventory: p._inventory,
	}
}
