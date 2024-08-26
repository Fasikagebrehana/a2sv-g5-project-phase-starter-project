package router

import (
	"blogs/config"
	"blogs/delivery/controller/blogcontroller"
	"blogs/delivery/controller/usercontroller"
	"blogs/delivery/middleware"
	"blogs/domain"
	"blogs/repository"
	"blogs/usecase/blogusecase"
	"blogs/usecase/userusecase"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func getBlogController(database *mongo.Database, cache domain.Cache) *blogcontroller.BlogController {
	blogRepository := repository.NewBlogRepository(database, cache)
	tagRepository := repository.NewTagRepository(database)
	blogUsecase := blogusecase.NewBlogUsecase(blogRepository, tagRepository)
	blogController := blogcontroller.NewBlogController(blogUsecase)
	return blogController
}

func getUserController(database *mongo.Database, cache domain.Cache) *usercontroller.UserController {
	userRepository := repository.NewUserRepository(database, cache)
	authRepository := repository.NewOAuthRepository(database)
	userUsecase := userusecase.NewUserUsecase(userRepository, authRepository)
	userController := usercontroller.NewUserController(userUsecase)

	err := userUsecase.AddRoot()
	if err != nil {
		panic(err)
	}

	return userController
}

func publicRouter(router *gin.Engine, userController *usercontroller.UserController) {
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.LoginUser)
	router.POST("/users/forgot-password", userController.ForgotPassword)

	router.GET("/users/verify", userController.VerifyUser)
	router.GET("/users/reset-password", userController.ResetPassword)

	router.GET("/oauth2/login/google", userController.GoogleLogin)
	router.GET("/oauth2/callback/google", userController.GoogleCallback)
}

func protectedRouter(router *gin.Engine, userController *usercontroller.UserController) {
	router.GET(
		"/tokens/refresh",
		middleware.AuthMiddleware("refresh"),
		userController.RefreshToken,
	)
}

func privateUserRouter(router *gin.RouterGroup, userController *usercontroller.UserController) {
	router.GET("/users/:username", userController.GetUserByUsername)
	router.PATCH("/users", userController.UpdateProfile)
	router.PATCH("/users/promote", userController.PromoteUser)
	router.POST("/users/logout", userController.LogoutUser)
	router.DELETE("/users", userController.DeleteUser)
	router.PATCH("/users/change-password", userController.ChangePassword)
}

func privateBlogRouter(router *gin.RouterGroup, blogController *blogcontroller.BlogController) {
	router.POST("/blogs", blogController.InsertBlog)
	router.GET("/blogs", blogController.GetBlogs)
	router.GET("/blogs/:id", blogController.GetBlogByID)
	router.PUT("/blogs/:id", blogController.UpdateBlogByID)
	router.DELETE("/blogs/:id", blogController.DeleteBlogByID)

	router.POST("/blogs/:id/comments", blogController.AddComment)
	router.GET("/blogs/:id/comments", blogController.GetBlogComments)
	router.DELETE("/blogs/:id/comments/:commentid", blogController.DeleteComment)

	router.POST("/blogs/:id/likes", blogController.AddLike)
	router.DELETE("/blogs/:id/likes", blogController.RemoveLike)
	router.GET("/blogs/:id/likes", blogController.GetBlogLikes)

	router.POST("/blogs/views", blogController.AddView)

	router.GET("/blogs/search", blogController.SearchBlog)

	router.GET("/blogs/tags", blogController.GetTags)
	router.POST("/blogs/tags", blogController.InsertTag)
	router.DELETE("/blogs/tags", blogController.RemoveTags)

	router.POST("/blogs/generate", blogController.GenerateContent)
	
}

func SetupRouter(mongoClient *mongo.Client) *gin.Engine {
	cache := config.NewRedisCache()
	router := gin.Default()

	// Secure Headers Configuration
	secureMiddleware := secure.New(secure.Config{
		SSLRedirect:           true,
		STSPreload:            true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'; script-src 'self'; object-src 'none';",
		ReferrerPolicy:        "no-referrer",
		IsDevelopment:         true,
		BadHostHandler:        func(*gin.Context) {},
	})

	// Apply secure middleware to the router
	router.Use(secureMiddleware)

	// CORS Configuration
	corsConfig := cors.Config{
		AllowOrigins:     []string{"https://trusteddomain.com", "https://anothertrusteddomain.com"}, // Adjust based on your needs
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	// Apply CORS middleware with custom configuration
	router.Use(cors.New(corsConfig))

	database := mongoClient.Database("blog")
	blogController := getBlogController(database, cache)
	userController := getUserController(database, cache)

	publicRouter(router, userController)
	protectedRouter(router, userController)

	privateRouter := router.Group("")
	privateRouter.Use(middleware.AuthMiddleware("access"))

	privateUserRouter(privateRouter, userController)
	privateBlogRouter(privateRouter, blogController)

	return router
}
