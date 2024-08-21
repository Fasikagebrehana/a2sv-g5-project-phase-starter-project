// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "github.com/RealEskalate/blogpost/domain"
	mock "github.com/stretchr/testify/mock"
)

// Blog_Repository_interface is an autogenerated mock type for the Blog_Repository_interface type
type Blog_Repository_interface struct {
	mock.Mock
}

// CreateBlogDocunent provides a mock function with given fields: blog
func (_m *Blog_Repository_interface) CreateBlogDocunent(blog domain.Blog) (domain.Blog, error) {
	ret := _m.Called(blog)

	if len(ret) == 0 {
		panic("no return value specified for CreateBlogDocunent")
	}

	var r0 domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.Blog) (domain.Blog, error)); ok {
		return rf(blog)
	}
	if rf, ok := ret.Get(0).(func(domain.Blog) domain.Blog); ok {
		r0 = rf(blog)
	} else {
		r0 = ret.Get(0).(domain.Blog)
	}

	if rf, ok := ret.Get(1).(func(domain.Blog) error); ok {
		r1 = rf(blog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBlogDocument provides a mock function with given fields: id, userID
func (_m *Blog_Repository_interface) DeleteBlogDocument(id string, userID string) error {
	ret := _m.Called(id, userID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBlogDocument")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(id, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterBlogDocunent provides a mock function with given fields: _a0
func (_m *Blog_Repository_interface) FilterBlogDocunent(_a0 map[string]string) ([]domain.Blog, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for FilterBlogDocunent")
	}

	var r0 []domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]string) ([]domain.Blog, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(map[string]string) []domain.Blog); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(map[string]string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlogDocunents provides a mock function with given fields: offset, limit
func (_m *Blog_Repository_interface) GetBlogDocunents(offset int, limit int) ([]domain.Blog, error) {
	ret := _m.Called(offset, limit)

	if len(ret) == 0 {
		panic("no return value specified for GetBlogDocunents")
	}

	var r0 []domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int) ([]domain.Blog, error)); ok {
		return rf(offset, limit)
	}
	if rf, ok := ret.Get(0).(func(int, int) []domain.Blog); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(offset, limit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOneBlogDocunent provides a mock function with given fields: id
func (_m *Blog_Repository_interface) GetOneBlogDocunent(id string) (domain.Blog, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetOneBlogDocunent")
	}

	var r0 domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.Blog, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) domain.Blog); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.Blog)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBlogDocunent provides a mock function with given fields: id, blog
func (_m *Blog_Repository_interface) UpdateBlogDocunent(id string, blog domain.Blog) (domain.Blog, error) {
	ret := _m.Called(id, blog)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBlogDocunent")
	}

	var r0 domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(string, domain.Blog) (domain.Blog, error)); ok {
		return rf(id, blog)
	}
	if rf, ok := ret.Get(0).(func(string, domain.Blog) domain.Blog); ok {
		r0 = rf(id, blog)
	} else {
		r0 = ret.Get(0).(domain.Blog)
	}

	if rf, ok := ret.Get(1).(func(string, domain.Blog) error); ok {
		r1 = rf(id, blog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBlog_Repository_interface creates a new instance of Blog_Repository_interface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlog_Repository_interface(t interface {
	mock.TestingT
	Cleanup(func())
}) *Blog_Repository_interface {
	mock := &Blog_Repository_interface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
