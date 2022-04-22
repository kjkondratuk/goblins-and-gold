package encounter

import (
	"fmt"
	"github.com/kjkondratuk/goblins-and-gold/monster"
	"github.com/kjkondratuk/goblins-and-gold/player"
	"github.com/pterm/pterm"
)

const (
	TypeCombat = Type("combat")
)

type Type string

type EncounterData struct {
	Type        Type              `yaml:"type"`
	Description string            `yaml:"description"`
	Enemies     []monster.Monster `yaml:"enemies"`
}

type encouter struct {
	_type        Type
	_description string
	_enemies     []monster.Monster
}

type Encounter interface {
	Run(p player.Player) Result
}

type Result struct {
}

func NewEncounter(d EncounterData) Encounter {
	return &encouter{
		_type:        d.Type,
		_description: d.Description,
		_enemies:     d.Enemies,
	}
}

func (e *encouter) Run(p player.Player) Result {
	pterm.Info.Println(fmt.Sprintf("Running encounter..."))
	return Result{}
}
