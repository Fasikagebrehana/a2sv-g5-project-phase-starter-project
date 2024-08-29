package repository

import (
	"context"
	"errors"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	custom_error "github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/errors"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userProjection = bson.M{
	"password": 0,
}

type userRepository struct {
	database       mongo.Database
	collectionName string
}

func NewUserRepository(db mongo.Database, collection string) entities.UserRepository {
	return &userRepository{
		database:       db,
		collectionName: collection,
	}
}

func (ur *userRepository) CreateUser(c context.Context, user *entities.User) (*entities.User, error) {

	collection := ur.database.Collection(ur.collectionName)

	res, err := collection.InsertOne(c, user)

	if err != nil {
		return nil, err
	}
	// Find the inserted user by ID
	insertedID, _ := res.InsertedID.(primitive.ObjectID)
	var insertedUser entities.User
	err = collection.FindOne(c, bson.M{"_id": insertedID}).Decode(&insertedUser)

	insertedUser.Password = ""

	if err != nil {
		return nil, custom_error.ErrErrorCreatingUser
	}

	return &insertedUser, err
}
func (ur *userRepository) IsOwner(c context.Context) (bool, error) {
	collection := ur.database.Collection(ur.collectionName)
	count, err := collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return false, custom_error.ErrErrorCreatingUser
	}
	return count == 0, nil
}
func (ur *userRepository) GetAllUsers(c context.Context) ([]entities.User, error) {
	collection := ur.database.Collection(ur.collectionName)
	var users []entities.User

	opts := options.Find().SetProjection(userProjection)

	cursor, err := collection.Find(c, bson.M{}, opts)
	if err != nil {
		return []entities.User{}, err
	}
	err = cursor.All(c, &users)
	if err != nil {
		return []entities.User{}, err
	}

	return users, nil

}
func (ur *userRepository) UpdateRefreshToken(c context.Context, userID string, refreshToken string) error {
	collection := ur.database.Collection(ur.collectionName)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": id}
	_, err = collection.UpdateOne(c, filter, bson.M{"$push": bson.M{"tokens": refreshToken}})
	return err
}

func (ur *userRepository) UpdateLastLogin(c context.Context, userID string) error {
	collection := ur.database.Collection(ur.collectionName)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": id}
	_, err = collection.UpdateOne(c, filter, bson.M{"$set": bson.M{"last_login": primitive.NewDateTimeFromTime(time.Now())}})

	if err != nil {
		return custom_error.ErrErrorUpdatingUser
	}

	return err
}

func (ur *userRepository) GetUserById(c context.Context, userId string) (*entities.User, error) {
	collection := ur.database.Collection(ur.collectionName)
	var user entities.User

	id, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return nil, custom_error.ErrInvalidID
	}

	opts := options.FindOne().SetProjection(userProjection)

	err = collection.FindOne(c, bson.M{"_id": id}, opts).Decode(&user)

	if err != nil {
		return nil, custom_error.ErrUserNotFound
	}

	return &user, err

}

func (ur *userRepository) GetUserByEmail(c context.Context, email string) (*entities.User, error) {
	collection := ur.database.Collection(ur.collectionName)
	var user entities.User

	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, custom_error.ErrUserNotFound
	}
	return &user, err
}

func (ur *userRepository) GetUsers(c context.Context, filter bson.M, userFilter entities.UserFilter) (*[]entities.User, mongopagination.PaginationData, error) {
	collection := ur.database.Collection(ur.collectionName)

	projectQuery := bson.M{"$project": userProjection}

	var aggUserList []entities.User = make([]entities.User, 0)

	paginatedData, err := mongopagination.New(collection).Context(c).Limit(userFilter.Limit).Page(userFilter.Pages).Aggregate(filter, projectQuery)

	if err != nil {
		return &[]entities.User{}, mongopagination.PaginationData{}, custom_error.ErrFilteringUsers
	}

	for _, raw := range paginatedData.Data {
		var user *entities.User
		if marshallErr := bson.Unmarshal(raw, &user); marshallErr == nil {
			aggUserList = append(aggUserList, *user)
		}

	}

	return &aggUserList, paginatedData.Pagination, nil

}

