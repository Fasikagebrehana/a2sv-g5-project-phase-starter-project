package main

// import (
// 	// domain "blogs/Domain"
// 	// domain "blogs/Domain"
// 	"fmt"
// 	"time"

// 	"github.com/golang-jwt/jwt/v4"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// func CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
// 	exp := time.Now().Add(time.Hour * time.Duration(expiry))

// 	// Create claims
// 	claims := &domain.JwtCustomClaims{
// 		UserName: user.Username,
// 		ID:       user.ID.Hex(),
// 		Role:     user.Role,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(exp), // Convert expiration time to *jwt.NumericDate
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	t, err := token.SignedString([]byte(secret))
// 	if err != nil {
// 		return "", err
// 	}
// 	return t, err
// }
// func main() {
// 	id, _ := primitive.ObjectIDFromHex("66c0a0dcdb2272faca4591ae")
// 	fmt.Println(CreateAccessToken(&domain.User{
// 		Username: "Full_Name",
// 		ID:       id,
// 		Role:     "user",
// 	}, "access_token_secret", 10))
// 	startDate := "2024-08-20T10:48:03+03:00"
// 	endDate := "2024-08-21T10:48:03+03:00"

// 	StartDate, err := time.Parse(time.RFC3339, startDate)
// 	if err != nil {
// 		fmt.Println("Error parsing start date:", err)
// 		return
// 	}

// 	EndDate, err := time.Parse(time.RFC3339, endDate)
// 	if err != nil {
// 		fmt.Println("Error parsing end date:", err)
// 		return
// 	}
// 	fmt.Println(StartDate, EndDate)
// }
