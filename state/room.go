package state

type RoomDefinition struct {
	Name                string                 `yaml:"name"`
	Description         string                 `yaml:"description"`
	Paths               []*PathDefinition      `yaml:"paths"`
	Containers          []*Container           `yaml:"containers"`
	MandatoryEncounters []*EncounterDefinition `yaml:"mandatory_encounters"`
}

func (r *RoomDefinition) RunEncounters(s State) {
	for _, e := range r.MandatoryEncounters {
		// TODO : returns an outcome.  do we need it?
		NewEncounter(*e).Run(s)
		p := s.Player()
		if p.Unconscious() {
			break
		}
	}
}
