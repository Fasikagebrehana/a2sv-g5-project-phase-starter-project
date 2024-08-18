package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username   string             `bson:"username" json:"username" validate:"required,min=3,max=30"`
	Name       string             `bson:"name" json:"name" validate:"required"`
	Email      string             `bson:"email" json:"email" validate:"required,email"`
	Password   string             `bson:"password" json:"password"`
	Role       Role               `bson:"role" json:"role"`
	IsVerified bool               `bson:"is_verified" json:"is_verified"`
}
