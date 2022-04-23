package room

import (
	"github.com/kjkondratuk/goblins-and-gold/actors"
	"github.com/kjkondratuk/goblins-and-gold/container"
	"github.com/kjkondratuk/goblins-and-gold/encounter"
	"github.com/kjkondratuk/goblins-and-gold/world/path"
)

type Definition struct {
	Name                string                  `yaml:"name"`
	Description         string                  `yaml:"description"`
	Paths               []*path.Definition      `yaml:"paths"`
	Containers          []*container.Container  `yaml:"containers"`
	MandatoryEncounters []*encounter.Definition `yaml:"mandatory_encounters"`
}

func (r *Definition) RunEncounters(p actors.Player) {
	for _, e := range r.MandatoryEncounters {
		// TODO : returns an outcome.  do we need it?
		encounter.NewEncounter(*e).Run(p)
		if p.Unconscious() {
			break
		}
	}
}
