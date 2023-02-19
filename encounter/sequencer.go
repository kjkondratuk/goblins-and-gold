package encounter

import (
	"container/ring"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/encounter/hasher"
	"github.com/pterm/pterm"
)

type sequencer struct {
	_turnHasher hasher.TurnHasher
	_turnOrder  *ring.Ring
	_player     actors.Player
	_fighters   map[string]actors.Combatant
	_turn       int
}

type Sequencer interface {
	IsDone() bool
	DoTurn(handler Handler)
}

type SequencerOpt func(s *sequencer)

func NewCombatSequencer(p actors.Player, m []actors.Monster, opts ...SequencerOpt) Sequencer {
	if p == nil || m == nil {
		return nil
	}

	combatants := make([]actors.Combatant, len(m)+1)
	for i, en := range m {
		combatants[i] = en
	}
	combatants[len(combatants)-1] = p

	s := &sequencer{
		_turn:       0,
		_player:     p,
		_turnHasher: hasher.InitiativeRolled,
	}

	for _, o := range opts {
		o(s)
	}

	turnOrder, fighters := s._turnHasher.HashTurnOrder(combatants...)
	s._turnOrder = turnOrder
	s._fighters = fighters

	return s
}

func WithHasher(h hasher.TurnHasher) SequencerOpt {
	return func(s *sequencer) {
		s._turnHasher = h
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
