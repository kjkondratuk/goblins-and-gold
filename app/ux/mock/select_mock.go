package mock

import (
	"github.com/kjkondratuk/goblins-and-gold/app/ux"
	"github.com/stretchr/testify/mock"
)

type SelectMock struct {
	mock.Mock
	ux.SelectBuilder
	ux.Select
}

func (b *SelectMock) Create(cancel string, label string) ux.Select {
	args := b.Called(cancel, label)
	return args.Get(0).(ux.Select)
}

func (b *SelectMock) Run(items []ux.Described) (int, string, error) {
	args := b.Called(items)
	return args.Int(0), args.String(1), args.Error(2)
}
