package interaction

import (
	"context"
	"github.com/kjkondratuk/goblins-and-gold/damage"
	"github.com/kjkondratuk/goblins-and-gold/item"
)

const (
	InteractionDataKey = iota
	ContainerDataKey
	PlayerDataKey
)

type Type string

func (t Type) Describe() string {
	return string(t)
}

type Func func(c context.Context) (Result, error)

type ResultType string

const (
	RT_Success = ResultType("Success")
	RT_Failure = ResultType("Failure")
)

type Result struct {
	Type          ResultType
	Message       string
	AcquiredItems []item.Item
	Damage        damage.Damage
}
