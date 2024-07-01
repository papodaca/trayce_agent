// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	docker "github.com/evanrolfe/trayce_agent/internal/docker"
	mock "github.com/stretchr/testify/mock"
)

// MockContainersI is an autogenerated mock type for the ContainersI type
type MockContainersI struct {
	mock.Mock
}

type MockContainersI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockContainersI) EXPECT() *MockContainersI_Expecter {
	return &MockContainersI_Expecter{mock: &_m.Mock}
}

// GetContainersToIntercept provides a mock function with given fields:
func (_m *MockContainersI) GetContainersToIntercept() map[string]docker.Container {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetContainersToIntercept")
	}

	var r0 map[string]docker.Container
	if rf, ok := ret.Get(0).(func() map[string]docker.Container); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]docker.Container)
		}
	}

	return r0
}

// MockContainersI_GetContainersToIntercept_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContainersToIntercept'
type MockContainersI_GetContainersToIntercept_Call struct {
	*mock.Call
}

// GetContainersToIntercept is a helper method to define mock.On call
func (_e *MockContainersI_Expecter) GetContainersToIntercept() *MockContainersI_GetContainersToIntercept_Call {
	return &MockContainersI_GetContainersToIntercept_Call{Call: _e.mock.On("GetContainersToIntercept")}
}

func (_c *MockContainersI_GetContainersToIntercept_Call) Run(run func()) *MockContainersI_GetContainersToIntercept_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContainersI_GetContainersToIntercept_Call) Return(_a0 map[string]docker.Container) *MockContainersI_GetContainersToIntercept_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContainersI_GetContainersToIntercept_Call) RunAndReturn(run func() map[string]docker.Container) *MockContainersI_GetContainersToIntercept_Call {
	_c.Call.Return(run)
	return _c
}

// GetProcsToIntercept provides a mock function with given fields:
func (_m *MockContainersI) GetProcsToIntercept() map[uint32]docker.Proc {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetProcsToIntercept")
	}

	var r0 map[uint32]docker.Proc
	if rf, ok := ret.Get(0).(func() map[uint32]docker.Proc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32]docker.Proc)
		}
	}

	return r0
}

// MockContainersI_GetProcsToIntercept_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProcsToIntercept'
type MockContainersI_GetProcsToIntercept_Call struct {
	*mock.Call
}

// GetProcsToIntercept is a helper method to define mock.On call
func (_e *MockContainersI_Expecter) GetProcsToIntercept() *MockContainersI_GetProcsToIntercept_Call {
	return &MockContainersI_GetProcsToIntercept_Call{Call: _e.mock.On("GetProcsToIntercept")}
}

func (_c *MockContainersI_GetProcsToIntercept_Call) Run(run func()) *MockContainersI_GetProcsToIntercept_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockContainersI_GetProcsToIntercept_Call) Return(_a0 map[uint32]docker.Proc) *MockContainersI_GetProcsToIntercept_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockContainersI_GetProcsToIntercept_Call) RunAndReturn(run func() map[uint32]docker.Proc) *MockContainersI_GetProcsToIntercept_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockContainersI creates a new instance of MockContainersI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockContainersI(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockContainersI {
	mock := &MockContainersI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
