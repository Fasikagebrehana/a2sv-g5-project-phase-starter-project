package userusecase

import (
	"blogs/config"
	"blogs/domain"
)

func (u *UserUsecase) LoginUser(usernameoremail string, password string) (string, string, error) {
	user, err := u.UserRepo.GetUserByUsernameorEmail(usernameoremail)
	if err != nil {
		return "", "", config.ErrIncorrectPassword
	}

	if !user.IsVerified {
		return "", "", config.ErrUserNotVerified
	}

	// Compare the hashed password
	err = config.ComparePassword(user.Password, password)
	if err != nil {
		return "", "", config.ErrIncorrectPassword
	}

	// Generate access token
	accessToken, _, err := config.GenerateToken(
		&domain.LoginClaims{
			Username: user.Username,
			Role:     user.Role,
			Type:     "access",
		}, "access")

	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshToken, tokenEntry, err := config.GenerateToken(
		&domain.LoginClaims{
			Username: user.Username,
			Role:     user.Role,
			Type:     "refresh",
		}, "refresh")

	if err != nil {
		return "", "", err
	}

	// Save the refresh token in the repository
	err = u.UserRepo.InsertToken(tokenEntry)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
