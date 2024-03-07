// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	v1 "gitlab.com/beabys/go-http-template/internal/api/v1"
)

// ServerOption is an autogenerated mock type for the ServerOption type
type ServerOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *ServerOption) Execute(_a0 *v1.ServerOptions) {
	_m.Called(_a0)
}

// NewServerOption creates a new instance of ServerOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServerOption(t interface {
	mock.TestingT
	Cleanup(func())
}) *ServerOption {
	mock := &ServerOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}