package controllers

import (
	"net/http"
	"log"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

type LogoutController struct {
	LogoutUsecase interfaces.LogoutUsecase
	JwtService    interfaces.JwtService
}

func NewLogoutController(logoutUsecase interfaces.LogoutUsecase, jwtService interfaces.JwtService) *LogoutController {
	return &LogoutController{
		LogoutUsecase: logoutUsecase,
		JwtService:    jwtService,
	}
}

func (logoutController *LogoutController) Logout(ctx *gin.Context) {
	// get the userId from the context
	userId := ctx.GetString("id")

	log.Println("userId", userId)

	e := logoutController.LogoutUsecase.LogoutUser(ctx, userId)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
