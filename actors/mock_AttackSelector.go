// Code generated by mockery v2.34.2. DO NOT EDIT.

package actors

import (
	attack "github.com/kjkondratuk/goblins-and-gold/model/attack"
	mock "github.com/stretchr/testify/mock"
)

// MockAttackSelector is an autogenerated mock type for the AttackSelector type
type MockAttackSelector struct {
	mock.Mock
}

type MockAttackSelector_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAttackSelector) EXPECT() *MockAttackSelector_Expecter {
	return &MockAttackSelector_Expecter{mock: &_m.Mock}
}

// Select provides a mock function with given fields: c
func (_m *MockAttackSelector) Select(c Combatant) (string, attack.Attack) {
	ret := _m.Called(c)

	var r0 string
	var r1 attack.Attack
	if rf, ok := ret.Get(0).(func(Combatant) (string, attack.Attack)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(Combatant) string); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(Combatant) attack.Attack); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Get(1).(attack.Attack)
	}

	return r0, r1
}

// MockAttackSelector_Select_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Select'
type MockAttackSelector_Select_Call struct {
	*mock.Call
}

// Select is a helper method to define mock.On call
//   - c Combatant
func (_e *MockAttackSelector_Expecter) Select(c interface{}) *MockAttackSelector_Select_Call {
	return &MockAttackSelector_Select_Call{Call: _e.mock.On("Select", c)}
}

func (_c *MockAttackSelector_Select_Call) Run(run func(c Combatant)) *MockAttackSelector_Select_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(Combatant))
	})
	return _c
}

func (_c *MockAttackSelector_Select_Call) Return(_a0 string, _a1 attack.Attack) *MockAttackSelector_Select_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAttackSelector_Select_Call) RunAndReturn(run func(Combatant) (string, attack.Attack)) *MockAttackSelector_Select_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAttackSelector creates a new instance of MockAttackSelector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAttackSelector(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAttackSelector {
	mock := &MockAttackSelector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}