package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"AAIT-backend-group-3/internal/infrastructures/database"
	// "AAIT-backend-group-3/internal/repositories/implementation"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dbName := os.Getenv("DB_NAME")
	dbClient, err := database.NewMongoDBClient(context.Background(), dbName)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB!", dbClient.Database.Name())
	// userRepo := repositories.NewMongoUserRepository(dbClient.Database, "users")
	// blogRepo := repositories.NewMongoBlogRepository(dbClient.Database, "blogs")
	// commentRepo := repositories.NewMongoCommentRepository(dbClient.Database, "comments")
	// otpRepo := repositories.NewMongoOtpRepository(dbClient.Database, "otps")

	// Initialize services
	// userService := service.NewUserService(userRepo)
	// blogService := service.NewBlogService(blogRepo)
	// commentService := service.NewCommentService(commentRepo)

	// Initialize controllers
	// userController := delivery.NewUserController(userService)
	// blogController := delivery.NewBlogController(blogService)
	// commentController := delivery.NewCommentController(commentService)

	// Use the controllers for handling HTTP requests (assuming you have a router set up)
	fmt.Println("Application is running...")
}
