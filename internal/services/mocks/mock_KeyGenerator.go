// Code generated by mockery v2.52.2. DO NOT EDIT.

package servicesmocks

import (
	jwa "github.com/a-novel-kit/jwt/jwa"
	mock "github.com/stretchr/testify/mock"
)

// MockKeyGenerator is an autogenerated mock type for the KeyGenerator type
type MockKeyGenerator struct {
	mock.Mock
}

type MockKeyGenerator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockKeyGenerator) EXPECT() *MockKeyGenerator_Expecter {
	return &MockKeyGenerator_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with no fields
func (_m *MockKeyGenerator) Execute() (*jwa.JWK, *jwa.JWK, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *jwa.JWK
	var r1 *jwa.JWK
	var r2 error
	if rf, ok := ret.Get(0).(func() (*jwa.JWK, *jwa.JWK, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *jwa.JWK); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwa.JWK)
		}
	}

	if rf, ok := ret.Get(1).(func() *jwa.JWK); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*jwa.JWK)
		}
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockKeyGenerator_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockKeyGenerator_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *MockKeyGenerator_Expecter) Execute() *MockKeyGenerator_Execute_Call {
	return &MockKeyGenerator_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *MockKeyGenerator_Execute_Call) Run(run func()) *MockKeyGenerator_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockKeyGenerator_Execute_Call) Return(privateKey *jwa.JWK, publicKey *jwa.JWK, err error) *MockKeyGenerator_Execute_Call {
	_c.Call.Return(privateKey, publicKey, err)
	return _c
}

func (_c *MockKeyGenerator_Execute_Call) RunAndReturn(run func() (*jwa.JWK, *jwa.JWK, error)) *MockKeyGenerator_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockKeyGenerator creates a new instance of MockKeyGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockKeyGenerator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockKeyGenerator {
	mock := &MockKeyGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
