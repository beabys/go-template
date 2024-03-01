// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	v1 "gitlab.com/beabys/go-http-template/internal/api/v1"
)

// ClientInterface is an autogenerated mock type for the ClientInterface type
type ClientInterface struct {
	mock.Mock
}

// HelloWorld provides a mock function with given fields: ctx, reqEditors
func (_m *ClientInterface) HelloWorld(ctx context.Context, reqEditors ...v1.RequestEditorFn) (*http.Response, error) {
	_va := make([]interface{}, len(reqEditors))
	for _i := range reqEditors {
		_va[_i] = reqEditors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for HelloWorld")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...v1.RequestEditorFn) (*http.Response, error)); ok {
		return rf(ctx, reqEditors...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...v1.RequestEditorFn) *http.Response); ok {
		r0 = rf(ctx, reqEditors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...v1.RequestEditorFn) error); ok {
		r1 = rf(ctx, reqEditors...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewClientInterface creates a new instance of ClientInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClientInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ClientInterface {
	mock := &ClientInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
