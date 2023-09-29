// Code generated by mockery v2.34.2. DO NOT EDIT.

package state

import (
	context "context"

	actors "github.com/kjkondratuk/goblins-and-gold/actors"

	mock "github.com/stretchr/testify/mock"

	ux "github.com/kjkondratuk/goblins-and-gold/ux"

	world "github.com/kjkondratuk/goblins-and-gold/model/world"
)

// MockState is an autogenerated mock type for the State type
type MockState struct {
	mock.Mock
}

type MockState_Expecter struct {
	mock *mock.Mock
}

func (_m *MockState) EXPECT() *MockState_Expecter {
	return &MockState_Expecter{mock: &_m.Mock}
}

// Context provides a mock function with given fields:
func (_m *MockState) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// MockState_Context_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Context'
type MockState_Context_Call struct {
	*mock.Call
}

// Context is a helper method to define mock.On call
func (_e *MockState_Expecter) Context() *MockState_Context_Call {
	return &MockState_Context_Call{Call: _e.mock.On("Context")}
}

func (_c *MockState_Context_Call) Run(run func()) *MockState_Context_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockState_Context_Call) Return(_a0 context.Context) *MockState_Context_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockState_Context_Call) RunAndReturn(run func() context.Context) *MockState_Context_Call {
	_c.Call.Return(run)
	return _c
}

// CurrentRoom provides a mock function with given fields:
func (_m *MockState) CurrentRoom() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockState_CurrentRoom_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CurrentRoom'
type MockState_CurrentRoom_Call struct {
	*mock.Call
}

// CurrentRoom is a helper method to define mock.On call
func (_e *MockState_Expecter) CurrentRoom() *MockState_CurrentRoom_Call {
	return &MockState_CurrentRoom_Call{Call: _e.mock.On("CurrentRoom")}
}

func (_c *MockState_CurrentRoom_Call) Run(run func()) *MockState_CurrentRoom_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockState_CurrentRoom_Call) Return(_a0 string) *MockState_CurrentRoom_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockState_CurrentRoom_Call) RunAndReturn(run func() string) *MockState_CurrentRoom_Call {
	_c.Call.Return(run)
	return _c
}

// Player provides a mock function with given fields:
func (_m *MockState) Player() actors.Player {
	ret := _m.Called()

	var r0 actors.Player
	if rf, ok := ret.Get(0).(func() actors.Player); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(actors.Player)
		}
	}

	return r0
}

// MockState_Player_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Player'
type MockState_Player_Call struct {
	*mock.Call
}

// Player is a helper method to define mock.On call
func (_e *MockState_Expecter) Player() *MockState_Player_Call {
	return &MockState_Player_Call{Call: _e.mock.On("Player")}
}

func (_c *MockState_Player_Call) Run(run func()) *MockState_Player_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockState_Player_Call) Return(_a0 actors.Player) *MockState_Player_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockState_Player_Call) RunAndReturn(run func() actors.Player) *MockState_Player_Call {
	_c.Call.Return(run)
	return _c
}

// Prompter provides a mock function with given fields:
func (_m *MockState) Prompter() ux.PromptLib {
	ret := _m.Called()

	var r0 ux.PromptLib
	if rf, ok := ret.Get(0).(func() ux.PromptLib); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ux.PromptLib)
		}
	}

	return r0
}

// MockState_Prompter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Prompter'
type MockState_Prompter_Call struct {
	*mock.Call
}

// Prompter is a helper method to define mock.On call
func (_e *MockState_Expecter) Prompter() *MockState_Prompter_Call {
	return &MockState_Prompter_Call{Call: _e.mock.On("Prompter")}
}

func (_c *MockState_Prompter_Call) Run(run func()) *MockState_Prompter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockState_Prompter_Call) Return(_a0 ux.PromptLib) *MockState_Prompter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockState_Prompter_Call) RunAndReturn(run func() ux.PromptLib) *MockState_Prompter_Call {
	_c.Call.Return(run)
	return _c
}

// SetPrompter provides a mock function with given fields: p
func (_m *MockState) SetPrompter(p ux.PromptLib) {
	_m.Called(p)
}

// MockState_SetPrompter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPrompter'
type MockState_SetPrompter_Call struct {
	*mock.Call
}

// SetPrompter is a helper method to define mock.On call
//   - p ux.PromptLib
func (_e *MockState_Expecter) SetPrompter(p interface{}) *MockState_SetPrompter_Call {
	return &MockState_SetPrompter_Call{Call: _e.mock.On("SetPrompter", p)}
}

func (_c *MockState_SetPrompter_Call) Run(run func(p ux.PromptLib)) *MockState_SetPrompter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(ux.PromptLib))
	})
	return _c
}

func (_c *MockState_SetPrompter_Call) Return() *MockState_SetPrompter_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockState_SetPrompter_Call) RunAndReturn(run func(ux.PromptLib)) *MockState_SetPrompter_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCurrentRoom provides a mock function with given fields: r
func (_m *MockState) UpdateCurrentRoom(r string) {
	_m.Called(r)
}

// MockState_UpdateCurrentRoom_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCurrentRoom'
type MockState_UpdateCurrentRoom_Call struct {
	*mock.Call
}

// UpdateCurrentRoom is a helper method to define mock.On call
//   - r string
func (_e *MockState_Expecter) UpdateCurrentRoom(r interface{}) *MockState_UpdateCurrentRoom_Call {
	return &MockState_UpdateCurrentRoom_Call{Call: _e.mock.On("UpdateCurrentRoom", r)}
}

func (_c *MockState_UpdateCurrentRoom_Call) Run(run func(r string)) *MockState_UpdateCurrentRoom_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockState_UpdateCurrentRoom_Call) Return() *MockState_UpdateCurrentRoom_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockState_UpdateCurrentRoom_Call) RunAndReturn(run func(string)) *MockState_UpdateCurrentRoom_Call {
	_c.Call.Return(run)
	return _c
}

// World provides a mock function with given fields:
func (_m *MockState) World() *world.WorldDefinition {
	ret := _m.Called()

	var r0 *world.WorldDefinition
	if rf, ok := ret.Get(0).(func() *world.WorldDefinition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*world.WorldDefinition)
		}
	}

	return r0
}

// MockState_World_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'World'
type MockState_World_Call struct {
	*mock.Call
}

// World is a helper method to define mock.On call
func (_e *MockState_Expecter) World() *MockState_World_Call {
	return &MockState_World_Call{Call: _e.mock.On("World")}
}

func (_c *MockState_World_Call) Run(run func()) *MockState_World_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockState_World_Call) Return(_a0 *world.WorldDefinition) *MockState_World_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockState_World_Call) RunAndReturn(run func() *world.WorldDefinition) *MockState_World_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockState creates a new instance of MockState. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockState(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockState {
	mock := &MockState{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
