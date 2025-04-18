// Code generated by mockery. DO NOT EDIT.

package v1

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// HttpRequestDoer is an autogenerated mock type for the HttpRequestDoer type
type HttpRequestDoer struct {
	mock.Mock
}

type HttpRequestDoer_Expecter struct {
	mock *mock.Mock
}

func (_m *HttpRequestDoer) EXPECT() *HttpRequestDoer_Expecter {
	return &HttpRequestDoer_Expecter{mock: &_m.Mock}
}

// Do provides a mock function with given fields: req
func (_m *HttpRequestDoer) Do(req *http.Request) (*http.Response, error) {
	ret := _m.Called(req)

	if len(ret) == 0 {
		panic("no return value specified for Do")
	}

	var r0 *http.Response
	var r1 error
	if rf, ok := ret.Get(0).(func(*http.Request) (*http.Response, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*http.Request) *http.Response); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	if rf, ok := ret.Get(1).(func(*http.Request) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HttpRequestDoer_Do_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Do'
type HttpRequestDoer_Do_Call struct {
	*mock.Call
}

// Do is a helper method to define mock.On call
//   - req *http.Request
func (_e *HttpRequestDoer_Expecter) Do(req interface{}) *HttpRequestDoer_Do_Call {
	return &HttpRequestDoer_Do_Call{Call: _e.mock.On("Do", req)}
}

func (_c *HttpRequestDoer_Do_Call) Run(run func(req *http.Request)) *HttpRequestDoer_Do_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*http.Request))
	})
	return _c
}

func (_c *HttpRequestDoer_Do_Call) Return(_a0 *http.Response, _a1 error) *HttpRequestDoer_Do_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *HttpRequestDoer_Do_Call) RunAndReturn(run func(*http.Request) (*http.Response, error)) *HttpRequestDoer_Do_Call {
	_c.Call.Return(run)
	return _c
}

// NewHttpRequestDoer creates a new instance of HttpRequestDoer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHttpRequestDoer(t interface {
	mock.TestingT
	Cleanup(func())
}) *HttpRequestDoer {
	mock := &HttpRequestDoer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
