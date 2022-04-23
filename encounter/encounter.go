package encounter

import (
	"container/ring"
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

func calcIndexKey(c actors.Combatant, dupCounter int) string {
	return fmt.Sprintf("%02d/%02d/%d", c.Roll(dice.D20)+c.BaseStats().DexMod(), c.BaseStats().Dex, dupCounter)
}

func (e *encounter) Run(p player.Player) Outcome {
	pterm.Info.Println(fmt.Sprintf("Running encounter..."))

	// Create a map of combatants with an index representing their turn order
	dupCounter := 0
	fighters := make(map[string]actors.Combatant)
	initiatives := make([]string, 1+len(e._enemies))

	// Roll initiative for the player and add them to the combatant map at the given index
	// index format is: <2 digit padded roll total>/<2 digit padded dexterity score>/<duplicate count>
	initiatives[0] = calcIndexKey(p, dupCounter)
	fighters[initiatives[0]] = p

	for i, enemy := range e._enemies {
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

	// populate the turn order key values into a circular data structure we can iterate
	turnOrder := ring.New(len(initiatives))
	for _, x := range initiatives {
		turnOrder.Value = x
		turnOrder = turnOrder.Next()
	}

	// TODO : finish implementing encounters
	// loop over list, taking a turn for each combatant
	curr := turnOrder
	for i := 0; i < len(initiatives); i++ {
		pterm.Success.Printfln("%s -> %s", curr.Value, curr.Next().Value)
		curr = curr.Move(1)
	}
	//turnOrder.Do(func(a any) {
	//	key := a.(int)
	//	pterm.Success.Printfln("Taking turn: %s", key)
	//	turnOrder.
	//})
	// monsters will roll a die to select an attack from an attack distribution, or by random by default
	// Calculate to-hit based on the combatant's modifier
	// If the attack hits, calculate damage
	// Apply damage to the target(s)
	// if the player hasn't been defeated and there are still monsters loop again

	return Outcome{}
}

func (e *encounter) Describe() string {
	return e._description
}
