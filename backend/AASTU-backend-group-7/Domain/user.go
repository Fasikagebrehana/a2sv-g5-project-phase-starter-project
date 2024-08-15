package Domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email" validate:"required"`
	Password string             `json:"password,omitempty" validate:"required"`
	Is_Admin bool               `json:"is_admin,omitempty" default:"false"`
}

// this could have been handled in a better way but i was too lazy to do it
type OmitedUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" bson:"_id,omitempty"`
	Email    string             `json:"email" validate:"required"`
	Password string             `json:"-"`
	Is_Admin bool               `json:"is_admin" default:"false"`
}
