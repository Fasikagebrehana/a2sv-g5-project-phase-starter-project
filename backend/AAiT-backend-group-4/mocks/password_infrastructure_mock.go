// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PasswordInfrastructure is an autogenerated mock type for the PasswordInfrastructure type
type PasswordInfrastructure struct {
	mock.Mock
}

// ComparePasswords provides a mock function with given fields: password, hashedPassword
func (_m *PasswordInfrastructure) ComparePasswords(password string, hashedPassword string) error {
	ret := _m.Called(password, hashedPassword)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(password, hashedPassword)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HashPassword provides a mock function with given fields: password
func (_m *PasswordInfrastructure) HashPassword(password string) (string, error) {
	ret := _m.Called(password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPasswordInfrastructure interface {
	mock.TestingT
	Cleanup(func())
}

// NewPasswordInfrastructure creates a new instance of PasswordInfrastructure. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPasswordInfrastructure(t mockConstructorTestingTNewPasswordInfrastructure) *PasswordInfrastructure {
	mock := &PasswordInfrastructure{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
