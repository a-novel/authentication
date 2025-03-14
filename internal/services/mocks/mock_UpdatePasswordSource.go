// Code generated by mockery v2.53.1. DO NOT EDIT.

package servicesmocks

import (
	context "github.com/a-novel-kit/context"
	dao "github.com/a-novel/authentication/internal/dao"
	mock "github.com/stretchr/testify/mock"

	models "github.com/a-novel/authentication/models"

	services "github.com/a-novel/authentication/internal/services"

	uuid "github.com/google/uuid"
)

// MockUpdatePasswordSource is an autogenerated mock type for the UpdatePasswordSource type
type MockUpdatePasswordSource struct {
	mock.Mock
}

type MockUpdatePasswordSource_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUpdatePasswordSource) EXPECT() *MockUpdatePasswordSource_Expecter {
	return &MockUpdatePasswordSource_Expecter{mock: &_m.Mock}
}

// ConsumeShortCode provides a mock function with given fields: ctx, request
func (_m *MockUpdatePasswordSource) ConsumeShortCode(ctx context.Context, request services.ConsumeShortCodeRequest) (*models.ShortCode, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ConsumeShortCode")
	}

	var r0 *models.ShortCode
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, services.ConsumeShortCodeRequest) (*models.ShortCode, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, services.ConsumeShortCodeRequest) *models.ShortCode); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ShortCode)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, services.ConsumeShortCodeRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUpdatePasswordSource_ConsumeShortCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConsumeShortCode'
type MockUpdatePasswordSource_ConsumeShortCode_Call struct {
	*mock.Call
}

// ConsumeShortCode is a helper method to define mock.On call
//   - ctx context.Context
//   - request services.ConsumeShortCodeRequest
func (_e *MockUpdatePasswordSource_Expecter) ConsumeShortCode(ctx interface{}, request interface{}) *MockUpdatePasswordSource_ConsumeShortCode_Call {
	return &MockUpdatePasswordSource_ConsumeShortCode_Call{Call: _e.mock.On("ConsumeShortCode", ctx, request)}
}

func (_c *MockUpdatePasswordSource_ConsumeShortCode_Call) Run(run func(ctx context.Context, request services.ConsumeShortCodeRequest)) *MockUpdatePasswordSource_ConsumeShortCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(services.ConsumeShortCodeRequest))
	})
	return _c
}

func (_c *MockUpdatePasswordSource_ConsumeShortCode_Call) Return(_a0 *models.ShortCode, _a1 error) *MockUpdatePasswordSource_ConsumeShortCode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUpdatePasswordSource_ConsumeShortCode_Call) RunAndReturn(run func(context.Context, services.ConsumeShortCodeRequest) (*models.ShortCode, error)) *MockUpdatePasswordSource_ConsumeShortCode_Call {
	_c.Call.Return(run)
	return _c
}

// SelectCredentials provides a mock function with given fields: ctx, id
func (_m *MockUpdatePasswordSource) SelectCredentials(ctx context.Context, id uuid.UUID) (*dao.CredentialsEntity, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for SelectCredentials")
	}

	var r0 *dao.CredentialsEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*dao.CredentialsEntity, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *dao.CredentialsEntity); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.CredentialsEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUpdatePasswordSource_SelectCredentials_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SelectCredentials'
type MockUpdatePasswordSource_SelectCredentials_Call struct {
	*mock.Call
}

// SelectCredentials is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *MockUpdatePasswordSource_Expecter) SelectCredentials(ctx interface{}, id interface{}) *MockUpdatePasswordSource_SelectCredentials_Call {
	return &MockUpdatePasswordSource_SelectCredentials_Call{Call: _e.mock.On("SelectCredentials", ctx, id)}
}

func (_c *MockUpdatePasswordSource_SelectCredentials_Call) Run(run func(ctx context.Context, id uuid.UUID)) *MockUpdatePasswordSource_SelectCredentials_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockUpdatePasswordSource_SelectCredentials_Call) Return(_a0 *dao.CredentialsEntity, _a1 error) *MockUpdatePasswordSource_SelectCredentials_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUpdatePasswordSource_SelectCredentials_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*dao.CredentialsEntity, error)) *MockUpdatePasswordSource_SelectCredentials_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateCredentialsPassword provides a mock function with given fields: ctx, userID, data
func (_m *MockUpdatePasswordSource) UpdateCredentialsPassword(ctx context.Context, userID uuid.UUID, data dao.UpdateCredentialsPasswordData) (*dao.CredentialsEntity, error) {
	ret := _m.Called(ctx, userID, data)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCredentialsPassword")
	}

	var r0 *dao.CredentialsEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, dao.UpdateCredentialsPasswordData) (*dao.CredentialsEntity, error)); ok {
		return rf(ctx, userID, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, dao.UpdateCredentialsPasswordData) *dao.CredentialsEntity); ok {
		r0 = rf(ctx, userID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dao.CredentialsEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, dao.UpdateCredentialsPasswordData) error); ok {
		r1 = rf(ctx, userID, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUpdatePasswordSource_UpdateCredentialsPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateCredentialsPassword'
type MockUpdatePasswordSource_UpdateCredentialsPassword_Call struct {
	*mock.Call
}

// UpdateCredentialsPassword is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
//   - data dao.UpdateCredentialsPasswordData
func (_e *MockUpdatePasswordSource_Expecter) UpdateCredentialsPassword(ctx interface{}, userID interface{}, data interface{}) *MockUpdatePasswordSource_UpdateCredentialsPassword_Call {
	return &MockUpdatePasswordSource_UpdateCredentialsPassword_Call{Call: _e.mock.On("UpdateCredentialsPassword", ctx, userID, data)}
}

func (_c *MockUpdatePasswordSource_UpdateCredentialsPassword_Call) Run(run func(ctx context.Context, userID uuid.UUID, data dao.UpdateCredentialsPasswordData)) *MockUpdatePasswordSource_UpdateCredentialsPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(dao.UpdateCredentialsPasswordData))
	})
	return _c
}

func (_c *MockUpdatePasswordSource_UpdateCredentialsPassword_Call) Return(_a0 *dao.CredentialsEntity, _a1 error) *MockUpdatePasswordSource_UpdateCredentialsPassword_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUpdatePasswordSource_UpdateCredentialsPassword_Call) RunAndReturn(run func(context.Context, uuid.UUID, dao.UpdateCredentialsPasswordData) (*dao.CredentialsEntity, error)) *MockUpdatePasswordSource_UpdateCredentialsPassword_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUpdatePasswordSource creates a new instance of MockUpdatePasswordSource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUpdatePasswordSource(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUpdatePasswordSource {
	mock := &MockUpdatePasswordSource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
