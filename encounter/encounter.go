package encounter

import (
	"context"
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/actors/monster"
	"github.com/kjkondratuk/goblins-and-gold/actors/player"
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/pterm/pterm"
)

const (
	TypeCombat = Type("Combat")

	ActionAttack = CombatAction("Attack")
	ActionRun    = CombatAction("Run")
)

var (
	combatActions = []ux.Described{
		ActionAttack,
		ActionRun,
	}
)

type CombatAction string

func (ca CombatAction) Describe() string {
	return string(ca)
}

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
	Enemies() []monster.Monster
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

func (e *encounter) Enemies() []monster.Monster {
	return e._enemies
}

func (e *encounter) Run(p player.Player) Outcome {
	pterm.Info.Println(e._description)

	seq := NewCombatSequencer(p, e)

	// loop over list, taking a combat for each combatant, until done
	for !seq.IsDone() {
		seq.DoTurn(func(c actors.Combatant) {
			switch c.(type) {
			case player.Player:
				// take player turn
				// TODO : figure out how to handle the deeply nested nature of combat--it makes it difficult to check stats and such as combat progresses
				_, _ = ux.NewSelector("Pass", "How do you respond?", func(ctx context.Context, idx int, val string, err error) (interface{}, error) {
					//_, _ = ux.NewSelector("Back")
					return nil, nil
				}).Run(context.Background(), combatActions)
			case monster.Monster:
				// take monster turn
				c.Attack(p)
				//pterm.Success.Printfln("Taking monster turn: %s", c.Name())
			default:
				pterm.Error.Printfln("Invalid combatant type: %T", c)
			}
		})
	}
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
