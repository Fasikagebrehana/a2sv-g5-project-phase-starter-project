// Code generated by MockGen. DO NOT EDIT.
// Source: Domain/Interfaces/oauth_interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockOAuthController is a mock of OAuthController interface.
type MockOAuthController struct {
	ctrl     *gomock.Controller
	recorder *MockOAuthControllerMockRecorder
}

// MockOAuthControllerMockRecorder is the mock recorder for MockOAuthController.
type MockOAuthControllerMockRecorder struct {
	mock *MockOAuthController
}

// NewMockOAuthController creates a new mock instance.
func NewMockOAuthController(ctrl *gomock.Controller) *MockOAuthController {
	mock := &MockOAuthController{ctrl: ctrl}
	mock.recorder = &MockOAuthControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOAuthController) EXPECT() *MockOAuthControllerMockRecorder {
	return m.recorder
}

// LoginHandlerController mocks base method.
func (m *MockOAuthController) LoginHandlerController(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LoginHandlerController", arg0)
}

// LoginHandlerController indicates an expected call of LoginHandlerController.
func (mr *MockOAuthControllerMockRecorder) LoginHandlerController(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginHandlerController", reflect.TypeOf((*MockOAuthController)(nil).LoginHandlerController), arg0)
}

// OAuthCallbackHandler mocks base method.
func (m *MockOAuthController) OAuthCallbackHandler(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OAuthCallbackHandler", arg0)
}

// OAuthCallbackHandler indicates an expected call of OAuthCallbackHandler.
func (mr *MockOAuthControllerMockRecorder) OAuthCallbackHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OAuthCallbackHandler", reflect.TypeOf((*MockOAuthController)(nil).OAuthCallbackHandler), arg0)
}

// OAuthHanderController mocks base method.
func (m *MockOAuthController) OAuthHanderController(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OAuthHanderController", arg0)
}

// OAuthHanderController indicates an expected call of OAuthHanderController.
func (mr *MockOAuthControllerMockRecorder) OAuthHanderController(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OAuthHanderController", reflect.TypeOf((*MockOAuthController)(nil).OAuthHanderController), arg0)
}

// MockOAuthUseCase is a mock of OAuthUseCase interface.
type MockOAuthUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockOAuthUseCaseMockRecorder
}

// MockOAuthUseCaseMockRecorder is the mock recorder for MockOAuthUseCase.
type MockOAuthUseCaseMockRecorder struct {
	mock *MockOAuthUseCase
}

// NewMockOAuthUseCase creates a new mock instance.
func NewMockOAuthUseCase(ctrl *gomock.Controller) *MockOAuthUseCase {
	mock := &MockOAuthUseCase{ctrl: ctrl}
	mock.recorder = &MockOAuthUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOAuthUseCase) EXPECT() *MockOAuthUseCaseMockRecorder {
	return m.recorder
}

// LoginHandlerUseCase mocks base method.
func (m *MockOAuthUseCase) LoginHandlerUseCase(ctx context.Context, user dtos.OAuthRequest) *models.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginHandlerUseCase", ctx, user)
	ret0, _ := ret[0].(*models.ErrorResponse)
	return ret0
}

// LoginHandlerUseCase indicates an expected call of LoginHandlerUseCase.
func (mr *MockOAuthUseCaseMockRecorder) LoginHandlerUseCase(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginHandlerUseCase", reflect.TypeOf((*MockOAuthUseCase)(nil).LoginHandlerUseCase), ctx, user)
}

// SaveSession mocks base method.
func (m *MockOAuthUseCase) SaveSession(ctx context.Context, user dtos.OAuthRequest) *models.ErrorResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveSession", ctx, user)
	ret0, _ := ret[0].(*models.ErrorResponse)
	return ret0
}

// SaveSession indicates an expected call of SaveSession.
func (mr *MockOAuthUseCaseMockRecorder) SaveSession(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveSession", reflect.TypeOf((*MockOAuthUseCase)(nil).SaveSession), ctx, user)
}

// MockOAuthService is a mock of OAuthService interface.
type MockOAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockOAuthServiceMockRecorder
}

// MockOAuthServiceMockRecorder is the mock recorder for MockOAuthService.
type MockOAuthServiceMockRecorder struct {
	mock *MockOAuthService
}

// NewMockOAuthService creates a new mock instance.
func NewMockOAuthService(ctrl *gomock.Controller) *MockOAuthService {
	mock := &MockOAuthService{ctrl: ctrl}
	mock.recorder = &MockOAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOAuthService) EXPECT() *MockOAuthServiceMockRecorder {
	return m.recorder
}

// GenerateAccessToken mocks base method.
func (m *MockOAuthService) GenerateAccessToken(ctx context.Context, refreshToken string) (string, *models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateAccessToken", ctx, refreshToken)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*models.ErrorResponse)
	return ret0, ret1
}

// GenerateAccessToken indicates an expected call of GenerateAccessToken.
func (mr *MockOAuthServiceMockRecorder) GenerateAccessToken(ctx, refreshToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateAccessToken", reflect.TypeOf((*MockOAuthService)(nil).GenerateAccessToken), ctx, refreshToken)
}

// OAuthTokenValidator mocks base method.
func (m *MockOAuthService) OAuthTokenValidator(token string, ctx context.Context) (*models.JWTCustome, *models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OAuthTokenValidator", token, ctx)
	ret0, _ := ret[0].(*models.JWTCustome)
	ret1, _ := ret[1].(*models.ErrorResponse)
	return ret0, ret1
}

// OAuthTokenValidator indicates an expected call of OAuthTokenValidator.
func (mr *MockOAuthServiceMockRecorder) OAuthTokenValidator(token, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OAuthTokenValidator", reflect.TypeOf((*MockOAuthService)(nil).OAuthTokenValidator), token, ctx)
}

// RefreshTokenValidator mocks base method.
func (m *MockOAuthService) RefreshTokenValidator(refreshToken string, ctx context.Context) (*models.JWTCustome, *models.ErrorResponse) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTokenValidator", refreshToken, ctx)
	ret0, _ := ret[0].(*models.JWTCustome)
	ret1, _ := ret[1].(*models.ErrorResponse)
	return ret0, ret1
}

// RefreshTokenValidator indicates an expected call of RefreshTokenValidator.
func (mr *MockOAuthServiceMockRecorder) RefreshTokenValidator(refreshToken, ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTokenValidator", reflect.TypeOf((*MockOAuthService)(nil).RefreshTokenValidator), refreshToken, ctx)
}
