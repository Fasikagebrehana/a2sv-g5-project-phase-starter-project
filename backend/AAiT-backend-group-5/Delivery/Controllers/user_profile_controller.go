package controllers

import (
	"fmt"
	"net/http"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	"github.com/gin-gonic/gin"
)

type UserProfileController struct {
	UserProfileUC interfaces.UserProfileUpdateUsecase
}

func NewUserProfileController(userProfileUC interfaces.UserProfileUpdateUsecase) *UserProfileController {
	return &UserProfileController{
		UserProfileUC: userProfileUC,
	}
}

func (userProfileController *UserProfileController) ProfileUpdate(ctx *gin.Context) {
	var updatedUser *dtos.ProfileUpdateRequest
	userID := ctx.GetString("id")
	fmt.Println(userID)

	err := ctx.ShouldBind(&updatedUser)
	if err != nil {
		fmt.Println(updatedUser)
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "invalid request"})
		return
	}

	e := userProfileController.UserProfileUC.UpdateUserProfile(ctx, userID, updatedUser)
	if e != nil {
		ctx.IndentedJSON(e.Code, gin.H{"error": e.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user profile successfully updated"})
}
