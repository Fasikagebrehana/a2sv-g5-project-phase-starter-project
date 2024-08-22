package usecase

import (
	infrastructure "astu-backend-g1/Infrastructure"
	"astu-backend-g1/domain"
	"crypto/rand"
	"errors"
	"math/big"
	"time"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) (domain.UserUsecase, error) {
	return &userUsecase{userRepository: u}, nil
}

func (useCase *userUsecase) Get() ([]domain.User, error) {
	return useCase.userRepository.Get(domain.UserFilterOption{})
}

func (useCase *userUsecase) LoginUser(uname string, password string) (string, error) {
	user, err := useCase.GetByUsername(uname)
	if err != nil {
		return "", err
	}
	if user.IsActive == false {
		return "", errors.New("Account not activated")
	}
	accesstoken, refreshToken, err := infrastructure.GenerateToken(&user, password)
	if err != nil {
		return "", err
	}
	user.RefreshToken = refreshToken
	useCase.userRepository.Update(user.ID, domain.User{RefreshToken: refreshToken})

	return accesstoken, nil
}

func (useCase *userUsecase) Logout(email string) error {
	user, err := useCase.GetByEmail(email)
	if err != nil {
		return err
	}
	user.RefreshToken = ""
	useCase.userRepository.Update(user.ID, user)
	return nil
}

func (useCase *userUsecase) ForgetPassword(email string) (string, error) {
	user, err := useCase.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if user.IsActive == false {
		return "", errors.New("Account not activated")
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	confirmationToken := make([]byte, 64)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < 64; i++ {
		num, _ := rand.Int(rand.Reader, charsetLength)
		confirmationToken[i] = charset[num.Int64()]
	}

	user.VerifyToken = string(confirmationToken)

	expirationTime := time.Now().Add(2 * time.Hour)
	useCase.userRepository.Update(user.ID, domain.User{VerifyToken: string(confirmationToken), ExpirationDate: expirationTime})
	link := "http://localhost:8000/users/resetPassword/?email=" + user.Email + "&token=" + string(confirmationToken)
	err = infrastructure.SendEmail(user.Email, "Password Reset", "This is the password reset link: ", link)
	if err != nil {
		return "", err
	}

	return "Password reset token sent to your email", nil
}

func (useCase *userUsecase) ResetPassword(email string, token string, password string) (string, error) {
	user, err := useCase.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if !user.IsActive {
		return "", errors.New("Account not activated")
	}
	if user.VerifyToken == token {
		if user.ExpirationDate.Before(time.Now()) {
			return "Token has expired", errors.New("Token expired")
		}
		user.Password, _ = infrastructure.PasswordHasher(password)
		_, err := useCase.userRepository.Update(user.ID, domain.User{Password: user.Password})
		if err != nil {
			return "password has not been updated", err
		}
		return "Password reset successful", nil
	}
	return "Invalid token", errors.New("Invalid token")
}

func (useCase *userUsecase) GetByID(userID string) (domain.User, error) {
	filter := domain.UserFilter{UserId: userID}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	return users[0], err
}

func (useCase *userUsecase) GetByUsername(username string) (domain.User, error) {
	filter := domain.UserFilter{Username: username}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	return users[0], err
}

// function for account verification

func (useCase *userUsecase) AccountVerification(uemail string, confirmationToken string) error {
	user, err := useCase.GetByEmail(uemail)
	if err != nil {
		return err
	}
	if user.VerifyToken == confirmationToken {
		_, err := useCase.userRepository.Update(user.ID, domain.User{IsActive: true})
		return err
	} else {
		return errors.New("invalid token")
	}
}

func (useCase *userUsecase) GetByEmail(email string) (domain.User, error) {
	filter := domain.UserFilter{Email: email}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	return users[0], err
}

func (useCase *userUsecase) Create(u *domain.User) (domain.User, error) {
	u.Password, _ = infrastructure.PasswordHasher(u.Password)
	u.IsActive = false
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	confirmationToken := make([]byte, 64)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < 64; i++ {
		num, _ := rand.Int(rand.Reader, charsetLength)
		confirmationToken[i] = charset[num.Int64()]
	}
	u.VerifyToken = string(confirmationToken)
	nUser, err := useCase.userRepository.Create(u)
	if !nUser.IsAdmin {
		link := "`http://localhost:8000/`users/accountVerification/?email=" + u.Email + "&token=" + string(confirmationToken)
		err := infrastructure.SendEmail(u.Email, "Registration Confirmation", "This sign up Confirmation email to verify: ", link)
		if err != nil {
			return nUser, err
		}
	}
	return nUser, err
}

func (useCase *userUsecase) Update(userId string, updateData domain.User) (domain.User, error) {
	return useCase.userRepository.Update(userId, updateData)
}

func (useCase *userUsecase) Delete(userId string) error {
	return useCase.userRepository.Delete(userId)
}
