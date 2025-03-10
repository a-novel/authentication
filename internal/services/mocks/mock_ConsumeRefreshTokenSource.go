// Code generated by mockery v2.53.1. DO NOT EDIT.

package servicesmocks

import (
	context "github.com/a-novel-kit/context"
	mock "github.com/stretchr/testify/mock"

	services "github.com/a-novel/authentication/internal/services"
)

// MockConsumeRefreshTokenSource is an autogenerated mock type for the ConsumeRefreshTokenSource type
type MockConsumeRefreshTokenSource struct {
	mock.Mock
}

type MockConsumeRefreshTokenSource_Expecter struct {
	mock *mock.Mock
}

func (_m *MockConsumeRefreshTokenSource) EXPECT() *MockConsumeRefreshTokenSource_Expecter {
	return &MockConsumeRefreshTokenSource_Expecter{mock: &_m.Mock}
}

// IssueToken provides a mock function with given fields: ctx, request
func (_m *MockConsumeRefreshTokenSource) IssueToken(ctx context.Context, request services.IssueTokenRequest) (string, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for IssueToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, services.IssueTokenRequest) (string, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, services.IssueTokenRequest) string); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, services.IssueTokenRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockConsumeRefreshTokenSource_IssueToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IssueToken'
type MockConsumeRefreshTokenSource_IssueToken_Call struct {
	*mock.Call
}

// IssueToken is a helper method to define mock.On call
//   - ctx context.Context
//   - request services.IssueTokenRequest
func (_e *MockConsumeRefreshTokenSource_Expecter) IssueToken(ctx interface{}, request interface{}) *MockConsumeRefreshTokenSource_IssueToken_Call {
	return &MockConsumeRefreshTokenSource_IssueToken_Call{Call: _e.mock.On("IssueToken", ctx, request)}
}

func (_c *MockConsumeRefreshTokenSource_IssueToken_Call) Run(run func(ctx context.Context, request services.IssueTokenRequest)) *MockConsumeRefreshTokenSource_IssueToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(services.IssueTokenRequest))
	})
	return _c
}

func (_c *MockConsumeRefreshTokenSource_IssueToken_Call) Return(_a0 string, _a1 error) *MockConsumeRefreshTokenSource_IssueToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockConsumeRefreshTokenSource_IssueToken_Call) RunAndReturn(run func(context.Context, services.IssueTokenRequest) (string, error)) *MockConsumeRefreshTokenSource_IssueToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockConsumeRefreshTokenSource creates a new instance of MockConsumeRefreshTokenSource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockConsumeRefreshTokenSource(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockConsumeRefreshTokenSource {
	mock := &MockConsumeRefreshTokenSource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
