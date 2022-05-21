package mock

import (
	"github.com/kjkondratuk/goblins-and-gold/app/state"
	"github.com/stretchr/testify/mock"
	"github.com/urfave/cli"
)

type MockContext struct {
	mock.Mock
}

func (c *MockContext) State() *state.State {
	args := c.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*state.State)
}

func (c *MockContext) Context() *cli.Context {
	panic("not implemented")
}

func (c *MockContext) RunCommandByName(name string) error {
	args := c.Called(name)
	return args.Error(0)
}
