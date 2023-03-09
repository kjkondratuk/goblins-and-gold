// Code generated by mockery v2.20.0. DO NOT EDIT.

package ux

import mock "github.com/stretchr/testify/mock"

// MockPromptLib is an autogenerated mock type for the PromptLib type
type MockPromptLib struct {
	mock.Mock
}

type MockPromptLib_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPromptLib) EXPECT() *MockPromptLib_Expecter {
	return &MockPromptLib_Expecter{mock: &_m.Mock}
}

// Prompt provides a mock function with given fields: label, defaultValue
func (_m *MockPromptLib) Prompt(label string, defaultValue string) (string, error) {
	ret := _m.Called(label, defaultValue)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(label, defaultValue)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(label, defaultValue)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(label, defaultValue)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPromptLib_Prompt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Prompt'
type MockPromptLib_Prompt_Call struct {
	*mock.Call
}

// Prompt is a helper method to define mock.On call
//   - label string
//   - defaultValue string
func (_e *MockPromptLib_Expecter) Prompt(label interface{}, defaultValue interface{}) *MockPromptLib_Prompt_Call {
	return &MockPromptLib_Prompt_Call{Call: _e.mock.On("Prompt", label, defaultValue)}
}

func (_c *MockPromptLib_Prompt_Call) Run(run func(label string, defaultValue string)) *MockPromptLib_Prompt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *MockPromptLib_Prompt_Call) Return(_a0 string, _a1 error) *MockPromptLib_Prompt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPromptLib_Prompt_Call) RunAndReturn(run func(string, string) (string, error)) *MockPromptLib_Prompt_Call {
	_c.Call.Return(run)
	return _c
}

// Select provides a mock function with given fields: label, items
func (_m *MockPromptLib) Select(label string, items []string) (int, string, error) {
	ret := _m.Called(label, items)

	var r0 int
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, []string) (int, string, error)); ok {
		return rf(label, items)
	}
	if rf, ok := ret.Get(0).(func(string, []string) int); ok {
		r0 = rf(label, items)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(string, []string) string); ok {
		r1 = rf(label, items)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, []string) error); ok {
		r2 = rf(label, items)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockPromptLib_Select_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Select'
type MockPromptLib_Select_Call struct {
	*mock.Call
}

// Select is a helper method to define mock.On call
//   - label string
//   - items []string
func (_e *MockPromptLib_Expecter) Select(label interface{}, items interface{}) *MockPromptLib_Select_Call {
	return &MockPromptLib_Select_Call{Call: _e.mock.On("Select", label, items)}
}

func (_c *MockPromptLib_Select_Call) Run(run func(label string, items []string)) *MockPromptLib_Select_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].([]string))
	})
	return _c
}

func (_c *MockPromptLib_Select_Call) Return(_a0 int, _a1 string, _a2 error) *MockPromptLib_Select_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockPromptLib_Select_Call) RunAndReturn(run func(string, []string) (int, string, error)) *MockPromptLib_Select_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPromptLib interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPromptLib creates a new instance of MockPromptLib. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPromptLib(t mockConstructorTestingTNewMockPromptLib) *MockPromptLib {
	mock := &MockPromptLib{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
