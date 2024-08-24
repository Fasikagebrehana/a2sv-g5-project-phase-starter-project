// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "blogs/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CheckRoot provides a mock function with given fields:
func (_m *UserRepository) CheckRoot() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CheckRoot")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckUsernameAndEmail provides a mock function with given fields: username, email
func (_m *UserRepository) CheckUsernameAndEmail(username string, email string) error {
	ret := _m.Called(username, email)

	if len(ret) == 0 {
		panic("no return value specified for CheckUsernameAndEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(username, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteToken provides a mock function with given fields: username
func (_m *UserRepository) DeleteToken(username string) error {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for DeleteToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: username
func (_m *UserRepository) DeleteUser(username string) error {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTokenByUsername provides a mock function with given fields: username
func (_m *UserRepository) GetTokenByUsername(username string) (*domain.Token, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for GetTokenByUsername")
	}

	var r0 *domain.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Token, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Token); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsernameorEmail provides a mock function with given fields: usernameoremail
func (_m *UserRepository) GetUserByUsernameorEmail(usernameoremail string) (*domain.User, error) {
	ret := _m.Called(usernameoremail)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByUsernameorEmail")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(usernameoremail)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(usernameoremail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(usernameoremail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertToken provides a mock function with given fields: token
func (_m *UserRepository) InsertToken(token *domain.Token) error {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for InsertToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Token) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterUser provides a mock function with given fields: user
func (_m *UserRepository) RegisterUser(user *domain.User) error {
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

// Resetpassword provides a mock function with given fields: usernameoremail, password
func (_m *UserRepository) Resetpassword(usernameoremail string, password string) error {
	ret := _m.Called(usernameoremail, password)

	if len(ret) == 0 {
		panic("no return value specified for Resetpassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(usernameoremail, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateProfile provides a mock function with given fields: usernameoremail, user
func (_m *UserRepository) UpdateProfile(usernameoremail string, user *domain.User) error {
	ret := _m.Called(usernameoremail, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateProfile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *domain.User) error); ok {
		r0 = rf(usernameoremail, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
