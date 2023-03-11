package actors

import (
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/kjkondratuk/goblins-and-gold/ux"
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

func (m *monster) Describe() string {
	return m._desc
}
