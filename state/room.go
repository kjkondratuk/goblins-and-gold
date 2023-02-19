package state

import (
	"context"
)

type RoomDefinition struct {
	Name                string                 `yaml:"name"`
	Description         string                 `yaml:"description"`
	Paths               []*PathDefinition      `yaml:"paths"`
	Containers          []*Container           `yaml:"containers"`
	MandatoryEncounters []*EncounterDefinition `yaml:"mandatory_encounters"`
}

func (r *RoomDefinition) RunEncounters(ctx context.Context) {
	for _, e := range r.MandatoryEncounters {
		// TODO : returns an outcome.  do we need it?
		NewEncounter(*e).Run(ctx)
		p := ctx.Value(StateKey).(*State).Player
		if p.Unconscious() {
			break
		}
	}
}
