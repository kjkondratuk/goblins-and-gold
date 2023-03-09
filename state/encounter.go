package state

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/sequencer"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"github.com/pterm/pterm"
)

const (
	TypeCombat = Type("Combat")

	ActionAttack = CombatAction("Attack")
	ActionRun    = CombatAction("Run")
)

type CombatAction string

func (ca CombatAction) Describe() string {
	return string(ca)
}

type Type string

type EncounterDefinition struct {
	Type        Type                   `yaml:"type"`
	Description string                 `yaml:"description"`
	Enemies     []actors.MonsterParams `yaml:"enemies"`
}

type encounter struct {
	_type        Type
	_description string
	_enemies     []actors.Monster
}

type Encounter interface {
	Run(s State) Outcome
	Enemies() []actors.Monster
}

type Outcome struct {
}

func NewEncounter(d EncounterDefinition) Encounter {
	monsters := make([]actors.Monster, len(d.Enemies))
	for i, m := range d.Enemies {
		monsters[i] = actors.NewMonster(m)
	}
	return &encounter{
		_type:        d.Type,
		_description: d.Description,
		_enemies:     monsters,
	}
}

func (e *encounter) Enemies() []actors.Monster {
	return e._enemies
}

func (e *encounter) Run(s State) Outcome {
	p := s.Player()
	pterm.Info.Println(e._description)

	m := make([]actors.Monster, len(e.Enemies()))
	md := make([]ux.Described, len(e.Enemies()))
	for i, en := range e.Enemies() {
		m[i] = en
		md[i] = en
	}
	//m[len(m)-1] = p

	seq := sequencer.NewCombatSequencer(p, m)

	// loop over list, taking a combat for each combatant, until done
	for !seq.IsDone() {
		seq.DoTurn(func(c actors.Combatant) {
			switch c.(type) {
			case actors.Player:
				// take player turn
				_, action, err := s.Prompter().Select("How do you respond?", []string{
					"Pass",
					"Attack",
					"Run",
				})
				if err != nil {
					pterm.Error.Printfln("There was a problem performing action [%s]: %s", action, err)
				}
				switch action {
				case "Attack":
					// TODO : finish implementing assailant choice
					//atkIdx, _, err := s.Prompter().Select("Who do you attack?", append([]string{}, ux.DescribeToList(md)...))
					//if err != nil {
					//	return
					//}
				case "Run":
					// TODO : implement dex contest to escape
				}
			case actors.Monster:
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
