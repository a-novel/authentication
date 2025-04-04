// Code generated by mockery v2.53.1. DO NOT EDIT.

package apimocks

import (
	context "github.com/a-novel-kit/context"
	mock "github.com/stretchr/testify/mock"

	models "github.com/a-novel/authentication/models"

	services "github.com/a-novel/authentication/internal/services"
)

// MockUpdateRoleService is an autogenerated mock type for the UpdateRoleService type
type MockUpdateRoleService struct {
	mock.Mock
}

type MockUpdateRoleService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUpdateRoleService) EXPECT() *MockUpdateRoleService_Expecter {
	return &MockUpdateRoleService_Expecter{mock: &_m.Mock}
}

// UpdateRole provides a mock function with given fields: ctx, request
func (_m *MockUpdateRoleService) UpdateRole(ctx context.Context, request services.UpdateRoleRequest) (*models.User, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRole")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, services.UpdateRoleRequest) (*models.User, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, services.UpdateRoleRequest) *models.User); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, services.UpdateRoleRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUpdateRoleService_UpdateRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRole'
type MockUpdateRoleService_UpdateRole_Call struct {
	*mock.Call
}

// UpdateRole is a helper method to define mock.On call
//   - ctx context.Context
//   - request services.UpdateRoleRequest
func (_e *MockUpdateRoleService_Expecter) UpdateRole(ctx interface{}, request interface{}) *MockUpdateRoleService_UpdateRole_Call {
	return &MockUpdateRoleService_UpdateRole_Call{Call: _e.mock.On("UpdateRole", ctx, request)}
}

func (_c *MockUpdateRoleService_UpdateRole_Call) Run(run func(ctx context.Context, request services.UpdateRoleRequest)) *MockUpdateRoleService_UpdateRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(services.UpdateRoleRequest))
	})
	return _c
}

func (_c *MockUpdateRoleService_UpdateRole_Call) Return(_a0 *models.User, _a1 error) *MockUpdateRoleService_UpdateRole_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUpdateRoleService_UpdateRole_Call) RunAndReturn(run func(context.Context, services.UpdateRoleRequest) (*models.User, error)) *MockUpdateRoleService_UpdateRole_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUpdateRoleService creates a new instance of MockUpdateRoleService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUpdateRoleService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUpdateRoleService {
	mock := &MockUpdateRoleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
