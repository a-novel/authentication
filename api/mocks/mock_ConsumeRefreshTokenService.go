// Code generated by mockery v2.52.2. DO NOT EDIT.

package apimocks

import (
	context "github.com/a-novel-kit/context"
	mock "github.com/stretchr/testify/mock"

	services "github.com/a-novel/authentication/internal/services"
)

// MockConsumeRefreshTokenService is an autogenerated mock type for the ConsumeRefreshTokenService type
type MockConsumeRefreshTokenService struct {
	mock.Mock
}

type MockConsumeRefreshTokenService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConsumeRefreshTokenService) EXPECT() *MockConsumeRefreshTokenService_Expecter {
	return &MockConsumeRefreshTokenService_Expecter{mock: &_m.Mock}
}

// ConsumeRefreshToken provides a mock function with given fields: ctx, request
func (_m *MockConsumeRefreshTokenService) ConsumeRefreshToken(ctx context.Context, request services.ConsumeRefreshTokenRequest) (string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ConsumeRefreshToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, services.ConsumeRefreshTokenRequest) (string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, services.ConsumeRefreshTokenRequest) string); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, services.ConsumeRefreshTokenRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConsumeRefreshTokenService_ConsumeRefreshToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConsumeRefreshToken'
type MockConsumeRefreshTokenService_ConsumeRefreshToken_Call struct {
	*mock.Call
}

// ConsumeRefreshToken is a helper method to define mock.On call
//   - ctx context.Context
//   - request services.ConsumeRefreshTokenRequest
func (_e *MockConsumeRefreshTokenService_Expecter) ConsumeRefreshToken(ctx interface{}, request interface{}) *MockConsumeRefreshTokenService_ConsumeRefreshToken_Call {
	return &MockConsumeRefreshTokenService_ConsumeRefreshToken_Call{Call: _e.mock.On("ConsumeRefreshToken", ctx, request)}
}

func (_c *MockConsumeRefreshTokenService_ConsumeRefreshToken_Call) Run(run func(ctx context.Context, request services.ConsumeRefreshTokenRequest)) *MockConsumeRefreshTokenService_ConsumeRefreshToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(services.ConsumeRefreshTokenRequest))
	})
	return _c
}

func (_c *MockConsumeRefreshTokenService_ConsumeRefreshToken_Call) Return(_a0 string, _a1 error) *MockConsumeRefreshTokenService_ConsumeRefreshToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumeRefreshTokenService_ConsumeRefreshToken_Call) RunAndReturn(run func(context.Context, services.ConsumeRefreshTokenRequest) (string, error)) *MockConsumeRefreshTokenService_ConsumeRefreshToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConsumeRefreshTokenService creates a new instance of MockConsumeRefreshTokenService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConsumeRefreshTokenService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConsumeRefreshTokenService {
	mock := &MockConsumeRefreshTokenService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
