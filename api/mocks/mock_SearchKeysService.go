// Code generated by mockery v2.53.1. DO NOT EDIT.

package apimocks

import (
	context "github.com/a-novel-kit/context"
	jwa "github.com/a-novel-kit/jwt/jwa"

	mock "github.com/stretchr/testify/mock"

	services "github.com/a-novel/authentication/internal/services"
)

// MockSearchKeysService is an autogenerated mock type for the SearchKeysService type
type MockSearchKeysService struct {
	mock.Mock
}

type MockSearchKeysService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSearchKeysService) EXPECT() *MockSearchKeysService_Expecter {
	return &MockSearchKeysService_Expecter{mock: &_m.Mock}
}

// SearchKeys provides a mock function with given fields: ctx, request
func (_m *MockSearchKeysService) SearchKeys(ctx context.Context, request services.SearchKeysRequest) ([]*jwa.JWK, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for SearchKeys")
	}

	var r0 []*jwa.JWK
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, services.SearchKeysRequest) ([]*jwa.JWK, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, services.SearchKeysRequest) []*jwa.JWK); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*jwa.JWK)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, services.SearchKeysRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSearchKeysService_SearchKeys_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SearchKeys'
type MockSearchKeysService_SearchKeys_Call struct {
	*mock.Call
}

// SearchKeys is a helper method to define mock.On call
//   - ctx context.Context
//   - request services.SearchKeysRequest
func (_e *MockSearchKeysService_Expecter) SearchKeys(ctx interface{}, request interface{}) *MockSearchKeysService_SearchKeys_Call {
	return &MockSearchKeysService_SearchKeys_Call{Call: _e.mock.On("SearchKeys", ctx, request)}
}

func (_c *MockSearchKeysService_SearchKeys_Call) Run(run func(ctx context.Context, request services.SearchKeysRequest)) *MockSearchKeysService_SearchKeys_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(services.SearchKeysRequest))
	})
	return _c
}

func (_c *MockSearchKeysService_SearchKeys_Call) Return(_a0 []*jwa.JWK, _a1 error) *MockSearchKeysService_SearchKeys_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSearchKeysService_SearchKeys_Call) RunAndReturn(run func(context.Context, services.SearchKeysRequest) ([]*jwa.JWK, error)) *MockSearchKeysService_SearchKeys_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSearchKeysService creates a new instance of MockSearchKeysService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSearchKeysService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSearchKeysService {
	mock := &MockSearchKeysService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
