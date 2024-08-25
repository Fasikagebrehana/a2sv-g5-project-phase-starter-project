// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	domain "Blog_Starter/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// LoginRepository is an autogenerated mock type for the LoginRepository type
type LoginRepository struct {
	mock.Mock
}

// Login provides a mock function with given fields: c, user
func (_m *LoginRepository) Login(c context.Context, user *domain.UserLogin) (*domain.LoginResponse, error) {
	ret := _m.Called(c, user)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *domain.LoginResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.UserLogin) (*domain.LoginResponse, error)); ok {
		return rf(c, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.UserLogin) *domain.LoginResponse); ok {
		r0 = rf(c, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.LoginResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.UserLogin) error); ok {
		r1 = rf(c, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePassword provides a mock function with given fields: c, req, userID
func (_m *LoginRepository) UpdatePassword(c context.Context, req domain.ChangePasswordRequest, userID string) error {
	ret := _m.Called(c, req, userID)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.ChangePasswordRequest, string) error); ok {
		r0 = rf(c, req, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewLoginRepository creates a new instance of LoginRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLoginRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *LoginRepository {
	mock := &LoginRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
