package repository

import (
	"blog/database"
	"blog/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userRepository struct {
	database   database.Database
	collection string
}

// CreateUser implements domain.UserRepository.
func (u *userRepository) CreateUser(c context.Context, user *domain.User) error {
	collection := u.database.Collection(u.collection)
	_, err := collection.InsertOne(c, user)
	return err
}

// DeleteUser implements domain.UserRepository.
func (u *userRepository) DeleteUser(c context.Context, id primitive.ObjectID) error {
	collection := u.database.Collection(u.collection)
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(c, filter)
	return err
}

// GetUserByEmail implements domain.UserRepository.
func (u *userRepository) GetUserByEmail(c context.Context, email string) (*domain.User, error) {
	collection := u.database.Collection(u.collection)
	filter := bson.M{"email": email}
	user := &domain.User{}
	err := collection.FindOne(c, filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

// GetUserByID implements domain.UserRepository.
func (u *userRepository) GetUserByID(c context.Context, id primitive.ObjectID) (*domain.User, error) {
	collection := u.database.Collection(u.collection)
	filter := bson.M{"_id": id}
	user := &domain.User{}
	err := collection.FindOne(c, filter).Decode(user)
	return user, err
}

// GetUserByUsername implements domain.UserRepository.
func (u *userRepository) GetUserByUsername(c context.Context, username string) (*domain.User, error) {
	collection := u.database.Collection(u.collection)
	filter := bson.M{"username": username}
	user := &domain.User{}
	err := collection.FindOne(c, filter).Decode(user)
	if err != nil {
		return nil, err
	}
	return user, err
}

// GetUsers implements domain.UserRepository.
func (u *userRepository) GetAllUsers(c context.Context) ([]*domain.User, error) {
	collection := u.database.Collection(u.collection)
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(c)
	
	var users []*domain.User
	for cursor.Next(c) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// UpdateUser implements domain.UserRepository.
func (u *userRepository) UpdateUser(c context.Context, user *domain.User) error {
	collection := u.database.Collection(u.collection)
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": user}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}
func (u *userRepository) UpdatePassword(c context.Context, user *domain.User) error {
	collection := u.database.Collection(u.collection)
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{"password": user.Password}}
	_, err := collection.UpdateOne(c, filter, update)
	return err
}
func (u *userRepository) PromoteUser(c context.Context, id primitive.ObjectID)  error {
	collection := u.database.Collection(u.collection)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"role": "admin"}}
	_, err := collection.UpdateOne(c, filter, update)
	return  err
}
func (u *userRepository) DemoteUser(c context.Context, id primitive.ObjectID)  error {
	collection := u.database.Collection(u.collection)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"role": "user"}}
	_, err := collection.UpdateOne(c, filter, update)
	return  err
}

func NewUserRepository(db database.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}
