package encounter

import (
	"container/ring"
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/pterm/pterm"
	"sort"
)

type sequencer struct {
	_turnOrder *ring.Ring
	_player    actors.Combatant
	_fighters  map[string]actors.Combatant
	_turn      int
}

type Sequencer interface {
	IsDone() bool
	DoTurn(handler Handler)
}

func NewCombatSequencer(p actors.Player, e Encounter) Sequencer {
	// Create a map of combatants with an index representing their combat order
	dupCounter := 0
	fighters := make(map[string]actors.Combatant)
	initiatives := make([]string, 1+len(e.Enemies()))

	// Roll initiative for the player and add them to the combatant map at the given index
	// index format is: <2 digit padded roll total>/<2 digit padded dexterity score>/<duplicate count>
	initiatives[0] = calcIndexKey(p, dupCounter)
	fighters[initiatives[0]] = p

	for i, enemy := range e.Enemies() {
		// calculate an index key -- increment and recalculate if there is a collision
		newIndexKey := calcIndexKey(enemy, dupCounter)
		if _, collides := fighters[newIndexKey]; collides {
			dupCounter++
			newIndexKey = calcIndexKey(enemy, dupCounter)
		}

		// add the combatant to the participants and add index value
		initiatives[i+1] = newIndexKey
		fighters[initiatives[i+1]] = enemy
	}

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

	return &sequencer{
		_turn:      0,
		_turnOrder: turnOrder,
		_fighters:  fighters,
		_player:    p,
	}
}

type Handler func(c actors.Combatant)

func (s *sequencer) DoTurn(handler Handler) {
	fighter := s._fighters[s._turnOrder.Value.(string)]
	pterm.Debug.Printfln("Doing turn: %d - %s", s._turn, fighter.Name())
	handler(fighter)
	s._turn++
	s._turnOrder = s._turnOrder.Move(1)
}

func (s *sequencer) IsDone() bool {
	return s._player.Health() <= 0 || s._turnOrder.Len() <= 1
}

func calcIndexKey(c actors.Combatant, dupCounter int) string {
	return fmt.Sprintf("%02d/%02d/%d", c.Roll(dice.D20)+c.BaseStats().DexMod(), c.BaseStats().Dex, dupCounter)
}
