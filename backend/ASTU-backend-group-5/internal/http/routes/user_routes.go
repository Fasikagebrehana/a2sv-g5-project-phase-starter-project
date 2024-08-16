package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterUserRoutes(collection *mongo.Collection, router *gin.Engine) {
	userHandler := InstantaiteUserHandler(collection)
	adminRoute := router.Group("/api/v1/accounts")
	{
		adminRoute.GET("/me", userHandler.GetUser)
		adminRoute.DELETE("/me", userHandler.DeleteUser)
		adminRoute.PUT("/me", userHandler.UpdateUser)

	}

}