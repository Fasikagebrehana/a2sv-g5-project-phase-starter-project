// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "blogs/Domain"

	mock "github.com/stretchr/testify/mock"
)

// Infrastructure is an autogenerated mock type for the Infrastructure type
type Infrastructure struct {
	mock.Mock
}

// CreateAccessToken provides a mock function with given fields: user, secret, expiry
func (_m *Infrastructure) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
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
func (_m *Infrastructure) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
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

// ExtractIDFromToken provides a mock function with given fields: requestToken, secret
func (_m *Infrastructure) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	ret := _m.Called(requestToken, secret)

	if len(ret) == 0 {
		panic("no return value specified for ExtractIDFromToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (string, error)); ok {
		return rf(requestToken, secret)
	}
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(requestToken, secret)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(requestToken, secret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewInfrastructure creates a new instance of Infrastructure. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInfrastructure(t interface {
	mock.TestingT
	Cleanup(func())
}) *Infrastructure {
	mock := &Infrastructure{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
