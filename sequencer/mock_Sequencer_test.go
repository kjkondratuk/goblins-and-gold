// Code generated by mockery v2.34.2. DO NOT EDIT.

package sequencer

import (
	actors "github.com/kjkondratuk/goblins-and-gold/actors"
	mock "github.com/stretchr/testify/mock"
)

// MockSequencer is an autogenerated mock type for the Sequencer type
type MockSequencer struct {
	mock.Mock
}

type MockSequencer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSequencer) EXPECT() *MockSequencer_Expecter {
	return &MockSequencer_Expecter{mock: &_m.Mock}
}

// DoTurn provides a mock function with given fields: handler
func (_m *MockSequencer) DoTurn(handler Handler) {
	_m.Called(handler)
}

// MockSequencer_DoTurn_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DoTurn'
type MockSequencer_DoTurn_Call struct {
	*mock.Call
}

// DoTurn is a helper method to define mock.On call
//   - handler Handler
func (_e *MockSequencer_Expecter) DoTurn(handler interface{}) *MockSequencer_DoTurn_Call {
	return &MockSequencer_DoTurn_Call{Call: _e.mock.On("DoTurn", handler)}
}

func (_c *MockSequencer_DoTurn_Call) Run(run func(handler Handler)) *MockSequencer_DoTurn_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Handler))
	})
	return _c
}

func (_c *MockSequencer_DoTurn_Call) Return() *MockSequencer_DoTurn_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSequencer_DoTurn_Call) RunAndReturn(run func(Handler)) *MockSequencer_DoTurn_Call {
	_c.Call.Return(run)
	return _c
}

// IsDone provides a mock function with given fields:
func (_m *MockSequencer) IsDone() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockSequencer_IsDone_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsDone'
type MockSequencer_IsDone_Call struct {
	*mock.Call
}

// IsDone is a helper method to define mock.On call
func (_e *MockSequencer_Expecter) IsDone() *MockSequencer_IsDone_Call {
	return &MockSequencer_IsDone_Call{Call: _e.mock.On("IsDone")}
}

func (_c *MockSequencer_IsDone_Call) Run(run func()) *MockSequencer_IsDone_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockSequencer_IsDone_Call) Return(_a0 bool) *MockSequencer_IsDone_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSequencer_IsDone_Call) RunAndReturn(run func() bool) *MockSequencer_IsDone_Call {
	_c.Call.Return(run)
	return _c
}

// Terminate provides a mock function with given fields: c
func (_m *MockSequencer) Terminate(c actors.Combatant) {
	_m.Called(c)
}

// MockSequencer_Terminate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Terminate'
type MockSequencer_Terminate_Call struct {
	*mock.Call
}

// Terminate is a helper method to define mock.On call
//   - c actors.Combatant
func (_e *MockSequencer_Expecter) Terminate(c interface{}) *MockSequencer_Terminate_Call {
	return &MockSequencer_Terminate_Call{Call: _e.mock.On("Terminate", c)}
}

func (_c *MockSequencer_Terminate_Call) Run(run func(c actors.Combatant)) *MockSequencer_Terminate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(actors.Combatant))
	})
	return _c
}

func (_c *MockSequencer_Terminate_Call) Return() *MockSequencer_Terminate_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockSequencer_Terminate_Call) RunAndReturn(run func(actors.Combatant)) *MockSequencer_Terminate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSequencer creates a new instance of MockSequencer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSequencer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSequencer {
	mock := &MockSequencer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
