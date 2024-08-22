package main

import (
	infrastructure "astu-backend-g1/Infrastructure"
	"astu-backend-g1/delivery/controllers"
	_ "astu-backend-g1/delivery/docs"
	router "astu-backend-g1/delivery/routers"
	"astu-backend-g1/repository"
	usecase "astu-backend-g1/usecases"
	"context"

	"github.com/sv-tools/mongoifc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// @title TODO APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apiKey JWT
// @in header
// @name token
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
// @schemes http
func main() {
	clientOptions := options.Client().ApplyURI("mongodb://hundera:55969362@localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	blogCollections := client.Database("BlogAPI").Collection("Blogs")
	userCollections := client.Database("BlogAPI").Collection("Users")
	_ = client.Database("BlogAPI").Collection("Tokens")
	blogRepo := repository.NewBlogRepository(mongoifc.WrapCollection(blogCollections))
	blogUsecase := usecase.NewBlogUsecase(blogRepo)
	blogController := controllers.NewBlogController(*blogUsecase)
	authController := infrastructure.NewAuthController(blogRepo)
	userRepo := repository.NewUserRepository(mongoifc.WrapCollection(userCollections))
	userUsecase, err := usecase.NewUserUsecase(userRepo)
	if err != nil {
		panic(err)
	}
	UserController := controllers.NewUserController(userUsecase)
	Router := router.NewMainRouter(*UserController, *blogController, authController)
	Router.GinBlogRouter()
}
