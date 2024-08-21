// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "blogs/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// AddRoot provides a mock function with given fields:
func (_m *UserUsecase) AddRoot() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AddRoot")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForgotPassword provides a mock function with given fields: email, newPassword
func (_m *UserUsecase) ForgotPassword(email string, newPassword string) error {
	ret := _m.Called(email, newPassword)

	if len(ret) == 0 {
		panic("no return value specified for ForgotPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(email, newPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GoogleCallback provides a mock function with given fields: state, code
func (_m *UserUsecase) GoogleCallback(state string, code string) (string, string, error) {
	ret := _m.Called(state, code)

	if len(ret) == 0 {
		panic("no return value specified for GoogleCallback")
	}

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (string, string, error)); ok {
		return rf(state, code)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(state, code)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(state, code)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(state, code)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GoogleLogin provides a mock function with given fields:
func (_m *UserUsecase) GoogleLogin() (string, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GoogleLogin")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginUser provides a mock function with given fields: usernameoremail, password
func (_m *UserUsecase) LoginUser(usernameoremail string, password string) (string, string, error) {
	ret := _m.Called(usernameoremail, password)

	if len(ret) == 0 {
		panic("no return value specified for LoginUser")
	}

	var r0 string
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (string, string, error)); ok {
		return rf(usernameoremail, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(usernameoremail, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(usernameoremail, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(usernameoremail, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// LogoutUser provides a mock function with given fields: username
func (_m *UserUsecase) LogoutUser(username string) error {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for LogoutUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PromoteUser provides a mock function with given fields: username, promoted, claims
func (_m *UserUsecase) PromoteUser(username string, promoted bool, claims *domain.LoginClaims) error {
	ret := _m.Called(username, promoted, claims)

	if len(ret) == 0 {
		panic("no return value specified for PromoteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool, *domain.LoginClaims) error); ok {
		r0 = rf(username, promoted, claims)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RefreshToken provides a mock function with given fields: claims
func (_m *UserUsecase) RefreshToken(claims *domain.LoginClaims) (string, error) {
	ret := _m.Called(claims)

	if len(ret) == 0 {
		panic("no return value specified for RefreshToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.LoginClaims) (string, error)); ok {
		return rf(claims)
	}
	if rf, ok := ret.Get(0).(func(*domain.LoginClaims) string); ok {
		r0 = rf(claims)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*domain.LoginClaims) error); ok {
		r1 = rf(claims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: user
func (_m *UserUsecase) RegisterUser(user *domain.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetPassword provides a mock function with given fields: tokenString
func (_m *UserUsecase) ResetPassword(tokenString string) error {
	ret := _m.Called(tokenString)

	if len(ret) == 0 {
		panic("no return value specified for ResetPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(tokenString)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProfile provides a mock function with given fields: user, claims
func (_m *UserUsecase) UpdateProfile(user *domain.User, claims *domain.LoginClaims) error {
	ret := _m.Called(user, claims)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProfile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User, *domain.LoginClaims) error); ok {
		r0 = rf(user, claims)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyUser provides a mock function with given fields: token
func (_m *UserUsecase) VerifyUser(token string) error {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for VerifyUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
