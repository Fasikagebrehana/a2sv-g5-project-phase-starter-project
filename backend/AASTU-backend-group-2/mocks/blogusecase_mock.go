// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "blog_g2/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// BlogUsecase is an autogenerated mock type for the BlogUsecase type
type BlogUsecase struct {
	mock.Mock
}

// CreateBlog provides a mock function with given fields: c, blog
func (_m *BlogUsecase) CreateBlog(c context.Context, blog *domain.Blog) error {
	ret := _m.Called(c, blog)

	if len(ret) == 0 {
		panic("no return value specified for CreateBlog")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Blog) error); ok {
		r0 = rf(c, blog)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBlog provides a mock function with given fields: c, blogID, isadmin, userid
func (_m *BlogUsecase) DeleteBlog(c context.Context, blogID string, isadmin bool, userid string) error {
	ret := _m.Called(c, blogID, isadmin, userid)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBlog")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool, string) error); ok {
		r0 = rf(c, blogID, isadmin, userid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterBlog provides a mock function with given fields: c, tag, date
func (_m *BlogUsecase) FilterBlog(c context.Context, tag []string, date time.Time) ([]domain.Blog, error) {
	ret := _m.Called(c, tag, date)

	if len(ret) == 0 {
		panic("no return value specified for FilterBlog")
	}

	var r0 []domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, time.Time) ([]domain.Blog, error)); ok {
		return rf(c, tag, date)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, time.Time) []domain.Blog); ok {
		r0 = rf(c, tag, date)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, time.Time) error); ok {
		r1 = rf(c, tag, date)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveBlog provides a mock function with given fields: c, page, sortby, dir
func (_m *BlogUsecase) RetrieveBlog(c context.Context, page int, sortby string, dir string) ([]domain.Blog, error) {
	ret := _m.Called(c, page, sortby, dir)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveBlog")
	}

	var r0 []domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string) ([]domain.Blog, error)); ok {
		return rf(c, page, sortby, dir)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string) []domain.Blog); ok {
		r0 = rf(c, page, sortby, dir)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string, string) error); ok {
		r1 = rf(c, page, sortby, dir)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchBlog provides a mock function with given fields: c, postName, authorName
func (_m *BlogUsecase) SearchBlog(c context.Context, postName string, authorName string) ([]domain.Blog, error) {
	ret := _m.Called(c, postName, authorName)

	if len(ret) == 0 {
		panic("no return value specified for SearchBlog")
	}

	var r0 []domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) ([]domain.Blog, error)); ok {
		return rf(c, postName, authorName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []domain.Blog); ok {
		r0 = rf(c, postName, authorName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(c, postName, authorName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBlog provides a mock function with given fields: c, updatedblog, blogID, isadmin, userid
func (_m *BlogUsecase) UpdateBlog(c context.Context, updatedblog domain.Blog, blogID string, isadmin bool, userid string) error {
	ret := _m.Called(c, updatedblog, blogID, isadmin, userid)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBlog")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Blog, string, bool, string) error); ok {
		r0 = rf(c, updatedblog, blogID, isadmin, userid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBlogUsecase creates a new instance of BlogUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlogUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlogUsecase {
	mock := &BlogUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
