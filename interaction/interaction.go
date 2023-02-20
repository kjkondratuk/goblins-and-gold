package interaction

import (
	"github.com/kjkondratuk/goblins-and-gold/attack"
	"github.com/kjkondratuk/goblins-and-gold/item"
)

type Type string

func (t Type) Describe() string {
	return string(t)
}

type ResultType string

const (
	RT_Success = ResultType("Success")
	RT_Failure = ResultType("Failure")
)

type Result struct {
	Type          ResultType
	Message       string
	AcquiredItems []item.Item
	Damage        attack.DamageEffect
}
