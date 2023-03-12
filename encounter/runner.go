package encounter

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/model/encounter"
	"github.com/kjkondratuk/goblins-and-gold/model/item"
	"github.com/kjkondratuk/goblins-and-gold/sequencer"
	"github.com/kjkondratuk/goblins-and-gold/state"
	"github.com/kjkondratuk/goblins-and-gold/ux"
	"github.com/pterm/pterm"
)

type Outcome struct {
	Won   bool
	Slain []actors.Monster
}

type encounterRunner struct{}

type EncounterRunner interface {
	Run(s state.State, e encounter.Encounter) Outcome
}

func NewRunner() EncounterRunner {
	return &encounterRunner{}
}

func (er *encounterRunner) Run(s state.State, e encounter.Encounter) Outcome {
	p := s.Player()
	pterm.Warning.Println(e.Describe())

	m := make([]actors.Monster, len(e.Enemies()))
	md := make([]ux.Described, len(e.Enemies()))
	for i, en := range e.Enemies() {
		m[i] = en
		md[i] = en
	}
	//m[len(m)-1] = p

	seq := sequencer.NewCombatSequencer(p, m)

	outcome := Outcome{}

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
					targetIdx, _, _ := s.Prompter().Select("Who do you attack?", append([]string{}, ux.DescribeToList(md)...))
					ad := make([]string, len(p.Attacks()))
					i := 0
					for k, _ := range p.Attacks() {
						ad[i] = k
						i++
					}
					_, atkLabel, _ := s.Prompter().Select("How do you attack?", append([]string{}, ad...))

					killed := p.Attack(m[targetIdx], actors.ElectiveAttackSelector{Attack: atkLabel})
					if killed {
						seq.Terminate(m[targetIdx])
						outcome.Slain = append(outcome.Slain, m[targetIdx])
						outcome.Won = true

						pterm.Success.Printfln("%s vanquished the %s", s.Player().Name(), m[targetIdx].Name())

						var lootableInventory []item.Item
						var itemDesc []ux.Described
						for _, m := range outcome.Slain {
							for _, i := range m.Inventory() {
								lootableInventory = append(lootableInventory, i)
								itemDesc = append(itemDesc, i)
							}
						}

						// handle distribution of loot resulting from encounter
						var done bool
						for !done {
							i, _, _ := s.Prompter().Select("Which items do you want to loot?",
								append([]string{"Done"}, ux.DescribeToList(itemDesc)...))
							if i <= 0 {
								done = true
							} else {
								s.Player().Acquire(lootableInventory[i-1])
								lootableInventory = append(lootableInventory[:i-1], lootableInventory[i:]...)
								itemDesc = append(itemDesc[:i-1], itemDesc[i:]...)
							}
						}
					}
				case "Run":
					// TODO : implement dex contest to escape
				}
			case actors.Monster:
				// take monster turn
				_ = c.Attack(p, actors.RandomAttackSelector{})
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

	return outcome
}
