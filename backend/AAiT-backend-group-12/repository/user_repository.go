package repository

import (
	"blog_api/domain"
	"blog_api/domain/dtos"
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{collection: collection}
}

func (r *UserRepository) CreateUser(c context.Context, user *domain.User) domain.CodedError {
	_, err := r.collection.InsertOne(c, user)

	if mongo.IsDuplicateKeyError(err) && strings.Contains(err.Error(), "email") {
		return *domain.NewError("email already taken", domain.ERR_BAD_REQUEST)
	}

	if mongo.IsDuplicateKeyError(err) && strings.Contains(err.Error(), "username") {
		return *domain.NewError("username already taken", domain.ERR_BAD_REQUEST)
	}

	if err != nil {
		return *domain.NewError("error: failed to create user, "+err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

func (r *UserRepository) FindUser(c context.Context, user *domain.User) (domain.User, domain.CodedError) {
	var foundUser domain.User
	filter := bson.M{
		"$or": []bson.M{
			{"username": user.Username},
			{"email": user.Email},
		},
	}

	res := r.collection.FindOne(c, filter)
	if res.Err() == mongo.ErrNoDocuments {
		return foundUser, domain.NewError("user not found", domain.ERR_NOT_FOUND)
	}

	if res.Err() != nil {
		return foundUser, domain.NewError(res.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	err := res.Decode(&foundUser)
	if err != nil {
		return foundUser, domain.NewError(err.Error(), domain.ERR_INTERNAL_SERVER)
	}

	return foundUser, nil
}

func (r *UserRepository) SetRefreshToken(c context.Context, user *domain.User, newRefreshToken string) domain.CodedError {
	filter := bson.M{
		"$or": []bson.M{
			{"username": user.Username},
			{"email": user.Email},
		},
	}

	res := r.collection.FindOneAndUpdate(c, filter, bson.D{{
		Key: "$set", Value: bson.D{{Key: "refreshtoken", Value: newRefreshToken}},
	}})
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return domain.NewError("User not found", domain.ERR_NOT_FOUND)
	}

	if res.Err() != nil {
		return domain.NewError(res.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	return nil
}

func (r *UserRepository) UpdateUser(c context.Context, username string, user *dtos.UpdateUser) (map[string]string, domain.CodedError) {
	var updatedData = make(map[string]string)
	var updates = bson.D{}

	if user.Bio != "" {
		updatedData["bio"] = user.Bio
		updates = append(updates, bson.E{Key: "bio", Value: user.Bio})
	}

	if user.PhoneNumber != "" {
		updatedData["phonenumber"] = user.PhoneNumber
		updates = append(updates, bson.E{Key: "phonenumber", Value: user.PhoneNumber})
	}

	res := r.collection.FindOneAndUpdate(c, bson.D{{Key: "username", Value: username}}, bson.D{{Key: "$set", Value: updates}})
	if res.Err() != nil && res.Err() == mongo.ErrNoDocuments {
		return updatedData, domain.NewError("User not found", domain.ERR_NOT_FOUND)
	}

	if res.Err() != nil {
		return updatedData, domain.NewError(res.Err().Error(), domain.ERR_INTERNAL_SERVER)
	}

	return updatedData, nil
}
