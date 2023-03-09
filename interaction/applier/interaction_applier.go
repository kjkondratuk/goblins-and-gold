package applier

import (
	"github.com/kjkondratuk/goblins-and-gold/interaction"
	"github.com/kjkondratuk/goblins-and-gold/state"
)

type InteractionApplier struct{}

func (a InteractionApplier) Apply(s state.State, r interaction.Result) {
	s.Player().Acquire(r.AcquiredItems...)
}
