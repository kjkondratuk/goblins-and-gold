package commands

import (
	"github.com/stretchr/testify/mock"
	"github.com/urfave/cli"
)

type MockContext struct {
	mock.Mock
	cli.Context
}

func (mc *MockContext) Args() cli.Args {
	args := mc.Called()
	return args.Get(0).(cli.Args)
}
