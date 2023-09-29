// Code generated by mockery v2.34.2. DO NOT EDIT.

package command

import (
	state "github.com/kjkondratuk/goblins-and-gold/state"
	mock "github.com/stretchr/testify/mock"
)

// MockCommand is an autogenerated mock type for the Command type
type MockCommand struct {
	mock.Mock
}

type MockCommand_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCommand) EXPECT() *MockCommand_Expecter {
	return &MockCommand_Expecter{mock: &_m.Mock}
}

// Aliases provides a mock function with given fields:
func (_m *MockCommand) Aliases() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockCommand_Aliases_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Aliases'
type MockCommand_Aliases_Call struct {
	*mock.Call
}

// Aliases is a helper method to define mock.On call
func (_e *MockCommand_Expecter) Aliases() *MockCommand_Aliases_Call {
	return &MockCommand_Aliases_Call{Call: _e.mock.On("Aliases")}
}

func (_c *MockCommand_Aliases_Call) Run(run func()) *MockCommand_Aliases_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommand_Aliases_Call) Return(_a0 []string) *MockCommand_Aliases_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommand_Aliases_Call) RunAndReturn(run func() []string) *MockCommand_Aliases_Call {
	_c.Call.Return(run)
	return _c
}

// Description provides a mock function with given fields:
func (_m *MockCommand) Description() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockCommand_Description_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Description'
type MockCommand_Description_Call struct {
	*mock.Call
}

// Description is a helper method to define mock.On call
func (_e *MockCommand_Expecter) Description() *MockCommand_Description_Call {
	return &MockCommand_Description_Call{Call: _e.mock.On("Description")}
}

func (_c *MockCommand_Description_Call) Run(run func()) *MockCommand_Description_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommand_Description_Call) Return(_a0 string) *MockCommand_Description_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommand_Description_Call) RunAndReturn(run func() string) *MockCommand_Description_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *MockCommand) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockCommand_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type MockCommand_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *MockCommand_Expecter) Name() *MockCommand_Name_Call {
	return &MockCommand_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *MockCommand_Name_Call) Run(run func()) *MockCommand_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommand_Name_Call) Return(_a0 string) *MockCommand_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommand_Name_Call) RunAndReturn(run func() string) *MockCommand_Name_Call {
	_c.Call.Return(run)
	return _c
}

// Run provides a mock function with given fields: s, args
func (_m *MockCommand) Run(s state.State, args ...string) error {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, s)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(state.State, ...string) error); ok {
		r0 = rf(s, args...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockCommand_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockCommand_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - s state.State
//   - args ...string
func (_e *MockCommand_Expecter) Run(s interface{}, args ...interface{}) *MockCommand_Run_Call {
	return &MockCommand_Run_Call{Call: _e.mock.On("Run",
		append([]interface{}{s}, args...)...)}
}

func (_c *MockCommand_Run_Call) Run(run func(s state.State, args ...string)) *MockCommand_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]string, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(string)
			}
		}
		run(args[0].(state.State), variadicArgs...)
	})
	return _c
}

func (_c *MockCommand_Run_Call) Return(_a0 error) *MockCommand_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommand_Run_Call) RunAndReturn(run func(state.State, ...string) error) *MockCommand_Run_Call {
	_c.Call.Return(run)
	return _c
}

// Subcommands provides a mock function with given fields:
func (_m *MockCommand) Subcommands() []Command {
	ret := _m.Called()

	var r0 []Command
	if rf, ok := ret.Get(0).(func() []Command); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Command)
		}
	}

	return r0
}

// MockCommand_Subcommands_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Subcommands'
type MockCommand_Subcommands_Call struct {
	*mock.Call
}

// Subcommands is a helper method to define mock.On call
func (_e *MockCommand_Expecter) Subcommands() *MockCommand_Subcommands_Call {
	return &MockCommand_Subcommands_Call{Call: _e.mock.On("Subcommands")}
}

func (_c *MockCommand_Subcommands_Call) Run(run func()) *MockCommand_Subcommands_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommand_Subcommands_Call) Return(_a0 []Command) *MockCommand_Subcommands_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommand_Subcommands_Call) RunAndReturn(run func() []Command) *MockCommand_Subcommands_Call {
	_c.Call.Return(run)
	return _c
}

// Usage provides a mock function with given fields:
func (_m *MockCommand) Usage() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockCommand_Usage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Usage'
type MockCommand_Usage_Call struct {
	*mock.Call
}

// Usage is a helper method to define mock.On call
func (_e *MockCommand_Expecter) Usage() *MockCommand_Usage_Call {
	return &MockCommand_Usage_Call{Call: _e.mock.On("Usage")}
}

func (_c *MockCommand_Usage_Call) Run(run func()) *MockCommand_Usage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCommand_Usage_Call) Return(_a0 string) *MockCommand_Usage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockCommand_Usage_Call) RunAndReturn(run func() string) *MockCommand_Usage_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCommand creates a new instance of MockCommand. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCommand(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCommand {
	mock := &MockCommand{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}