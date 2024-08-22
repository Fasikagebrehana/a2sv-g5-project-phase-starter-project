// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/RealEskalate/blogpost/domain"
	mock "github.com/stretchr/testify/mock"
)

// VerifyEmail_Usecase_interface is an autogenerated mock type for the VerifyEmail_Usecase_interface type
type VerifyEmail_Usecase_interface struct {
	mock.Mock
}

// SendForgretPasswordEmail provides a mock function with given fields: id, vuser
func (_m *VerifyEmail_Usecase_interface) SendForgretPasswordEmail(id string, vuser domain.VerifyEmail) error {
	ret := _m.Called(id, vuser)

	if len(ret) == 0 {
		panic("no return value specified for SendForgretPasswordEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.VerifyEmail) error); ok {
		r0 = rf(id, vuser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendVerifyEmail provides a mock function with given fields: id, vuser
func (_m *VerifyEmail_Usecase_interface) SendVerifyEmail(id string, vuser domain.VerifyEmail) error {
	ret := _m.Called(id, vuser)

	if len(ret) == 0 {
		panic("no return value specified for SendVerifyEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.VerifyEmail) error); ok {
		r0 = rf(id, vuser)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ValidateForgetPassword provides a mock function with given fields: id, token
func (_m *VerifyEmail_Usecase_interface) ValidateForgetPassword(id string, token string) error {
	ret := _m.Called(id, token)

	if len(ret) == 0 {
		panic("no return value specified for ValidateForgetPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(id, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyUser provides a mock function with given fields: token
func (_m *VerifyEmail_Usecase_interface) VerifyUser(token string) error {
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

// NewVerifyEmail_Usecase_interface creates a new instance of VerifyEmail_Usecase_interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVerifyEmail_Usecase_interface(t interface {
	mock.TestingT
	Cleanup(func())
}) *VerifyEmail_Usecase_interface {
	mock := &VerifyEmail_Usecase_interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
