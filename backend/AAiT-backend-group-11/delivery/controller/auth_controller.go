package controller

import (
	"backend-starter-project/bootstrap"
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthController struct {
	AuthService 		interfaces.AuthenticationService
	TokenRepo           interfaces.RefreshTokenRepository
	Env 				*bootstrap.Env
}

func (ac *AuthController)RegisterUser(c *gin.Context)  {

	var userRequest entities.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.User{
		ID: primitive.NewObjectID(),
		Username: userRequest.Username,
		Email: userRequest.Email,
		Password: userRequest.Password,
		IsVerified: false,
		Role: "user",
		Profile: entities.Profile{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := ac.AuthService.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdUser})

}

func (ac *AuthController) VerifyEmail(c *gin.Context)  {

	var emailVerification entities.EmailVerificationRequest

	if err := c.ShouldBindJSON(&emailVerification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ac.AuthService.VerifyEmail(emailVerification.Email, emailVerification.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}
