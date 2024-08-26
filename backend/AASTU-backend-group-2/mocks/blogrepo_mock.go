// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "blog_g2/domain"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// BlogRepository is an autogenerated mock type for the BlogRepository type
type BlogRepository struct {
	mock.Mock
}

// CreateBlog provides a mock function with given fields: blog
func (_m *BlogRepository) CreateBlog(blog *domain.Blog) error {
	ret := _m.Called(blog)

	if len(ret) == 0 {
		panic("no return value specified for CreateBlog")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Blog) error); ok {
		r0 = rf(blog)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBlog provides a mock function with given fields: blogID, isadmin, userid
func (_m *BlogRepository) DeleteBlog(blogID string, isadmin bool, userid string) error {
	ret := _m.Called(blogID, isadmin, userid)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBlog")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, bool, string) error); ok {
		r0 = rf(blogID, isadmin, userid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterBlog provides a mock function with given fields: tag, date
func (_m *BlogRepository) FilterBlog(tag []string, date time.Time) ([]domain.Blog, error) {
	ret := _m.Called(tag, date)

	if len(ret) == 0 {
		panic("no return value specified for FilterBlog")
	}

	var r0 []domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func([]string, time.Time) ([]domain.Blog, error)); ok {
		return rf(tag, date)
	}
	if rf, ok := ret.Get(0).(func([]string, time.Time) []domain.Blog); ok {
		r0 = rf(tag, date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func([]string, time.Time) error); ok {
		r1 = rf(tag, date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveBlog provides a mock function with given fields: pgnum, sortby, dir
func (_m *BlogRepository) RetrieveBlog(pgnum int, sortby string, dir string) ([]domain.Blog, int, error) {
	ret := _m.Called(pgnum, sortby, dir)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveBlog")
	}

	var r0 []domain.Blog
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func(int, string, string) ([]domain.Blog, int, error)); ok {
		return rf(pgnum, sortby, dir)
	}
	if rf, ok := ret.Get(0).(func(int, string, string) []domain.Blog); ok {
		r0 = rf(pgnum, sortby, dir)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(int, string, string) int); ok {
		r1 = rf(pgnum, sortby, dir)
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func(int, string, string) error); ok {
		r2 = rf(pgnum, sortby, dir)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SearchBlog provides a mock function with given fields: postName, authorName
func (_m *BlogRepository) SearchBlog(postName string, authorName string) ([]domain.Blog, error) {
	ret := _m.Called(postName, authorName)

	if len(ret) == 0 {
		panic("no return value specified for SearchBlog")
	}

	var r0 []domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]domain.Blog, error)); ok {
		return rf(postName, authorName)
	}
	if rf, ok := ret.Get(0).(func(string, string) []domain.Blog); ok {
		r0 = rf(postName, authorName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(postName, authorName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBlog provides a mock function with given fields: updatedblog, blogID, isadmin, userid
func (_m *BlogRepository) UpdateBlog(updatedblog domain.Blog, blogID string, isadmin bool, userid string) error {
	ret := _m.Called(updatedblog, blogID, isadmin, userid)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBlog")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.Blog, string, bool, string) error); ok {
		r0 = rf(updatedblog, blogID, isadmin, userid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBlogRepository creates a new instance of BlogRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlogRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlogRepository {
	mock := &BlogRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
