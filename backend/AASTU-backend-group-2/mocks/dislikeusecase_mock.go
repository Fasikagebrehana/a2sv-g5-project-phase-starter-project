// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "blog_g2/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// DisLikeUsecase is an autogenerated mock type for the DisLikeUsecase type
type DisLikeUsecase struct {
	mock.Mock
}

// CreateDisLike provides a mock function with given fields: _a0, user_id, post_id
func (_m *DisLikeUsecase) CreateDisLike(_a0 context.Context, user_id string, post_id string) error {
	ret := _m.Called(_a0, user_id, post_id)

	if len(ret) == 0 {
		panic("no return value specified for CreateDisLike")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(_a0, user_id, post_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDisLike provides a mock function with given fields: _a0, like_id
func (_m *DisLikeUsecase) DeleteDisLike(_a0 context.Context, like_id string) error {
	ret := _m.Called(_a0, like_id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDisLike")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, like_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDisLikes provides a mock function with given fields: _a0, post_id
func (_m *DisLikeUsecase) GetDisLikes(_a0 context.Context, post_id string) ([]domain.DisLike, error) {
	ret := _m.Called(_a0, post_id)

	if len(ret) == 0 {
		panic("no return value specified for GetDisLikes")
	}

	var r0 []domain.DisLike
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]domain.DisLike, error)); ok {
		return rf(_a0, post_id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []domain.DisLike); ok {
		r0 = rf(_a0, post_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.DisLike)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, post_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDisLikeUsecase creates a new instance of DisLikeUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDisLikeUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *DisLikeUsecase {
	mock := &DisLikeUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
