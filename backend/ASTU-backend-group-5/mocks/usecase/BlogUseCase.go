// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	domain "blogApp/internal/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// BlogUseCase is an autogenerated mock type for the BlogUseCase type
type BlogUseCase struct {
	mock.Mock
}

// AddComment provides a mock function with given fields: ctx, comment, userId
func (_m *BlogUseCase) AddComment(ctx context.Context, comment *domain.Comment, userId string) error {
	ret := _m.Called(ctx, comment, userId)

	if len(ret) == 0 {
		panic("no return value specified for AddComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Comment, string) error); ok {
		r0 = rf(ctx, comment, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddLike provides a mock function with given fields: ctx, like, userId
func (_m *BlogUseCase) AddLike(ctx context.Context, like *domain.Like, userId string) error {
	ret := _m.Called(ctx, like, userId)

	if len(ret) == 0 {
		panic("no return value specified for AddLike")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Like, string) error); ok {
		r0 = rf(ctx, like, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddTagToBlog provides a mock function with given fields: ctx, blogID, tag
func (_m *BlogUseCase) AddTagToBlog(ctx context.Context, blogID string, tag domain.BlogTag) error {
	ret := _m.Called(ctx, blogID, tag)

	if len(ret) == 0 {
		panic("no return value specified for AddTagToBlog")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.BlogTag) error); ok {
		r0 = rf(ctx, blogID, tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddView provides a mock function with given fields: ctx, view, userId
func (_m *BlogUseCase) AddView(ctx context.Context, view *domain.View, userId string) error {
	ret := _m.Called(ctx, view, userId)

	if len(ret) == 0 {
		panic("no return value specified for AddView")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.View, string) error); ok {
		r0 = rf(ctx, view, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateBlog provides a mock function with given fields: ctx, _a1, authorId
func (_m *BlogUseCase) CreateBlog(ctx context.Context, _a1 *domain.Blog, authorId string) error {
	ret := _m.Called(ctx, _a1, authorId)

	if len(ret) == 0 {
		panic("no return value specified for CreateBlog")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Blog, string) error); ok {
		r0 = rf(ctx, _a1, authorId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateTag provides a mock function with given fields: ctx, tag
func (_m *BlogUseCase) CreateTag(ctx context.Context, tag *domain.BlogTag) error {
	ret := _m.Called(ctx, tag)

	if len(ret) == 0 {
		panic("no return value specified for CreateTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.BlogTag) error); ok {
		r0 = rf(ctx, tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBlog provides a mock function with given fields: ctx, id, userId
func (_m *BlogUseCase) DeleteBlog(ctx context.Context, id string, userId string) error {
	ret := _m.Called(ctx, id, userId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBlog")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, id, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteComment provides a mock function with given fields: ctx, comment, userId
func (_m *BlogUseCase) DeleteComment(ctx context.Context, comment string, userId string) error {
	ret := _m.Called(ctx, comment, userId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteComment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, comment, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTag provides a mock function with given fields: ctx, id
func (_m *BlogUseCase) DeleteTag(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterBlogs provides a mock function with given fields: ctx, filter
func (_m *BlogUseCase) FilterBlogs(ctx context.Context, filter domain.BlogFilter) ([]*domain.Blog, error) {
	ret := _m.Called(ctx, filter)

	if len(ret) == 0 {
		panic("no return value specified for FilterBlogs")
	}

	var r0 []*domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.BlogFilter) ([]*domain.Blog, error)); ok {
		return rf(ctx, filter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.BlogFilter) []*domain.Blog); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.BlogFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllBlogs provides a mock function with given fields: ctx
func (_m *BlogUseCase) GetAllBlogs(ctx context.Context) ([]*domain.Blog, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllBlogs")
	}

	var r0 []*domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*domain.Blog, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.Blog); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllTags provides a mock function with given fields: ctx
func (_m *BlogUseCase) GetAllTags(ctx context.Context) ([]*domain.BlogTag, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllTags")
	}

	var r0 []*domain.BlogTag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*domain.BlogTag, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*domain.BlogTag); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.BlogTag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlogByID provides a mock function with given fields: ctx, id
func (_m *BlogUseCase) GetBlogByID(ctx context.Context, id string) (*domain.Blog, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetBlogByID")
	}

	var r0 *domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.Blog, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Blog); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommentsByBlogID provides a mock function with given fields: ctx, blogID
func (_m *BlogUseCase) GetCommentsByBlogID(ctx context.Context, blogID string) ([]*domain.Comment, error) {
	ret := _m.Called(ctx, blogID)

	if len(ret) == 0 {
		panic("no return value specified for GetCommentsByBlogID")
	}

	var r0 []*domain.Comment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*domain.Comment, error)); ok {
		return rf(ctx, blogID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*domain.Comment); ok {
		r0 = rf(ctx, blogID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Comment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, blogID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLikesByBlogID provides a mock function with given fields: ctx, blogID
func (_m *BlogUseCase) GetLikesByBlogID(ctx context.Context, blogID string) ([]*domain.Like, error) {
	ret := _m.Called(ctx, blogID)

	if len(ret) == 0 {
		panic("no return value specified for GetLikesByBlogID")
	}

	var r0 []*domain.Like
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*domain.Like, error)); ok {
		return rf(ctx, blogID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*domain.Like); ok {
		r0 = rf(ctx, blogID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Like)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, blogID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTagByID provides a mock function with given fields: ctx, id
func (_m *BlogUseCase) GetTagByID(ctx context.Context, id string) (*domain.BlogTag, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetTagByID")
	}

	var r0 *domain.BlogTag
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*domain.BlogTag, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.BlogTag); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.BlogTag)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetViewsByBlogID provides a mock function with given fields: ctx, blogID
func (_m *BlogUseCase) GetViewsByBlogID(ctx context.Context, blogID string) ([]*domain.View, error) {
	ret := _m.Called(ctx, blogID)

	if len(ret) == 0 {
		panic("no return value specified for GetViewsByBlogID")
	}

	var r0 []*domain.View
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*domain.View, error)); ok {
		return rf(ctx, blogID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*domain.View); ok {
		r0 = rf(ctx, blogID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.View)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, blogID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PaginateBlogs provides a mock function with given fields: ctx, filter, page, pageSize
func (_m *BlogUseCase) PaginateBlogs(ctx context.Context, filter domain.BlogFilter, page int, pageSize int) ([]*domain.Blog, error) {
	ret := _m.Called(ctx, filter, page, pageSize)

	if len(ret) == 0 {
		panic("no return value specified for PaginateBlogs")
	}

	var r0 []*domain.Blog
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.BlogFilter, int, int) ([]*domain.Blog, error)); ok {
		return rf(ctx, filter, page, pageSize)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.BlogFilter, int, int) []*domain.Blog); ok {
		r0 = rf(ctx, filter, page, pageSize)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Blog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.BlogFilter, int, int) error); ok {
		r1 = rf(ctx, filter, page, pageSize)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveLike provides a mock function with given fields: ctx, likeId, userId
func (_m *BlogUseCase) RemoveLike(ctx context.Context, likeId string, userId string) error {
	ret := _m.Called(ctx, likeId, userId)

	if len(ret) == 0 {
		panic("no return value specified for RemoveLike")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, likeId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveTagFromBlog provides a mock function with given fields: ctx, blogID, tagID
func (_m *BlogUseCase) RemoveTagFromBlog(ctx context.Context, blogID string, tagID string) error {
	ret := _m.Called(ctx, blogID, tagID)

	if len(ret) == 0 {
		panic("no return value specified for RemoveTagFromBlog")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, blogID, tagID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateBlog provides a mock function with given fields: ctx, id, _a2, authorId
func (_m *BlogUseCase) UpdateBlog(ctx context.Context, id string, _a2 *domain.Blog, authorId string) error {
	ret := _m.Called(ctx, id, _a2, authorId)

	if len(ret) == 0 {
		panic("no return value specified for UpdateBlog")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.Blog, string) error); ok {
		r0 = rf(ctx, id, _a2, authorId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTag provides a mock function with given fields: ctx, id, tag
func (_m *BlogUseCase) UpdateTag(ctx context.Context, id string, tag *domain.BlogTag) error {
	ret := _m.Called(ctx, id, tag)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTag")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *domain.BlogTag) error); ok {
		r0 = rf(ctx, id, tag)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewBlogUseCase creates a new instance of BlogUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlogUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlogUseCase {
	mock := &BlogUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
