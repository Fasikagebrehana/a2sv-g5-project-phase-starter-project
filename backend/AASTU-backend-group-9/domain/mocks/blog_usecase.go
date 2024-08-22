// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	domain "blog/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"

	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// BlogUsecase is an autogenerated mock type for the BlogUsecase type
type BlogUsecase struct {
	mock.Mock
}

// AddComment provides a mock function with given fields: ctx, id, userID, comment
func (_m *BlogUsecase) AddComment(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID, comment *domain.Comment) error {
	ret := _m.Called(ctx, id, userID, comment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, primitive.ObjectID, *domain.Comment) error); ok {
		r0 = rf(ctx, id, userID, comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateBlog provides a mock function with given fields: ctx, blog, claims
func (_m *BlogUsecase) CreateBlog(ctx context.Context, blog *domain.BlogCreationRequest, claims *domain.JwtCustomClaims) (*domain.BlogResponse, error) {
	ret := _m.Called(ctx, blog, claims)

	var r0 *domain.BlogResponse
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BlogCreationRequest, *domain.JwtCustomClaims) *domain.BlogResponse); ok {
		r0 = rf(ctx, blog, claims)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BlogResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.BlogCreationRequest, *domain.JwtCustomClaims) error); ok {
		r1 = rf(ctx, blog, claims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteBlog provides a mock function with given fields: ctx, id
func (_m *BlogUsecase) DeleteBlog(ctx context.Context, id primitive.ObjectID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteComment provides a mock function with given fields: ctx, post_id, comment_id, userID
func (_m *BlogUsecase) DeleteComment(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID) error {
	ret := _m.Called(ctx, post_id, comment_id, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, primitive.ObjectID, primitive.ObjectID) error); ok {
		r0 = rf(ctx, post_id, comment_id, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterBlogs provides a mock function with given fields: ctx, popularity, tags, startDate, endDate
func (_m *BlogUsecase) FilterBlogs(ctx context.Context, popularity string, tags []string, startDate string, endDate string) ([]*domain.Blog, error) {
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
func (_m *BlogUsecase) GetAllBlogs(ctx context.Context, page int, limit int, sortBy string) ([]*domain.BlogResponse, error) {
	ret := _m.Called(ctx, page, limit, sortBy)

	var r0 []*domain.BlogResponse
	if rf, ok := ret.Get(0).(func(context.Context, int, int, string) []*domain.BlogResponse); ok {
		r0 = rf(ctx, page, limit, sortBy)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.BlogResponse)
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
func (_m *BlogUsecase) GetBlogByID(ctx context.Context, id primitive.ObjectID) (*domain.BlogResponse, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.BlogResponse
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *domain.BlogResponse); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BlogResponse)
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

// GetComments provides a mock function with given fields: ctx, post_id
func (_m *BlogUsecase) GetComments(ctx context.Context, post_id primitive.ObjectID) (*domain.Comment, error) {
	ret := _m.Called(ctx, post_id)

	var r0 *domain.Comment
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) *domain.Comment); ok {
		r0 = rf(ctx, post_id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Comment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID) error); ok {
		r1 = rf(ctx, post_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchBlogs provides a mock function with given fields: ctx, title, author
func (_m *BlogUsecase) SearchBlogs(ctx context.Context, title string, author string) (*[]domain.Blog, error) {
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

// TrackDislike provides a mock function with given fields: ctx, id, userID
func (_m *BlogUsecase) TrackDislike(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) error {
	ret := _m.Called(ctx, id, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, primitive.ObjectID) error); ok {
		r0 = rf(ctx, id, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TrackLike provides a mock function with given fields: ctx, id, userID
func (_m *BlogUsecase) TrackLike(ctx context.Context, id primitive.ObjectID, userID primitive.ObjectID) error {
	ret := _m.Called(ctx, id, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, primitive.ObjectID) error); ok {
		r0 = rf(ctx, id, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TrackView provides a mock function with given fields: ctx, id
func (_m *BlogUsecase) TrackView(ctx context.Context, id primitive.ObjectID) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBlog provides a mock function with given fields: ctx, id, blog
func (_m *BlogUsecase) UpdateBlog(ctx context.Context, id primitive.ObjectID, blog *domain.BlogUpdateRequest) (*domain.BlogResponse, error) {
	ret := _m.Called(ctx, id, blog)

	var r0 *domain.BlogResponse
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, *domain.BlogUpdateRequest) *domain.BlogResponse); ok {
		r0 = rf(ctx, id, blog)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BlogResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, primitive.ObjectID, *domain.BlogUpdateRequest) error); ok {
		r1 = rf(ctx, id, blog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateComment provides a mock function with given fields: ctx, post_id, comment_id, userID, comment
func (_m *BlogUsecase) UpdateComment(ctx context.Context, post_id primitive.ObjectID, comment_id primitive.ObjectID, userID primitive.ObjectID, comment *domain.Comment) error {
	ret := _m.Called(ctx, post_id, comment_id, userID, comment)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, primitive.ObjectID, primitive.ObjectID, primitive.ObjectID, *domain.Comment) error); ok {
		r0 = rf(ctx, post_id, comment_id, userID, comment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBlogUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewBlogUsecase creates a new instance of BlogUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBlogUsecase(t mockConstructorTestingTNewBlogUsecase) *BlogUsecase {
	mock := &BlogUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
