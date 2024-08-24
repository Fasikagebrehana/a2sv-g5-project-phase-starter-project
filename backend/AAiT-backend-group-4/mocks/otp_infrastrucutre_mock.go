// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "aait-backend-group4/Domain"

	mock "github.com/stretchr/testify/mock"
)

// OtpInfrastructure is an autogenerated mock type for the OtpInfrastructure type
type OtpInfrastructure struct {
	mock.Mock
}

// CreateOTP provides a mock function with given fields: otp
func (_m *OtpInfrastructure) CreateOTP(otp *domain.UserOTPRequest) (string, error) {
	ret := _m.Called(otp)

	var r0 string
	if rf, ok := ret.Get(0).(func(*domain.UserOTPRequest) string); ok {
		r0 = rf(otp)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.UserOTPRequest) error); ok {
		r1 = rf(otp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendEmail provides a mock function with given fields: email, subject, key, otp
func (_m *OtpInfrastructure) SendEmail(email string, subject string, key string, otp string) error {
	ret := _m.Called(email, subject, key, otp)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(email, subject, key, otp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendPasswordResetEmail provides a mock function with given fields: email, subject, key
func (_m *OtpInfrastructure) SendPasswordResetEmail(email string, subject string, key string) error {
	ret := _m.Called(email, subject, key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(email, subject, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewOtpInfrastructure interface {
	mock.TestingT
	Cleanup(func())
}

// NewOtpInfrastructure creates a new instance of OtpInfrastructure. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOtpInfrastructure(t mockConstructorTestingTNewOtpInfrastructure) *OtpInfrastructure {
	mock := &OtpInfrastructure{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
