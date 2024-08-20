// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// AIUsecase is an autogenerated mock type for the AIUsecase type
type AIUsecase struct {
	mock.Mock
}

// GenerateBlogContent provides a mock function with given fields: ctx, keywords
func (_m *AIUsecase) GenerateBlogContent(ctx context.Context, keywords string) (string, error) {
	ret := _m.Called(ctx, keywords)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, keywords)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, keywords)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAIUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAIUsecase creates a new instance of AIUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAIUsecase(t mockConstructorTestingTNewAIUsecase) *AIUsecase {
	mock := &AIUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
