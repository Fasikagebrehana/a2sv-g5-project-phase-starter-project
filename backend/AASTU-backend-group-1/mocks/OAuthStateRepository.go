// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "blogs/domain"

	mock "github.com/stretchr/testify/mock"
)

// OAuthStateRepository is an autogenerated mock type for the OAuthStateRepository type
type OAuthStateRepository struct {
	mock.Mock
}

// DeleteState provides a mock function with given fields: state
func (_m *OAuthStateRepository) DeleteState(state *domain.OAuthState) error {
	ret := _m.Called(state)

	if len(ret) == 0 {
		panic("no return value specified for DeleteState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.OAuthState) error); ok {
		r0 = rf(state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetState provides a mock function with given fields: stateString
func (_m *OAuthStateRepository) GetState(stateString string) (*domain.OAuthState, error) {
	ret := _m.Called(stateString)

	if len(ret) == 0 {
		panic("no return value specified for GetState")
	}

	var r0 *domain.OAuthState
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.OAuthState, error)); ok {
		return rf(stateString)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.OAuthState); ok {
		r0 = rf(stateString)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.OAuthState)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(stateString)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertState provides a mock function with given fields: state
func (_m *OAuthStateRepository) InsertState(state *domain.OAuthState) error {
	ret := _m.Called(state)

	if len(ret) == 0 {
		panic("no return value specified for InsertState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.OAuthState) error); ok {
		r0 = rf(state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOAuthStateRepository creates a new instance of OAuthStateRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOAuthStateRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *OAuthStateRepository {
	mock := &OAuthStateRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
