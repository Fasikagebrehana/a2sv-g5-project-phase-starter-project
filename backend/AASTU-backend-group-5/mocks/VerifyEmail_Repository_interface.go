// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// VerifyEmail_Repository_interface is an autogenerated mock type for the VerifyEmail_Repository_interface type
type VerifyEmail_Repository_interface struct {
	mock.Mock
}

// VerifyUser provides a mock function with given fields: id
func (_m *VerifyEmail_Repository_interface) VerifyUser(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for VerifyUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewVerifyEmail_Repository_interface creates a new instance of VerifyEmail_Repository_interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewVerifyEmail_Repository_interface(t interface {
	mock.TestingT
	Cleanup(func())
}) *VerifyEmail_Repository_interface {
	mock := &VerifyEmail_Repository_interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
