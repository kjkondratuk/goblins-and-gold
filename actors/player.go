package actors

import (
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/kjkondratuk/goblins-and-gold/item"
	"github.com/pterm/pterm"
	"time"
)

type player struct {
	combatant
}

type PlayerParams struct {
	CombatantParams `yaml:",inline"`
}

type Player interface {
	Combatant
	Acquire(item ...item.Item)
	Summary() PlayerParams
}

func NewPlayer(pd PlayerParams) Player {
	return &player{
		combatant{
			_name:      pd.Name,
			_dice:      dice.NewDice(time.Now().UnixNano()),
			_hp:        pd.HP,
			_ac:        pd.AC,
			_baseStats: pd.BaseStats,
			_inventory: pd.Inventory,
			_attacks:   pd.Attacks,
		},
	}
}

func (p *player) Attack(c Combatant) {
	pterm.Error.Printfln("Not implemented yet")
}

func (c *combatant) Acquire(item ...item.Item) {
	c._inventory = append(c._inventory, item...)
}

func (p *player) Summary() PlayerParams {
	return PlayerParams{
		CombatantParams: CombatantParams{
			Name:      p._name,
			AC:        p._ac,
			HP:        p._hp,
			BaseStats: p._baseStats,
			Inventory: p._inventory,
		},
	}
}
