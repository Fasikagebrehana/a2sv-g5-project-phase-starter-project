package jwtservice

import (
	Config "blogapp/Config"
	"blogapp/Domain"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateAccessToken(existingUser Domain.User) (string, error) {
	userclaims := &Domain.AccessClaims{
		ID:   existingUser.ID,
		Role: existingUser.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	}

	// Create a new JWT token with the user claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userclaims)

	// Ensure Config.JwtSecret is of type []byte
	jwtToken, err := token.SignedString([]byte(Config.JwtSecret))
	return jwtToken, err
}

func CreateRefreshToken(existingUser Domain.User) (refreshToken string, err error) {
	userclaims := &Domain.RefreshClaims{
		ID:   existingUser.ID,
		Role: existingUser.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
		},
	}

	// Create a new JWT token with the user claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userclaims)

	// Ensure Config.JwtSecret is of type []byte
	jwtToken, err := token.SignedString([]byte(Config.JwtSecret))
	return jwtToken, err
}

func VerifyRefreshToken(tokenString string) ( error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(Config.JwtSecret), nil
	})
	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*Domain.RefreshClaims)
	if !ok || !token.Valid {
		return errors.New("Invalid refresh token")
	}

	if time.Now().Unix() > claims.ExpiresAt {
		return errors.New("Refresh token has expired")
	}

	return nil
}
