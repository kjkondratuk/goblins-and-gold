package encounter

import (
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/actors/monster"
	"github.com/kjkondratuk/goblins-and-gold/actors/player"
	"github.com/kjkondratuk/goblins-and-gold/dice"
	"github.com/pterm/pterm"
	"sort"
)

const (
	TypeCombat = Type("Combat")
)

type Type string

type Definition struct {
	Type        Type                 `yaml:"type"`
	Description string               `yaml:"description"`
	Enemies     []monster.Definition `yaml:"enemies"`
}

type encounter struct {
	_type        Type
	_description string
	_enemies     []monster.Monster
}

type Encounter interface {
	Run(p player.Player) Outcome
}

type Outcome struct {
}

func NewEncounter(d Definition) Encounter {
	monsters := make([]monster.Monster, len(d.Enemies))
	for i, m := range d.Enemies {
		monsters[i] = monster.NewMonster(m)
	}
	return &encounter{
		_type:        d.Type,
		_description: d.Description,
		_enemies:     monsters,
	}
}

func (e *encounter) Run(p player.Player) Outcome {
	pterm.Info.Println(fmt.Sprintf("Running encounter..."))

	// Create a map of combatants and keep a list of keys for sequencing combat
	turns := make(map[int]actors.Combatant)
	inits := make([]int, 1+len(e._enemies))
	pInit := p.Roll(dice.D20) + p.BaseStats().DexMod()
	turns[pInit] = p
	inits[0] = pInit
	for i, e := range e._enemies {
		inits[i+1] = e.BaseStats().DexMod() + e.Roll(dice.D20)
		turns[inits[i+1]] = e
	}

	// Sort in ascending order, then reverse it
	sort.Ints(inits)
	for i, j := 0, len(inits)-1; i < j; i, j = i+1, j-1 {
		inits[i], inits[j] = inits[j], inits[i]
	}

	return Outcome{}
}

func (e *encounter) Describe() string {
	return e._description
}
