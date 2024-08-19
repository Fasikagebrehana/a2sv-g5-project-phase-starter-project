// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/RealEskalate/blogpost/domain"
	mock "github.com/stretchr/testify/mock"
)

// User_Usecase_interface is an autogenerated mock type for the User_Usecase_interface type
type User_Usecase_interface struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: id
func (_m *User_Usecase_interface) DeleteUser(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterUser provides a mock function with given fields: _a0
func (_m *User_Usecase_interface) FilterUser(_a0 map[string]string) ([]domain.ResponseUser, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FilterUser")
	}

	var r0 []domain.ResponseUser
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]string) ([]domain.ResponseUser, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(map[string]string) []domain.ResponseUser); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ResponseUser)
		}
	}

	if rf, ok := ret.Get(1).(func(map[string]string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneUser provides a mock function with given fields: id
func (_m *User_Usecase_interface) GetOneUser(id string) (domain.ResponseUser, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetOneUser")
	}

	var r0 domain.ResponseUser
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.ResponseUser, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) domain.ResponseUser); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.ResponseUser)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields:
func (_m *User_Usecase_interface) GetUsers() ([]domain.ResponseUser, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 []domain.ResponseUser
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.ResponseUser, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.ResponseUser); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ResponseUser)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePassword provides a mock function with given fields: id, updated_user
func (_m *User_Usecase_interface) UpdatePassword(id string, updated_user domain.UpdatePassword) (domain.ResponseUser, error) {
	ret := _m.Called(id, updated_user)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePassword")
	}

	var r0 domain.ResponseUser
	var r1 error
	if rf, ok := ret.Get(0).(func(string, domain.UpdatePassword) (domain.ResponseUser, error)); ok {
		return rf(id, updated_user)
	}
	if rf, ok := ret.Get(0).(func(string, domain.UpdatePassword) domain.ResponseUser); ok {
		r0 = rf(id, updated_user)
	} else {
		r0 = ret.Get(0).(domain.ResponseUser)
	}

	if rf, ok := ret.Get(1).(func(string, domain.UpdatePassword) error); ok {
		r1 = rf(id, updated_user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: id, user
func (_m *User_Usecase_interface) UpdateUser(id string, user domain.UpdateUser) (domain.ResponseUser, error) {
	ret := _m.Called(id, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 domain.ResponseUser
	var r1 error
	if rf, ok := ret.Get(0).(func(string, domain.UpdateUser) (domain.ResponseUser, error)); ok {
		return rf(id, user)
	}
	if rf, ok := ret.Get(0).(func(string, domain.UpdateUser) domain.ResponseUser); ok {
		r0 = rf(id, user)
	} else {
		r0 = ret.Get(0).(domain.ResponseUser)
	}

	if rf, ok := ret.Get(1).(func(string, domain.UpdateUser) error); ok {
		r1 = rf(id, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUser_Usecase_interface creates a new instance of User_Usecase_interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUser_Usecase_interface(t interface {
	mock.TestingT
	Cleanup(func())
}) *User_Usecase_interface {
	mock := &User_Usecase_interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
