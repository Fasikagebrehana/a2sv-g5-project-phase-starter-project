package usecase

import (
	"context"
	"testing"

	"github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/aait.backend.g5.main/backend/Mocks"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type SignupUsecaseTestSuite struct {
	suite.Suite
	repositoryMock   *mocks.MockUserRepository
	emailServiceMock *mocks.MockEmailService
	jwtServiceMock   *mocks.MockJwtService
	urlServiceMock   *mocks.MockURLService
	signupUsecase    interfaces.SignupUsecase
	ctrl             *gomock.Controller
}

func (suite *SignupUsecaseTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repositoryMock = mocks.NewMockUserRepository(suite.ctrl)
	suite.emailServiceMock = mocks.NewMockEmailService(suite.ctrl)
	suite.jwtServiceMock = mocks.NewMockJwtService(suite.ctrl)
	suite.urlServiceMock = mocks.NewMockURLService(suite.ctrl)

	suite.signupUsecase = usecases.NewSignupUsecase(
		suite.repositoryMock,
		suite.emailServiceMock,
		suite.jwtServiceMock,
		suite.urlServiceMock,
	)
}

func (suite *SignupUsecaseTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *SignupUsecaseTestSuite) TestCreateUser_Success() {
	ctx := context.Background()
	user := &models.User{Username: "newuser", Email: "newuser@example.com"}
	token := "token"
	url := "http://verification.url"

	suite.repositoryMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, user.Username, user.Email).
		Return(nil, nil)

	suite.emailServiceMock.
		EXPECT().
		IsValidEmail(user.Email).
		Return(true)

	suite.jwtServiceMock.
		EXPECT().
		CreateURLToken(*user, 3600).
		Return(token, nil)

	suite.urlServiceMock.
		EXPECT().
		GenerateURL(token, "confirmRegistration").
		Return(url, nil)

	suite.emailServiceMock.
		EXPECT().
		SendEmail(user.Email, "Email Verification", gomock.Any()).
		Return(nil)

	err := suite.signupUsecase.CreateUser(ctx, user)
	suite.Nil(err)
}

func (suite *SignupUsecaseTestSuite) TestCreateUser_UserExists() {
	ctx := context.Background()
	user := &models.User{Username: "existinguser", Email: "existinguser@example.com"}

	suite.repositoryMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, user.Username, user.Email).
		Return(user, nil)

	err := suite.signupUsecase.CreateUser(ctx, user)
	suite.Equal(models.BadRequest("User already exists"), err)
}

func (suite *SignupUsecaseTestSuite) TestCreateUser_InvalidEmail() {
	ctx := context.Background()
	user := &models.User{Username: "newuser", Email: "invalidemail"}

	suite.repositoryMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, user.Username, user.Email).
		Return(nil, nil)

	suite.emailServiceMock.
		EXPECT().
		IsValidEmail(user.Email).
		Return(false)

	err := suite.signupUsecase.CreateUser(ctx, user)
	suite.Equal(models.BadRequest("Invalid Email"), err)
}

func (suite *SignupUsecaseTestSuite) TestCreateUser_TokenGenerationError() {
	ctx := context.Background()
	user := &models.User{Username: "newuser", Email: "newuser@example.com"}

	suite.repositoryMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, user.Username, user.Email).
		Return(nil, nil)

	suite.emailServiceMock.
		EXPECT().
		IsValidEmail(user.Email).
		Return(true)

	suite.jwtServiceMock.
		EXPECT().
		CreateURLToken(*user, 3600).
		Return("", models.InternalServerError("Error while creating token"))

	suite.urlServiceMock.
		EXPECT().
		GenerateURL("", "confirmRegistration").
		Return("", models.InternalServerError("Error while creating url")).
		Times(0)
	err := suite.signupUsecase.CreateUser(ctx, user)
	suite.Equal(models.InternalServerError("Error while creating token"), err)
}

func (suite *SignupUsecaseTestSuite) TestCreateUser_EmailSendingError() {
	ctx := context.Background()
	user := &models.User{Username: "newuser", Email: "newuser@example.com"}
	token := "token"
	url := "http://verification.url"

	suite.repositoryMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, user.Username, user.Email).
		Return(nil, nil)

	suite.emailServiceMock.
		EXPECT().
		IsValidEmail(user.Email).
		Return(true)

	suite.jwtServiceMock.
		EXPECT().
		CreateURLToken(*user, 3600).
		Return(token, nil)

	suite.urlServiceMock.
		EXPECT().
		GenerateURL(token, "confirmRegistration").
		Return(url, nil)

	suite.emailServiceMock.
		EXPECT().
		SendEmail(user.Email, "Email Verification", gomock.Any()).
		Return(models.InternalServerError("Error while sending email"))

	err := suite.signupUsecase.CreateUser(ctx, user)
	suite.Equal(models.InternalServerError("Error while sending email"), err)
}

func TestSignupUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(SignupUsecaseTestSuite))
}
