package player

import (
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/kjkondratuk/goblins-and-gold/item"
	"github.com/kjkondratuk/goblins-and-gold/stats"
	"time"
)

type PlayerStruct struct {
	_dice dice.Dice
	// TODO : this should probably be moved into BaseStats since it's dependent upon class and Con
	HP        int             `yaml:"hp"`
	BaseStats stats.BaseStats `yaml:"stats"`
	Inventory []item.Item     `yaml:"inventory"`
}

type Player interface {
	Dmg(hp int) bool
	Acquire(item ...item.Item)
	Roll(rollExp string) int
}

func NewPlayerStruct() *PlayerStruct {
	return &PlayerStruct{
		_dice: dice.NewDice(time.Now().UnixNano()),
	}
}

// Dmg : applies damage to a PlayerStruct's hitpoints.
// Parameters:
//   - hp - int - the number of hitpoints to remove
// Returns:
//   - bool - whether or not the damage could be applied
func (p *PlayerStruct) Dmg(hp int) bool {
	if hp > 0 {
		p.HP -= hp
		if p.HP < 0 {
			p.HP = 0
		}
		return true
	}
	return false
}

func (p *PlayerStruct) Acquire(item ...item.Item) {
	p.Inventory = append(p.Inventory, item...)
}

func (p *PlayerStruct) Roll(rollExp string) int {
	r, _ := p._dice.Roll(rollExp)
	return r
}