func (ur *userRepository) RevokeRefreshToken(c context.Context, userID, refreshToken string) error {
	collection := ur.database.Collection(ur.collectionName)
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("object id invalid")
	}

	res, err := collection.UpdateOne(c, bson.M{"_id": objID}, bson.M{"$pull": bson.M{"tokens": refreshToken}})

	if err != nil {
		return err
	}
	if res.MatchedCount < 1 {
		return custom_error.ErrTokenNotFound
	}
	return nil
}

func (ur *userRepository) UpdateUser(c context.Context, userID string, updatedUser *entities.UserUpdate) (*entities.User, error) {
	collection := ur.database.Collection(ur.collectionName)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, custom_error.ErrInvalidID
	}

	filter := bson.M{"_id": id}

	if updatedUser.FirstName == "" && updatedUser.LastName == "" && updatedUser.Bio == "" {
		return nil, custom_error.ErrNoUpdateFields
	}

	var update bson.M

	if updatedUser.FirstName != "" {
		update = bson.M{"$set": bson.M{"first_name": updatedUser.FirstName}}
	}

	if updatedUser.LastName != "" {
		update = bson.M{"$set": bson.M{"last_name": updatedUser.LastName}}
	}

	if updatedUser.Bio != "" {
		update = bson.M{"$set": bson.M{"bio": updatedUser.Bio}}
	}

	if updatedUser.FirstName != "" && updatedUser.LastName != "" {
		update = bson.M{"$set": bson.M{
			"first_name": updatedUser.FirstName,
			"last_name":  updatedUser.LastName,
		}}
	}

	if updatedUser.FirstName != "" && updatedUser.Bio != "" {
		update = bson.M{"$set": bson.M{
			"first_name": updatedUser.FirstName,
			"bio":        updatedUser.Bio,
		}}

	}

	if updatedUser.LastName != "" && updatedUser.Bio != "" {
		update = bson.M{"$set": bson.M{
			"last_name": updatedUser.LastName,
			"bio":       updatedUser.Bio,
		}}
	}

	if updatedUser.FirstName != "" && updatedUser.LastName != "" && updatedUser.Bio != "" {
		update = bson.M{"$set": bson.M{
			"first_name": updatedUser.FirstName,
			"last_name":  updatedUser.LastName,
			"bio":        updatedUser.Bio,
		}}
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetProjection(userProjection)

	var ResultUser entities.User
	err = collection.FindOneAndUpdate(c, filter, update, opts).Decode(&ResultUser)

	if err != nil {
		return nil, custom_error.ErrErrorUpdatingUser
	}

	return &ResultUser, nil
}

func (ur *userRepository) ActivateUser(c context.Context, userID string) (*entities.User, error) {
	collection := ur.database.Collection(ur.collectionName)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"is_active": true}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetProjection(userProjection)

	var ResultUser entities.User
	err = collection.FindOneAndUpdate(c, filter, update, opts).Decode(&ResultUser)

	if err != nil {
		return nil, custom_error.ErrErrorUpdatingUser
	}

	return &ResultUser, nil
}

