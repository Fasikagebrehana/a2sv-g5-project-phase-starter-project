// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "blog/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: c, user, claims
func (_m *UserUsecase) CreateUser(c context.Context, user *domain.CreateUser, claims *domain.JwtCustomClaims) error {
	ret := _m.Called(c, user, claims)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.CreateUser, *domain.JwtCustomClaims) error); ok {
		r0 = rf(c, user, claims)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: c, id, claims
func (_m *UserUsecase) DeleteUser(c context.Context, id primitive.ObjectID, claims *domain.JwtCustomClaims) error {
	ret := _m.Called(c, id, claims)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, *domain.JwtCustomClaims) error); ok {
		r0 = rf(c, id, claims)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DemoteUser provides a mock function with given fields: c, id, claims
func (_m *UserUsecase) DemoteUser(c context.Context, id primitive.ObjectID, claims *domain.JwtCustomClaims) error {
	ret := _m.Called(c, id, claims)

	if len(ret) == 0 {
		panic("no return value specified for DemoteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, *domain.JwtCustomClaims) error); ok {
		r0 = rf(c, id, claims)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllUsers provides a mock function with given fields: c
func (_m *UserUsecase) GetAllUsers(c context.Context) ([]*domain.User, error) {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for GetAllUsers")
	}

	var r0 []*domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*domain.User, error)); ok {
		return rf(c)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.User); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: c, email
func (_m *UserUsecase) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	ret := _m.Called(c, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.User, error)); ok {
		return rf(c, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(c, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: c, id
func (_m *UserUsecase) GetUserByID(c context.Context, id primitive.ObjectID) (*domain.User, error) {
	ret := _m.Called(c, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) (*domain.User, error)); ok {
		return rf(c, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *domain.User); ok {
		r0 = rf(c, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsername provides a mock function with given fields: c, username
func (_m *UserUsecase) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	ret := _m.Called(c, username)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByUsername")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.User, error)); ok {
		return rf(c, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.User); ok {
		r0 = rf(c, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PromoteUser provides a mock function with given fields: c, id, claims
func (_m *UserUsecase) PromoteUser(c context.Context, id primitive.ObjectID, claims *domain.JwtCustomClaims) error {
	ret := _m.Called(c, id, claims)

	if len(ret) == 0 {
		panic("no return value specified for PromoteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, *domain.JwtCustomClaims) error); ok {
		r0 = rf(c, id, claims)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: c, user, claims, existinguser
func (_m *UserUsecase) UpdateUser(c context.Context, user *domain.User, claims *domain.JwtCustomClaims, existinguser *domain.User) (*domain.User, error) {
	ret := _m.Called(c, user, claims, existinguser)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User, *domain.JwtCustomClaims, *domain.User) (*domain.User, error)); ok {
		return rf(c, user, claims, existinguser)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User, *domain.JwtCustomClaims, *domain.User) *domain.User); ok {
		r0 = rf(c, user, claims, existinguser)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.User, *domain.JwtCustomClaims, *domain.User) error); ok {
		r1 = rf(c, user, claims, existinguser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
