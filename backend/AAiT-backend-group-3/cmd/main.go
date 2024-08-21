package main

import (
	"AAIT-backend-group-3/internal/delivery/controllers"
	"AAIT-backend-group-3/internal/delivery/routers"
	"AAIT-backend-group-3/internal/infrastructures/database"
	"AAIT-backend-group-3/internal/infrastructures/middlewares"
	"AAIT-backend-group-3/internal/infrastructures/services"
	"AAIT-backend-group-3/internal/usecases"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"AAIT-backend-group-3/internal/repositories/implementation"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbName := os.Getenv("DB_NAME")
	secretKey := os.Getenv("SECRET_KEY")
	smtpPortStr := os.Getenv("SMTP_PORT")
	userName := os.Getenv("USERNAME")
	smtpHost := os.Getenv("SMTP_HOST")
	passWord := os.Getenv("PASSWORD")
	
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatalf("Invalid SMTP_PORT: %v", err)
	}

	dbClient, err := database.NewMongoDBClient(context.Background(), dbName)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB!", dbClient.Database.Name())


	//services
	emailSvc := services.NewEmailService(smtpHost, smtpPort, userName, passWord)
	passSvc := services.NewPasswordService()
	validationSvc := services.NewValidationService()
	jwtSvc := services.NewJWTService(secretKey)
	cacheSvc := services.NewCacheService("localhost:6379", "", 0)

	//repositories
	userRepo := repositories.NewMongoUserRepository(dbClient.Database, "users", cacheSvc)
	otpRepo := repositories.NewMongoOtpRepository(dbClient.Database, "otps")
	blogRepo := repositories.NewMongoBlogRepository(dbClient.Database, "blogs")
	// commentRepo := repositories.NewMongoCommentRepository(dbClient.Database, "comments")


	//middlewares
	authMiddleware := middlewares.NewAuthMiddleware(jwtSvc, cacheSvc)


	//usecases
	userUsecase := usecases.NewUserUsecase(userRepo, passSvc, validationSvc, emailSvc, jwtSvc)
	otpUsecase := usecases.NewOtpUseCase(otpRepo, userRepo, emailSvc, passSvc, "http://localhost:8080", validationSvc)
	blogService := usecases.NewBlogUsecase(blogRepo, cacheSvc)
	// commentService := service.NewCommentService(commentRepo)


	// controllers
	userController := controllers.NewUserController(userUsecase)
	otpController := controllers.NewOTPController(otpUsecase)
	blogController := controllers.NewBlogController(blogService)
	// commentController := delivery.NewCommentController(commentService)

	router := gin.New()
	router.Use(gin.Logger())


	// routers
	routers.CreateUserRouter(router, userController, otpController, authMiddleware)
	routers.CreateBlogRouter(router, blogController, authMiddleware)
	if err := router.Run(":" + os.Getenv("PORT")); err!= nil{
		log.Fatal(err)
	}
	fmt.Println("server connnected")
}
