// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "blogs/Domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// LoginUsecase is an autogenerated mock type for the LoginUsecase type
type LoginUsecase struct {
	mock.Mock
}

// CreateAccessToken provides a mock function with given fields: user, secret, expiry
func (_m *LoginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	ret := _m.Called(user, secret, expiry)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccessToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.User, string, int) (string, error)); ok {
		return rf(user, secret, expiry)
	}
	if rf, ok := ret.Get(0).(func(*domain.User, string, int) string); ok {
		r0 = rf(user, secret, expiry)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*domain.User, string, int) error); ok {
		r1 = rf(user, secret, expiry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRefreshToken provides a mock function with given fields: user, secret, expiry
func (_m *LoginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	ret := _m.Called(user, secret, expiry)

	if len(ret) == 0 {
		panic("no return value specified for CreateRefreshToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.User, string, int) (string, error)); ok {
		return rf(user, secret, expiry)
	}
	if rf, ok := ret.Get(0).(func(*domain.User, string, int) string); ok {
		r0 = rf(user, secret, expiry)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*domain.User, string, int) error); ok {
		r1 = rf(user, secret, expiry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: c, email
func (_m *LoginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ret := _m.Called(c, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (domain.User, error)); ok {
		return rf(c, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = rf(c, email)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAsActiveUser provides a mock function with given fields: user, refreshToken, c
func (_m *LoginUsecase) SaveAsActiveUser(user domain.ActiveUser, refreshToken string, c context.Context) error {
	ret := _m.Called(user, refreshToken, c)

	if len(ret) == 0 {
		panic("no return value specified for SaveAsActiveUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.ActiveUser, string, context.Context) error); ok {
		r0 = rf(user, refreshToken, c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewLoginUsecase creates a new instance of LoginUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLoginUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *LoginUsecase {
	mock := &LoginUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
