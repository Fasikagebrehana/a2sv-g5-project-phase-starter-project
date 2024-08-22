package user_controller

import (
	"blog-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *userController) Login(c *gin.Context) {
	var req domain.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := uc.userUsecase.Login(c.Request.Context(), req)
	if err != nil {
		if err == domain.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	// Set the refresh token in a secure cookie
	c.SetCookie(
		"refresh_token",
		resp.RefreshToken,
		uc.Env.RefreshTokenExpiryHour,
		"/",
		"localhost",
		false,
		true,
	)

	c.JSON(http.StatusOK, resp)
}
