// Code generated by mockery v2.53.1. DO NOT EDIT.

package apimocks

import (
	context "github.com/a-novel-kit/context"
	mock "github.com/stretchr/testify/mock"

	services "github.com/a-novel/authentication/internal/services"
)

// MockLoginService is an autogenerated mock type for the LoginService type
type MockLoginService struct {
	mock.Mock
}

type MockLoginService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockLoginService) EXPECT() *MockLoginService_Expecter {
	return &MockLoginService_Expecter{mock: &_m.Mock}
}

// Login provides a mock function with given fields: ctx, request
func (_m *MockLoginService) Login(ctx context.Context, request services.LoginRequest) (string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, services.LoginRequest) (string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, services.LoginRequest) string); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, services.LoginRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockLoginService_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type MockLoginService_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - ctx context.Context
//   - request services.LoginRequest
func (_e *MockLoginService_Expecter) Login(ctx interface{}, request interface{}) *MockLoginService_Login_Call {
	return &MockLoginService_Login_Call{Call: _e.mock.On("Login", ctx, request)}
}

func (_c *MockLoginService_Login_Call) Run(run func(ctx context.Context, request services.LoginRequest)) *MockLoginService_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(services.LoginRequest))
	})
	return _c
}

func (_c *MockLoginService_Login_Call) Return(_a0 string, _a1 error) *MockLoginService_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockLoginService_Login_Call) RunAndReturn(run func(context.Context, services.LoginRequest) (string, error)) *MockLoginService_Login_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockLoginService creates a new instance of MockLoginService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockLoginService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockLoginService {
	mock := &MockLoginService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
