// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	domain "Blog_Starter/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// BlogRatingRepository is an autogenerated mock type for the BlogRatingRepository type
type BlogRatingRepository struct {
	mock.Mock
}

// DeleteRating provides a mock function with given fields: c, ratingID
func (_m *BlogRatingRepository) DeleteRating(c context.Context, ratingID string) (*domain.BlogRating, error) {
	ret := _m.Called(c, ratingID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteRating")
	}

	var r0 *domain.BlogRating
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.BlogRating, error)); ok {
		return rf(c, ratingID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.BlogRating); ok {
		r0 = rf(c, ratingID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BlogRating)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, ratingID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRatingByBlogID provides a mock function with given fields: c, blogID
func (_m *BlogRatingRepository) GetRatingByBlogID(c context.Context, blogID string) ([]*domain.BlogRating, error) {
	ret := _m.Called(c, blogID)

	if len(ret) == 0 {
		panic("no return value specified for GetRatingByBlogID")
	}

	var r0 []*domain.BlogRating
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*domain.BlogRating, error)); ok {
		return rf(c, blogID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*domain.BlogRating); ok {
		r0 = rf(c, blogID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.BlogRating)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, blogID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRatingByID provides a mock function with given fields: c, ratingID
func (_m *BlogRatingRepository) GetRatingByID(c context.Context, ratingID string) (*domain.BlogRating, error) {
	ret := _m.Called(c, ratingID)

	if len(ret) == 0 {
		panic("no return value specified for GetRatingByID")
	}

	var r0 *domain.BlogRating
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.BlogRating, error)); ok {
		return rf(c, ratingID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.BlogRating); ok {
		r0 = rf(c, ratingID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BlogRating)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, ratingID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertRating provides a mock function with given fields: c, rating
func (_m *BlogRatingRepository) InsertRating(c context.Context, rating *domain.BlogRating) (*domain.BlogRating, error) {
	ret := _m.Called(c, rating)

	if len(ret) == 0 {
		panic("no return value specified for InsertRating")
	}

	var r0 *domain.BlogRating
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BlogRating) (*domain.BlogRating, error)); ok {
		return rf(c, rating)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BlogRating) *domain.BlogRating); ok {
		r0 = rf(c, rating)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BlogRating)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.BlogRating) error); ok {
		r1 = rf(c, rating)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRating provides a mock function with given fields: c, rating, ratingID
func (_m *BlogRatingRepository) UpdateRating(c context.Context, rating int, ratingID string) (*domain.BlogRating, int, error) {
	ret := _m.Called(c, rating, ratingID)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRating")
	}

	var r0 *domain.BlogRating
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string) (*domain.BlogRating, int, error)); ok {
		return rf(c, rating, ratingID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string) *domain.BlogRating); ok {
		r0 = rf(c, rating, ratingID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BlogRating)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string) int); ok {
		r1 = rf(c, rating, ratingID)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(context.Context, int, string) error); ok {
		r2 = rf(c, rating, ratingID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// NewBlogRatingRepository creates a new instance of BlogRatingRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlogRatingRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlogRatingRepository {
	mock := &BlogRatingRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
