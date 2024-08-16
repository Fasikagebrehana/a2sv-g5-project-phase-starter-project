package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{
	ID				primitive.ObjectID		`bson:"_id"`
	Name		    *string			`json:"name" validate:"required,min=2,max=100"`
	Username		*string			`json:"username" validate:"required,min=2,max=100"`
	Password		*string			`json:"password" validate:"required,min=6"`
	Email			*string			`json:"email" validate:"email,required"`
	User_type		string			`json:"user_type"`
	Created_at		time.Time		`json:"created_at"`
	Updated_at		time.Time		`json:"updated_at"`
	User_id			string			`json:"user_id"`
}