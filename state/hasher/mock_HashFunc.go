// Code generated by mockery v2.34.2. DO NOT EDIT.

package hasher

import (
	actors "github.com/kjkondratuk/goblins-and-gold/actors"
	mock "github.com/stretchr/testify/mock"
)

// MockHashFunc is an autogenerated mock type for the HashFunc type
type MockHashFunc struct {
	mock.Mock
}

type MockHashFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHashFunc) EXPECT() *MockHashFunc_Expecter {
	return &MockHashFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: combatant, unique
func (_m *MockHashFunc) Execute(combatant actors.Combatant, unique int) string {
	ret := _m.Called(combatant, unique)

	var r0 string
	if rf, ok := ret.Get(0).(func(actors.Combatant, int) string); ok {
		r0 = rf(combatant, unique)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockHashFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockHashFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - combatant actors.Combatant
//   - unique int
func (_e *MockHashFunc_Expecter) Execute(combatant interface{}, unique interface{}) *MockHashFunc_Execute_Call {
	return &MockHashFunc_Execute_Call{Call: _e.mock.On("Execute", combatant, unique)}
}

func (_c *MockHashFunc_Execute_Call) Run(run func(combatant actors.Combatant, unique int)) *MockHashFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(actors.Combatant), args[1].(int))
	})
	return _c
}

func (_c *MockHashFunc_Execute_Call) Return(_a0 string) *MockHashFunc_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHashFunc_Execute_Call) RunAndReturn(run func(actors.Combatant, int) string) *MockHashFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHashFunc creates a new instance of MockHashFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHashFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHashFunc {
	mock := &MockHashFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
