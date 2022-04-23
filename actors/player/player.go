package player

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/attack"
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/kjkondratuk/goblins-and-gold/item"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	"time"
)

type player struct {
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

type Player interface {
	actors.Combatant
	Acquire(item ...item.Item)
	Definition() Definition
}

func NewPlayer(pd Definition) Player {
	return &player{
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
func (p *player) Dmg(hp int) bool {
	if hp > 0 {
		p._hp -= hp
		if p._hp < 0 {
			p._hp = 0
		}
		return true
	}
	return false
}

func (p *player) Acquire(item ...item.Item) {
	p._inventory = append(p._inventory, item...)
}

func (p *player) Roll(rollExp string) int {
	r, _ := p._dice.Roll(rollExp)
	return r
}

func (p *player) BaseStats() stats.BaseStats {
	return p._baseStats
}

func (p *player) Definition() Definition {
	return Definition{
		Name:      p._name,
		HP:        p._hp,
		BaseStats: p._baseStats,
		Inventory: p._inventory,
	}
}
