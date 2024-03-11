// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "gitlab.com/beabys/go-http-template/internal/api/v1"
)

// ClientOption is an autogenerated mock type for the ClientOption type
type ClientOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *ClientOption) Execute(_a0 *v1.Client) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*v1.Client) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewClientOption creates a new instance of ClientOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientOption {
	mock := &ClientOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}