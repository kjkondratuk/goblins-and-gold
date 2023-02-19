package hasher

import (
	"container/ring"
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"sort"
	"strconv"
)

type turnHasher struct {
	_hasher HashFunc
}

type HashFunc func(combatant actors.Combatant, unique int) string

var (
	InitiativeRolled = NewTurnHasher(func(c actors.Combatant, unique int) string {
		return fmt.Sprintf("%02d/%02d/%d", c.Roll(dice.D20)+c.BaseStats().DexMod(), c.BaseStats().Dex, unique)
	})
	InitiativeDeclared = NewTurnHasher(func(c actors.Combatant, unique int) string {
		return strconv.Itoa(unique)
	})
)

type TurnHasher interface {
	HashTurnOrder(pc ...actors.Combatant) (*ring.Ring, map[string]actors.Combatant)
}

func NewTurnHasher(h HashFunc) TurnHasher {
	return &turnHasher{_hasher: h}
}

func (h *turnHasher) HashTurnOrder(c ...actors.Combatant) (*ring.Ring, map[string]actors.Combatant) {
	// Create a map of combatants with an index representing their combat order
	dupCounter := 0
	fighters := make(map[string]actors.Combatant)
	initiatives := make([]string, len(c))

	for i, enemy := range c {
		// calculate an index key -- increment and recalculate if there is a collision
		newIndexKey := h._hasher(enemy, dupCounter)
		if _, collides := fighters[newIndexKey]; collides {
			dupCounter++
			newIndexKey = h._hasher(enemy, dupCounter)
		}

		// add the combatant to the participants and add index value
		initiatives[i] = newIndexKey
		fighters[initiatives[i]] = enemy
	}

	return setTurnOrder(initiatives), fighters
}

func setTurnOrder(initiatives []string) *ring.Ring {
	// Sort in ascending order, then reverse it
	sort.Strings(initiatives)
	for i, j := 0, len(initiatives)-1; i < j; i, j = i+1, j-1 {
		initiatives[i], initiatives[j] = initiatives[j], initiatives[i]
	}

	// populate the combat order key values into a circular data structure we can iterate
	turnOrder := ring.New(len(initiatives))
	for _, x := range initiatives {
		turnOrder.Value = x
		turnOrder = turnOrder.Next()
	}

	return turnOrder
}
