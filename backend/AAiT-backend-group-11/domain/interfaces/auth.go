package interfaces

import  "backend-starter-project/domain/entities"

type AuthenticationService interface {
	RegisterUser(user *entities.User) (*entities.User, error)
	Login(emailOrUsername, password string) (*entities.RefreshToken,string, error)
	VerifyEmail(email, code string) error
	Logout(userId string) error
}

type PasswordResetService interface {
    RequestPasswordReset(email string) error
    ResetPassword(token, newPassword string) error
}