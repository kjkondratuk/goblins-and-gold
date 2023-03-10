package room

import (
	"github.com/kjkondratuk/goblins-and-gold/model/container"
	"github.com/kjkondratuk/goblins-and-gold/model/encounter"
	"github.com/kjkondratuk/goblins-and-gold/model/path"
)

type RoomDefinition struct {
	Name                string                           `yaml:"name"`
	Description         string                           `yaml:"description"`
	Paths               []*path.PathDefinition           `yaml:"paths"`
	Containers          []*container.Container           `yaml:"containers"`
	MandatoryEncounters []*encounter.EncounterDefinition `yaml:"mandatory_encounters"`
}
