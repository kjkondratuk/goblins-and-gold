package world

import (
	"github.com/kjkondratuk/goblins-and-gold/container"
	"github.com/kjkondratuk/goblins-and-gold/encounter"
)

type Room struct {
	Name                string                 `yaml:"name"`
	Description         string                 `yaml:"description"`
	Paths               []*Path                `yaml:"paths"`
	Containers          []*container.Container `yaml:"containers"`
	MandatoryEncounters []*encounter.Encounter `yaml:"mandatory_encounters"`
}
