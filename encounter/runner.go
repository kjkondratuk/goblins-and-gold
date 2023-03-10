package encounter

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/model/encounter"
	"github.com/kjkondratuk/goblins-and-gold/sequencer"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"github.com/pterm/pterm"
)

type Outcome struct{}

type encounterRunner struct{}

type EncounterRunner interface {
	Run(s state.State, e encounter.Encounter) Outcome
}

func NewRunner() EncounterRunner {
	return &encounterRunner{}
}

func (er *encounterRunner) Run(s state.State, e encounter.Encounter) Outcome {
	p := s.Player()
	pterm.Info.Println(e.Describe())

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
