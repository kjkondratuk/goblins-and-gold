// Code generated by mockery v2.34.2. DO NOT EDIT.

package container

import (
	interaction "github.com/kjkondratuk/goblins-and-gold/interaction"
	mock "github.com/stretchr/testify/mock"

	modelcontainer "github.com/kjkondratuk/goblins-and-gold/model/container"

	state "github.com/kjkondratuk/goblins-and-gold/state"
)

// MockInteractionFunc is an autogenerated mock type for the InteractionFunc type
type MockInteractionFunc struct {
	mock.Mock
}

type MockInteractionFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInteractionFunc) EXPECT() *MockInteractionFunc_Expecter {
	return &MockInteractionFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: s, c
func (_m *MockInteractionFunc) Execute(s state.State, c *modelcontainer.Container) (interaction.Result, error) {
	ret := _m.Called(s, c)

	var r0 interaction.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(state.State, *modelcontainer.Container) (interaction.Result, error)); ok {
		return rf(s, c)
	}
	if rf, ok := ret.Get(0).(func(state.State, *modelcontainer.Container) interaction.Result); ok {
		r0 = rf(s, c)
	} else {
		r0 = ret.Get(0).(interaction.Result)
	}

	if rf, ok := ret.Get(1).(func(state.State, *modelcontainer.Container) error); ok {
		r1 = rf(s, c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockInteractionFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockInteractionFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - s state.State
//   - c *modelcontainer.Container
func (_e *MockInteractionFunc_Expecter) Execute(s interface{}, c interface{}) *MockInteractionFunc_Execute_Call {
	return &MockInteractionFunc_Execute_Call{Call: _e.mock.On("Execute", s, c)}
}

func (_c *MockInteractionFunc_Execute_Call) Run(run func(s state.State, c *modelcontainer.Container)) *MockInteractionFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.State), args[1].(*modelcontainer.Container))
	})
	return _c
}

func (_c *MockInteractionFunc_Execute_Call) Return(_a0 interaction.Result, _a1 error) *MockInteractionFunc_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockInteractionFunc_Execute_Call) RunAndReturn(run func(state.State, *modelcontainer.Container) (interaction.Result, error)) *MockInteractionFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInteractionFunc creates a new instance of MockInteractionFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInteractionFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInteractionFunc {
	mock := &MockInteractionFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
