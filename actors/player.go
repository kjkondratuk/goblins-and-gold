package actors

import (
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/kjkondratuk/goblins-and-gold/model/item"
	"github.com/kjkondratuk/goblins-and-gold/model/loadout"
	"time"
)

type player struct {
	combatant
	equipped loadout.Loadout
}

type PlayerParams struct {
	CombatantParams `yaml:",inline"`
	//Loadout         string `yaml:"loadout"`
	loadout.Loadout `yaml:"loadout"`
}

type Player interface {
	Combatant
	Acquire(item ...item.Item)
	Summary() PlayerParams
}

type PlayerOption func(c *player)

func WithPlayerDice(d dice.Dice) PlayerOption {
	return func(m *player) {
		m._dice = d
	}
}

func NewPlayer(pd PlayerParams, opts ...PlayerOption) Player {
	p := &player{
		combatant{
			_name:      pd.Name,
			_type:      PlayerType,
			_dice:      dice.NewDice(time.Now().UnixNano()),
			_hp:        pd.HP,
			_ac:        pd.AC,
			_baseStats: pd.BaseStats,
			_inventory: pd.Inventory,
			_attacks:   pd.Attacks,
		}, pd.Loadout,
	}

	for _, o := range opts {
		o(p)
	}

	return p
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
			Attacks:   p._attacks,
			Inventory: p._inventory,
		}, Loadout: p.equipped,
	}
}
