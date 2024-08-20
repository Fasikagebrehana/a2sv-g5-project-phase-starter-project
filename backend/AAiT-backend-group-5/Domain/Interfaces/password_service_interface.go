package interfaces

import (
	"context"

	models "github.com/aait.backend.g5.main/backend/Domain/Models"
)

type PasswordService interface {
	EncryptPassword(password string) (string, error)
	ValidatePassword(password string, hashedPassword string) bool
	ValidatePasswordStrength(password, email, name, username string) *models.ErrorResponse
}

type PasswordUsecase interface {
	GenerateResetURL(ctx context.Context, email string) (string, *models.ErrorResponse)
	SendResetEmail(ctx context.Context, email string, resetURL string) *models.ErrorResponse
	SetPassword(ctx context.Context, shortURlCode string, password string) *models.ErrorResponse
}
