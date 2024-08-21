package usecase

import (
	"context"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/emailutil"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/internal/tokenutil"
)

type signupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase)GetUserById(c context.Context, userId string) (*domain.User, error) {
	user,err:=su.userRepository.GetUserById(c,userId)
	return user,err
}
func (su *signupUsecase) ActivateUser(c context.Context, userID string) error {
	_, err := su.userRepository.ActivateUser(c, userID)
	return err
}
func (su *signupUsecase) Create(c context.Context, user *domain.User) (*domain.User,error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.CreateUser(ctx, user)
}

func (su *signupUsecase) IsOwner(ctx context.Context) (bool,error) {
	result,err:=su.userRepository.IsOwner(ctx)
	return result,err
}
func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	user, err := su.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return *user, nil
}

func (su *signupUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (su *signupUsecase) CreateVerificationToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateVerificationToken(user, secret, expiry)
}

func (su *signupUsecase) SendVerificationEmail(recipientEmail string,encodedToken string,env *bootstrap.Env) (err error) {
return emailutil.SendVerificationEmail(recipientEmail,encodedToken,env)
}