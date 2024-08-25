package router

import (
	"astu-backend-g1/delivery/controllers"
	"astu-backend-g1/infrastructure"

	_ "astu-backend-g1/delivery/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type MainRouter struct {
	blogController controllers.BlogController
	authController infrastructure.GeneralAuthorizationController
	handler        controllers.UserController
}

func NewMainRouter(uc controllers.UserController, bc controllers.BlogController, authc infrastructure.GeneralAuthorizationController) *MainRouter {
	return &MainRouter{
		blogController: bc,
		authController: authc,
		handler:        uc,
	}
}

func (gr *MainRouter) GinBlogRouter() {

	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userrouter := router.Group("/users")
	{
		userrouter.GET("/", gr.handler.GetUsers)
		userrouter.GET("/:id", gr.handler.GetUserByID)
		userrouter.PUT("/:id", gr.handler.UpdateUser)
		userrouter.POST("/register", gr.handler.Register)
		userrouter.GET("/accountVerification", gr.handler.AccountVerification)
		userrouter.POST("/login", gr.handler.LoginUser)
		userrouter.GET("/forgetPassword", gr.handler.ForgetPassword)
		userrouter.POST("/resetPassword", gr.handler.ResetPassword)
		userrouter.GET("/logout", gr.handler.LogoutUser)
		userrouter.POST("/:uid/refresh", gr.handler.RefreshAccessToken)
	}
	router.GET("blogs/", gr.blogController.HandleGetAllBlogs)
	router.GET("blogs/popular", gr.blogController.HandleGetPopularBlog)
	router.GET("blogs/filter", gr.blogController.HandleFilterBlogs)
	router.GET("blogs/:blogId", gr.blogController.HandleGetBlogById)

	blogRouter := router.Group("/blogs")
	blogRouter.Use(gr.authController.AuthenticationMiddleware())
	{
		blogRouter.POST("/", gr.authController.USERMiddleware(), gr.blogController.HandleCreateBlog)
		blogRouter.PATCH("/:blogId", gr.authController.OWNERMiddleware(), gr.blogController.HandleBlogUpdate)
		blogRouter.DELETE("/:blogId", gr.authController.OWNERMiddleware(), gr.blogController.HandleBlogDelete)
		blogRouter.POST("/:blogId/:type", gr.authController.USERMiddleware(), gr.blogController.HandleBlogLikeOrDislike)
		commentRouter := blogRouter.Group("/:blogId/comments")
		commentRouter.Use(gr.authController.USERMiddleware())
		{
			commentRouter.GET("/", gr.blogController.HandleGetAllComments)
			commentRouter.POST("/", gr.blogController.HandleCommentOnBlog)
			commentRouter.GET("/:commentId", gr.blogController.HandleGetCommentById)
			commentRouter.POST("/:commentId/:type", gr.blogController.HandleCommentLikeOrDislike)

			repliesRouter := commentRouter.Group("/:commentId/replies")
			repliesRouter.Use(gr.authController.USERMiddleware())
			{
				repliesRouter.GET("/", gr.blogController.HandleGetAllRepliesForComment)
				repliesRouter.POST("/", gr.blogController.HandleReplyOnComment)
				repliesRouter.GET("/:replyId", gr.blogController.HandleGetReplyById)
				repliesRouter.POST("/:replyId/:type", gr.blogController.HandleReplyLikeOrDislike)
			}
		}
	}
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API get"})
		ctx.Abort()
	})
	router.POST("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API create"})
		ctx.Abort()
	})
	router.DELETE("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API delete"})
		ctx.Abort()
	})
	router.PATCH("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API patch"})
		ctx.Abort()
	})
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "Not such route"})
		ctx.Abort()
	})
	router.Run(":8000")
}
