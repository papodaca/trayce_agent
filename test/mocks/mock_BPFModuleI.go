// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	libbpfgo "github.com/aquasecurity/libbpfgo"
	mock "github.com/stretchr/testify/mock"
)

// MockBPFModuleI is an autogenerated mock type for the BPFModuleI type
type MockBPFModuleI struct {
	mock.Mock
}

type MockBPFModuleI_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBPFModuleI) EXPECT() *MockBPFModuleI_Expecter {
	return &MockBPFModuleI_Expecter{mock: &_m.Mock}
}

// BPFLoadObject provides a mock function with given fields:
func (_m *MockBPFModuleI) BPFLoadObject() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for BPFLoadObject")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBPFModuleI_BPFLoadObject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BPFLoadObject'
type MockBPFModuleI_BPFLoadObject_Call struct {
	*mock.Call
}

// BPFLoadObject is a helper method to define mock.On call
func (_e *MockBPFModuleI_Expecter) BPFLoadObject() *MockBPFModuleI_BPFLoadObject_Call {
	return &MockBPFModuleI_BPFLoadObject_Call{Call: _e.mock.On("BPFLoadObject")}
}

func (_c *MockBPFModuleI_BPFLoadObject_Call) Run(run func()) *MockBPFModuleI_BPFLoadObject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBPFModuleI_BPFLoadObject_Call) Return(_a0 error) *MockBPFModuleI_BPFLoadObject_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBPFModuleI_BPFLoadObject_Call) RunAndReturn(run func() error) *MockBPFModuleI_BPFLoadObject_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockBPFModuleI) Close() {
	_m.Called()
}

// MockBPFModuleI_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockBPFModuleI_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockBPFModuleI_Expecter) Close() *MockBPFModuleI_Close_Call {
	return &MockBPFModuleI_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockBPFModuleI_Close_Call) Run(run func()) *MockBPFModuleI_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBPFModuleI_Close_Call) Return() *MockBPFModuleI_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBPFModuleI_Close_Call) RunAndReturn(run func()) *MockBPFModuleI_Close_Call {
	_c.Call.Return(run)
	return _c
}

// GetMap provides a mock function with given fields: mapName
func (_m *MockBPFModuleI) GetMap(mapName string) (*libbpfgo.BPFMap, error) {
	ret := _m.Called(mapName)

	if len(ret) == 0 {
		panic("no return value specified for GetMap")
	}

	var r0 *libbpfgo.BPFMap
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*libbpfgo.BPFMap, error)); ok {
		return rf(mapName)
	}
	if rf, ok := ret.Get(0).(func(string) *libbpfgo.BPFMap); ok {
		r0 = rf(mapName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*libbpfgo.BPFMap)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(mapName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBPFModuleI_GetMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMap'
type MockBPFModuleI_GetMap_Call struct {
	*mock.Call
}

// GetMap is a helper method to define mock.On call
//   - mapName string
func (_e *MockBPFModuleI_Expecter) GetMap(mapName interface{}) *MockBPFModuleI_GetMap_Call {
	return &MockBPFModuleI_GetMap_Call{Call: _e.mock.On("GetMap", mapName)}
}

func (_c *MockBPFModuleI_GetMap_Call) Run(run func(mapName string)) *MockBPFModuleI_GetMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockBPFModuleI_GetMap_Call) Return(_a0 *libbpfgo.BPFMap, _a1 error) *MockBPFModuleI_GetMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBPFModuleI_GetMap_Call) RunAndReturn(run func(string) (*libbpfgo.BPFMap, error)) *MockBPFModuleI_GetMap_Call {
	_c.Call.Return(run)
	return _c
}

// GetProgram provides a mock function with given fields: progName
func (_m *MockBPFModuleI) GetProgram(progName string) (*libbpfgo.BPFProg, error) {
	ret := _m.Called(progName)

	if len(ret) == 0 {
		panic("no return value specified for GetProgram")
	}

	var r0 *libbpfgo.BPFProg
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*libbpfgo.BPFProg, error)); ok {
		return rf(progName)
	}
	if rf, ok := ret.Get(0).(func(string) *libbpfgo.BPFProg); ok {
		r0 = rf(progName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*libbpfgo.BPFProg)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(progName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBPFModuleI_GetProgram_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProgram'
type MockBPFModuleI_GetProgram_Call struct {
	*mock.Call
}

// GetProgram is a helper method to define mock.On call
//   - progName string
func (_e *MockBPFModuleI_Expecter) GetProgram(progName interface{}) *MockBPFModuleI_GetProgram_Call {
	return &MockBPFModuleI_GetProgram_Call{Call: _e.mock.On("GetProgram", progName)}
}

func (_c *MockBPFModuleI_GetProgram_Call) Run(run func(progName string)) *MockBPFModuleI_GetProgram_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockBPFModuleI_GetProgram_Call) Return(_a0 *libbpfgo.BPFProg, _a1 error) *MockBPFModuleI_GetProgram_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBPFModuleI_GetProgram_Call) RunAndReturn(run func(string) (*libbpfgo.BPFProg, error)) *MockBPFModuleI_GetProgram_Call {
	_c.Call.Return(run)
	return _c
}

// InitRingBuf provides a mock function with given fields: mapName, eventsChan
func (_m *MockBPFModuleI) InitRingBuf(mapName string, eventsChan chan []byte) (*libbpfgo.RingBuffer, error) {
	ret := _m.Called(mapName, eventsChan)

	if len(ret) == 0 {
		panic("no return value specified for InitRingBuf")
	}

	var r0 *libbpfgo.RingBuffer
	var r1 error
	if rf, ok := ret.Get(0).(func(string, chan []byte) (*libbpfgo.RingBuffer, error)); ok {
		return rf(mapName, eventsChan)
	}
	if rf, ok := ret.Get(0).(func(string, chan []byte) *libbpfgo.RingBuffer); ok {
		r0 = rf(mapName, eventsChan)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*libbpfgo.RingBuffer)
		}
	}

	if rf, ok := ret.Get(1).(func(string, chan []byte) error); ok {
		r1 = rf(mapName, eventsChan)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBPFModuleI_InitRingBuf_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitRingBuf'
type MockBPFModuleI_InitRingBuf_Call struct {
	*mock.Call
}

// InitRingBuf is a helper method to define mock.On call
//   - mapName string
//   - eventsChan chan []byte
func (_e *MockBPFModuleI_Expecter) InitRingBuf(mapName interface{}, eventsChan interface{}) *MockBPFModuleI_InitRingBuf_Call {
	return &MockBPFModuleI_InitRingBuf_Call{Call: _e.mock.On("InitRingBuf", mapName, eventsChan)}
}

func (_c *MockBPFModuleI_InitRingBuf_Call) Run(run func(mapName string, eventsChan chan []byte)) *MockBPFModuleI_InitRingBuf_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(chan []byte))
	})
	return _c
}

func (_c *MockBPFModuleI_InitRingBuf_Call) Return(_a0 *libbpfgo.RingBuffer, _a1 error) *MockBPFModuleI_InitRingBuf_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBPFModuleI_InitRingBuf_Call) RunAndReturn(run func(string, chan []byte) (*libbpfgo.RingBuffer, error)) *MockBPFModuleI_InitRingBuf_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBPFModuleI creates a new instance of MockBPFModuleI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBPFModuleI(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBPFModuleI {
	mock := &MockBPFModuleI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
