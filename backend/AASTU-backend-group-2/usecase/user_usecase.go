package usecase

import (
	"blog_g2/domain"
	"context"

	"time"
)

type UserUsecase struct {
	UserRepo       domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(Userrepo domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &UserUsecase{
		UserRepo:       Userrepo,
		contextTimeout: timeout,
	}

}

func (uuse *UserUsecase) RegisterUser(c context.Context, user domain.User) error {
	return nil
}

func (uuse *UserUsecase) LoginUser(c context.Context, user domain.User) (string, error) {
	return "", nil
}

func (uuse *UserUsecase) ForgotPassword(c context.Context, email string) error {
	return nil
}

func (uuse *UserUsecase) LogoutUser(c context.Context) error {
	return nil
}

func (uuse *UserUsecase) PromoteDemoteUser(c context.Context, userid string) error {
	return nil
}
