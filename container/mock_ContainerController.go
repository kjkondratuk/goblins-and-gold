// Code generated by mockery v2.20.0. DO NOT EDIT.

package container

import (
	interaction "github.com/kjkondratuk/goblins-and-gold/interaction"
	mock "github.com/stretchr/testify/mock"

	modelcontainer "github.com/kjkondratuk/goblins-and-gold/model/container"

	state "github.com/kjkondratuk/goblins-and-gold/state"
)

// MockContainerController is an autogenerated mock type for the ContainerController type
type MockContainerController struct {
	mock.Mock
}

type MockContainerController_Expecter struct {
	mock *mock.Mock
}

func (_m *MockContainerController) EXPECT() *MockContainerController_Expecter {
	return &MockContainerController_Expecter{mock: &_m.Mock}
}

// Do provides a mock function with given fields: s, c, t
func (_m *MockContainerController) Do(s state.State, c modelcontainer.Container, t interaction.Type) (interaction.Result, error) {
	ret := _m.Called(s, c, t)

	var r0 interaction.Result
	var r1 error
	if rf, ok := ret.Get(0).(func(state.State, modelcontainer.Container, interaction.Type) (interaction.Result, error)); ok {
		return rf(s, c, t)
	}
	if rf, ok := ret.Get(0).(func(state.State, modelcontainer.Container, interaction.Type) interaction.Result); ok {
		r0 = rf(s, c, t)
	} else {
		r0 = ret.Get(0).(interaction.Result)
	}

	if rf, ok := ret.Get(1).(func(state.State, modelcontainer.Container, interaction.Type) error); ok {
		r1 = rf(s, c, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockContainerController_Do_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Do'
type MockContainerController_Do_Call struct {
	*mock.Call
}

// Do is a helper method to define mock.On call
//   - s state.State
//   - c modelcontainer.Container
//   - t interaction.Type
func (_e *MockContainerController_Expecter) Do(s interface{}, c interface{}, t interface{}) *MockContainerController_Do_Call {
	return &MockContainerController_Do_Call{Call: _e.mock.On("Do", s, c, t)}
}

func (_c *MockContainerController_Do_Call) Run(run func(s state.State, c modelcontainer.Container, t interaction.Type)) *MockContainerController_Do_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.State), args[1].(modelcontainer.Container), args[2].(interaction.Type))
	})
	return _c
}

func (_c *MockContainerController_Do_Call) Return(_a0 interaction.Result, _a1 error) *MockContainerController_Do_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockContainerController_Do_Call) RunAndReturn(run func(state.State, modelcontainer.Container, interaction.Type) (interaction.Result, error)) *MockContainerController_Do_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockContainerController interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockContainerController creates a new instance of MockContainerController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockContainerController(t mockConstructorTestingTNewMockContainerController) *MockContainerController {
	mock := &MockContainerController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
