package interaction

import (
	"github.com/kjkondratuk/goblins-and-gold/model/attack"
	"github.com/kjkondratuk/goblins-and-gold/model/item"
)

type Type string

func (t Type) Describe() string {
	return string(t)
}

type ResultType string

const (
	RtSuccess = ResultType("Success")
	RtFailure = ResultType("Failure")

	InteractionTypeCancel = Type("Cancel")
	InteractionTypeOpen   = Type("Open")
	InteractionTypeLoot   = Type("Loot")
	InteractionTypeUnlock = Type("Unlock")
)

type Result struct {
	Type          ResultType
	Message       string
	AcquiredItems []item.Item
	Damage        attack.DamageEffect
}
