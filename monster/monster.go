package monster

import (
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/kjkondratuk/goblins-and-gold/item"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	"time"
)

type monster struct {
	_dice dice.Dice
	// TODO : this should probably be moved into BaseStats since it's dependent upon class and Con
	_hp        int
	_baseStats stats.BaseStats
	_inventory []item.Item
}

type MonsterData struct {
	HP        int             `yaml:"hp"`
	BaseStats stats.BaseStats `yaml:"stats"`
	Inventory []item.Item     `yaml:"inventory"`
}

type Monster interface {
	Dmg(hp int) bool
	Acquire(item ...item.Item)
	BaseStats() stats.BaseStats
	MonsterData() MonsterData
	Roll(rollExp string) int
}

func NewMonster(pd MonsterData) Monster {
	return &monster{
		_dice:      dice.NewDice(time.Now().UnixNano()),
		_hp:        pd.HP,
		_baseStats: pd.BaseStats,
		_inventory: pd.Inventory,
	}
}

// Dmg : applies damage to a PlayerStruct's hitpoints.
// Parameters:
//   - hp - int - the number of hitpoints to remove
// Returns:
//   - bool - whether or not the damage could be applied
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

func (p *monster) MonsterData() MonsterData {
	return MonsterData{
		HP:        p._hp,
		BaseStats: p._baseStats,
		Inventory: p._inventory,
	}
}
