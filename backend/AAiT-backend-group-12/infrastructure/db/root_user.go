package initdb

import (
	"blog_api/domain"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Creates a root user in the database with the given username and password
func CreateRootUser(db *mongo.Database, rootUsername string, rootPassword string, hashService domain.HashingServiceInterface) error {
	rootUser := domain.User{
		Username:   rootUsername,
		Email:      "root@root.root",
		Password:   rootPassword,
		Role:       "root",
		CreatedAt:  time.Now().Round(0),
		IsVerified: true,
	}

	hashedPwd, err := hashService.HashString(rootUser.Password)
	if err != nil {
		return fmt.Errorf("error hashing root user password: " + err.Error())
	}
	rootUser.Password = hashedPwd
	collection := db.Collection(domain.CollectionUsers)

	_, derr := collection.DeleteMany(context.Background(), bson.D{bson.E{Key: "role", Value: "root"}})
	if derr != nil {
		return fmt.Errorf("error clearing root users: " + derr.Error())
	}

	_, derr = collection.InsertOne(context.Background(), rootUser)
	if derr != nil {
		return fmt.Errorf("error creating root users: " + derr.Error())
	}

	res := collection.FindOneAndUpdate(context.Background(), bson.D{{Key: "username", Value: rootUsername}}, bson.D{{Key: "$unset", Value: bson.D{{Key: "verificationdata", Value: ""}}}})
	if res.Err() != nil {
		return domain.NewError(res.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	fmt.Println("Root user created successfully")

	return nil
}
