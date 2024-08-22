// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "blog/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// BlogRepository is an autogenerated mock type for the BlogRepository type
type BlogRepository struct {
	mock.Mock
}

// AddComment provides a mock function with given fields: ctx, id, comment
func (_m *BlogRepository) AddComment(ctx context.Context, id primitive.ObjectID, comment *domain.Comment) error {
	ret := _m.Called(ctx, id, comment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, *domain.Comment) error); ok {
		r0 = rf(ctx, id, comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateBlog provides a mock function with given fields: ctx, blog
func (_m *BlogRepository) CreateBlog(ctx context.Context, blog *domain.Blog) error {
	ret := _m.Called(ctx, blog)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Blog) error); ok {
		r0 = rf(ctx, blog)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DecrementPopularity provides a mock function with given fields: ctx, id, metric
func (_m *BlogRepository) DecrementPopularity(ctx context.Context, id primitive.ObjectID, metric string) error {
	ret := _m.Called(ctx, id, metric)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, string) error); ok {
		r0 = rf(ctx, id, metric)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBlog provides a mock function with given fields: ctx, id
func (_m *BlogRepository) DeleteBlog(ctx context.Context, id primitive.ObjectID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterBlogs provides a mock function with given fields: ctx, popularity, tags, startDate, endDate
func (_m *BlogRepository) FilterBlogs(ctx context.Context, popularity string, tags []string, startDate string, endDate string) ([]*domain.Blog, error) {
	ret := _m.Called(ctx, popularity, tags, startDate, endDate)

	var r0 []*domain.Blog
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, string, string) []*domain.Blog); ok {
		r0 = rf(ctx, popularity, tags, startDate, endDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Blog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string, string, string) error); ok {
		r1 = rf(ctx, popularity, tags, startDate, endDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllBlogs provides a mock function with given fields: ctx, page, limit, sortBy
func (_m *BlogRepository) GetAllBlogs(ctx context.Context, page int, limit int, sortBy string) ([]*domain.Blog, error) {
	ret := _m.Called(ctx, page, limit, sortBy)

	var r0 []*domain.Blog
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string) []*domain.Blog); ok {
		r0 = rf(ctx, page, limit, sortBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Blog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int, string) error); ok {
		r1 = rf(ctx, page, limit, sortBy)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlogByID provides a mock function with given fields: ctx, id
func (_m *BlogRepository) GetBlogByID(ctx context.Context, id primitive.ObjectID) (*domain.Blog, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Blog
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *domain.Blog); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Blog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HasUserDisliked provides a mock function with given fields: ctx, id, userID
func (_m *BlogRepository) HasUserDisliked(ctx context.Context, id primitive.ObjectID, userID string) (bool, error) {
	ret := _m.Called(ctx, id, userID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, string) bool); ok {
		r0 = rf(ctx, id, userID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID, string) error); ok {
		r1 = rf(ctx, id, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrementPopularity provides a mock function with given fields: ctx, id, metric
func (_m *BlogRepository) IncrementPopularity(ctx context.Context, id primitive.ObjectID, metric string) error {
	ret := _m.Called(ctx, id, metric)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, string) error); ok {
		r0 = rf(ctx, id, metric)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchBlogs provides a mock function with given fields: ctx, title, author
func (_m *BlogRepository) SearchBlogs(ctx context.Context, title string, author string) (*[]domain.Blog, error) {
	ret := _m.Called(ctx, title, author)

	var r0 *[]domain.Blog
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *[]domain.Blog); ok {
		r0 = rf(ctx, title, author)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]domain.Blog)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, title, author)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBlog provides a mock function with given fields: ctx, blog
func (_m *BlogRepository) UpdateBlog(ctx context.Context, blog *domain.Blog) error {
	ret := _m.Called(ctx, blog)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Blog) error); ok {
		r0 = rf(ctx, blog)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBlogRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewBlogRepository creates a new instance of BlogRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBlogRepository(t mockConstructorTestingTNewBlogRepository) *BlogRepository {
	mock := &BlogRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
