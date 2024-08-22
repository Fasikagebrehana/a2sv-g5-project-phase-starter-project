package router

import (
	infrastructure "astu-backend-g1/Infrastructure"
	"astu-backend-g1/delivery/controllers"

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

// @title Blog API in Golang
func (gr *MainRouter) GinBlogRouter() {
	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("blogs/", gr.blogController.HandleGetAllBlogs)
	router.GET("blogs/popular", gr.blogController.HandleGetPopularBlog)
	router.GET("blogs/filter", gr.blogController.HandleFilterBlogs)
	router.GET("blogs/:blogId", gr.blogController.HandleGetBlogById)
	userrouter := router.Group("/users")
	{

		userrouter.POST("/register", gr.handler.Register)
		userrouter.GET("/accountVerification", gr.handler.AccountVerification)
		userrouter.POST("/login", gr.handler.LoginUser)
		userrouter.GET("/forgetPassword", gr.handler.ForgetPassword)
		userrouter.POST("/resetPassword", gr.handler.ResetPassword)
		userrouter.POST("/:uid/refresh", gr.handler.RefreshAccessToken)
	}
	blogRouter := router.Group("/blogs")
	blogRouter.Use(gr.authController.AuthenticationMiddleware())
	{
		blogRouter.POST("/", gr.authController.USERMiddleware(), gr.blogController.HandleCreateBlog)
		blogRouter.PATCH("/:blogId", gr.authController.ADMINMiddleware(), gr.blogController.HandleBlogUpdate)
		blogRouter.DELETE("/:blogId", gr.authController.ADMINMiddleware(), gr.blogController.HandleBlogDelete)
		blogRouter.POST("/:blogId/interact/:type", gr.authController.USERMiddleware(), gr.blogController.HandleBlogLikeOrDislike)

		// TODO: check if there is a blog with such id
		commentRouter := blogRouter.Group("/:blogId/comments")
		commentRouter.Use(gr.authController.USERMiddleware())
		{
			commentRouter.GET("/", gr.blogController.HandleGetAllComments)
			commentRouter.POST("/", gr.blogController.HandleCommentOnBlog)
			commentRouter.GET("/:commentId", gr.blogController.HandleGetCommentById)
			commentRouter.POST("/:commentId/interact/:type", gr.blogController.HandleCommentLikeOrDislike)

			repliesRouter := commentRouter.Group("/:commentId/replies")
			repliesRouter.Use(gr.authController.USERMiddleware())
			{
				repliesRouter.GET("/", gr.blogController.HandleGetAllRepliesForComment)
				repliesRouter.POST("/", gr.blogController.HandleReplyOnComment)
				repliesRouter.GET("/:replyId", gr.blogController.HandleGetReplyById)
				// todo: test the below functions
				repliesRouter.POST("/:replyId/interact/:type", gr.blogController.HandleReplyLikeOrDislike)
			}
		}

	}
	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API get"})
		ctx.Abort()
		return
	})
	router.POST("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API create"})
		ctx.Abort()
		return
	})
	router.DELETE("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API delete"})
		ctx.Abort()
		return
	})
	router.PATCH("", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Welcome to Blog API patch"})
		ctx.Abort()
		return
	})
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "Not such route"})
		ctx.Abort()
		return
	})
	router.Run(":8000")
}
