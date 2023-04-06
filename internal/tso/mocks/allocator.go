// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocktso

import (
	"time"

	"github.com/stretchr/testify/mock"
)

// Allocator is an autogenerated mock type for the Allocator type
type Allocator struct {
	mock.Mock
}

// GenerateTSO provides a mock function with given fields: count
func (_m *Allocator) GenerateTSO(count uint32) (uint64, error) {
	ret := _m.Called(count)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(uint32) uint64); ok {
		r0 = rf(count)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint32) error); ok {
		r1 = rf(count)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastSavedTime provides a mock function with given fields:
func (_m *Allocator) GetLastSavedTime() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// Initialize provides a mock function with given fields:
func (_m *Allocator) Initialize() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reset provides a mock function with given fields:
func (_m *Allocator) Reset() {
	_m.Called()
}

// SetTSO provides a mock function with given fields: _a0
func (_m *Allocator) SetTSO(_a0 uint64) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTSO provides a mock function with given fields:
func (_m *Allocator) UpdateTSO() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewAllocator interface {
	mock.TestingT
	Cleanup(func())
}

// NewAllocator creates a new instance of Allocator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAllocator(t mockConstructorTestingTNewAllocator) *Allocator {
	mock := &Allocator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
