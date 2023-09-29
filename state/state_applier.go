package state

import "github.com/kjkondratuk/goblins-and-gold/interaction"

type AppliableType interface {
	interaction.Result
}

type Applier[T AppliableType] interface {
	Apply(s State, t T)
}
