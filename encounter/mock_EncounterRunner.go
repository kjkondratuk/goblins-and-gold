// Code generated by mockery v2.34.2. DO NOT EDIT.

package encounter

import (
	modelencounter "github.com/kjkondratuk/goblins-and-gold/model/encounter"
	state "github.com/kjkondratuk/goblins-and-gold/state"
	mock "github.com/stretchr/testify/mock"
)

// MockEncounterRunner is an autogenerated mock type for the EncounterRunner type
type MockEncounterRunner struct {
	mock.Mock
}

type MockEncounterRunner_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEncounterRunner) EXPECT() *MockEncounterRunner_Expecter {
	return &MockEncounterRunner_Expecter{mock: &_m.Mock}
}

// Run provides a mock function with given fields: s, e
func (_m *MockEncounterRunner) Run(s state.State, e modelencounter.Encounter) Outcome {
	ret := _m.Called(s, e)

	var r0 Outcome
	if rf, ok := ret.Get(0).(func(state.State, modelencounter.Encounter) Outcome); ok {
		r0 = rf(s, e)
	} else {
		r0 = ret.Get(0).(Outcome)
	}

	return r0
}

// MockEncounterRunner_Run_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Run'
type MockEncounterRunner_Run_Call struct {
	*mock.Call
}

// Run is a helper method to define mock.On call
//   - s state.State
//   - e modelencounter.Encounter
func (_e *MockEncounterRunner_Expecter) Run(s interface{}, e interface{}) *MockEncounterRunner_Run_Call {
	return &MockEncounterRunner_Run_Call{Call: _e.mock.On("Run", s, e)}
}

func (_c *MockEncounterRunner_Run_Call) Run(run func(s state.State, e modelencounter.Encounter)) *MockEncounterRunner_Run_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(state.State), args[1].(modelencounter.Encounter))
	})
	return _c
}

func (_c *MockEncounterRunner_Run_Call) Return(_a0 Outcome) *MockEncounterRunner_Run_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEncounterRunner_Run_Call) RunAndReturn(run func(state.State, modelencounter.Encounter) Outcome) *MockEncounterRunner_Run_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEncounterRunner creates a new instance of MockEncounterRunner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEncounterRunner(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEncounterRunner {
	mock := &MockEncounterRunner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
