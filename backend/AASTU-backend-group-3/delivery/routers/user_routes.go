package routers

// import (
// 	"group3-blogApi/delivery/controllers/authController"
// 	"group3-blogApi/delivery/controllers/userController"

// 	"github.com/gin-gonic/gin"
// )

// func SetUpUser(router *gin.Engine) {
// 	//user routes
// 	user := router.Group("/user")

// 	{
// 		user.GET("/profile", authMiddleware, userController.Profile)
// 		user.PUT("/update", authMiddleware, userController.Update)
// 		user.POST("/upload-image", authMiddleware, userController.UploadImage)
// 		user.POST("/logout", authMiddleware, authController.Logout)
// 		user.POST("/reset-password", authMiddleware, authController.ResetPassword)
// 		user.POST("/refresh-token", authController.RefreshToken)
// 	}
// }
