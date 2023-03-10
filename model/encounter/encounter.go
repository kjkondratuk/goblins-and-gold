package encounter

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/ux"
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
	Enemies() []actors.Monster
	ux.Described
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

func (e *encounter) Describe() string {
	return e._description
}