func (ur *userRepository) DeleteUser(c context.Context, userID string) error {
	collection := ur.database.Collection(ur.collectionName)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	filter := bson.M{"_id": id}
	res, err := collection.DeleteOne(c, filter)

	if res.DeletedCount == 0 {
		return custom_error.ErrUserNotFound
	}

	if err != nil {
		return custom_error.ErrErrorUpdatingUser
	}

	return nil
}
func (ur *userRepository) IsUserActive(c context.Context, userID string) (bool, error) {
	collection := ur.database.Collection(ur.collectionName)
	var user entities.User
	err := collection.FindOne(c, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return false, custom_error.ErrUserNotFound
	}
	return user.Active, err

}
func (ur *userRepository) ResetUserPassword(c context.Context, userID string, resetPassword *entities.ResetPasswordRequest) error {
	collection := ur.database.Collection(ur.collectionName)
	ObjID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return custom_error.ErrInvalidID
	}
	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"password": resetPassword.NewPassword}})
	if res.MatchedCount < 1 {
		return custom_error.ErrUserNotFound
	}
	if err != nil {
		return custom_error.ErrErrorUpdatingPassword
	}
	return nil
}
func (ur *userRepository) UpdateUserPassword(c context.Context, userID string, updatePassword *entities.UpdatePassword) error {
	collection := ur.database.Collection(ur.collectionName)
	ObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}

	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"password": updatePassword.NewPassword}})
	if res.MatchedCount < 1 {
		return custom_error.ErrUserNotFound
	}
	if err != nil {
		return custom_error.ErrErrorUpdatingPassword
	}
	return nil
}
func (ur *userRepository) PromoteUserToAdmin(c context.Context, userID string) error {
	collection := ur.database.Collection(ur.collectionName)
	ObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"role": "admin"}})
	if res.MatchedCount < 1 {
		return custom_error.ErrUserNotFound
	}
	if err != nil {
		return custom_error.ErrErrorUpdatingUser
	}
	return nil
}
func (ur *userRepository) DemoteAdminToUser(c context.Context, userID string) error {
	collection := ur.database.Collection(ur.collectionName)
	ObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"role": "user"}})
	if res.MatchedCount < 1 {
		return custom_error.ErrUserNotFound
	}
	if err != nil {
		return custom_error.ErrErrorUpdatingUser
	}
	return nil
}
func (ur *userRepository) UpdateProfilePicture(c context.Context, userID string, filename string) error {
	collection := ur.database.Collection(ur.collectionName)
	ObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return custom_error.ErrInvalidID
	}
	res, err := collection.UpdateOne(c, bson.M{"_id": ObjID}, bson.M{"$set": bson.M{"profile_img": filename}})
	if res.ModifiedCount < 1 {
		return custom_error.ErrUserNotFound
	}
	return nil
}

func (ur *userRepository) GetInactiveUsersForReactivation(c context.Context, emailTreshold primitive.DateTime, deleteTreshold primitive.DateTime) (*[]entities.User, error) {
	collection := ur.database.Collection(ur.collectionName)

	filter := bson.M{
		"is_active":  false,
		"created_at": bson.M{"$lt": emailTreshold, "$gte": deleteTreshold},
	}

	// Users who haven't activated their account within `emailTreshold` days
	var users []entities.User
	cur, err := collection.Find(c, filter)
	if err != nil {
		return &[]entities.User{}, custom_error.ErrErrorSendingReminderEmail
	}
	err = cur.All(c, &users)
	if err != nil {
		return &[]entities.User{}, custom_error.ErrErrorSendingReminderEmail
	}

	return &users, nil
}

func (ur *userRepository) DeleteInActiveUser(c context.Context, deleteTreshold primitive.DateTime) error {
	collection := ur.database.Collection(ur.collectionName)

	filter := bson.M{
		"is_active":  false,
		"created_at": bson.M{"$lt": deleteTreshold},
	}

	_, err := collection.DeleteMany(c, filter)
	if err != nil {
		return err
	}

	return nil
}

// GetRefreshToken implements entities.UserRepository.
func (ur *userRepository) RefreshTokenExist(c context.Context, userID, refreshToken string) (bool, error) {
	collection := ur.database.Collection(ur.collectionName)
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return false, err
	}
	filter := bson.M{
		"_id":    id,
		"tokens": refreshToken, // Check if the refreshToken exists tokens[]
	}
	err = collection.FindOne(c, filter).Decode(&entities.User{})
	if err != nil {
		return false, nil
	}

	return true, nil

}
