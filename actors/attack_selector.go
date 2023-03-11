package actors

import (
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/model/attack"
)

type AttackSelector interface {
	Select(c Combatant) (string, attack.Attack)
}

type RandomAttackSelector struct{}

func (s RandomAttackSelector) Select(c Combatant) (string, attack.Attack) {
	r := c.Roll(fmt.Sprintf("1d%d", len(c.Attacks())-1))
	keys := make([]string, len(c.Attacks()))
	i := 0
	for k := range c.Attacks() {
		keys[i] = k
		i++
	}

	// select randomly rolled attack and tabulate damage
	ak := keys[r]
	atk := c.Attacks()[ak]

	return ak, atk
}

type ElectiveAttackSelector struct {
	Attack string
}

func (s ElectiveAttackSelector) Select(c Combatant) (string, attack.Attack) {
	atk := c.Attacks()[s.Attack]
	ak := s.Attack
	return ak, atk
}
