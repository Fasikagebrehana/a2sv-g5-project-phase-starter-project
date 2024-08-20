package repositories

import (
	domain "blogs/Domain"
	"blogs/mongo"
	"context"
	"time"

	// "go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

type signupRepository struct {
	database   mongo.Database
	collection string
}


func NewSignupRepository(database mongo.Database, collection string) domain.SignupRepository {
	return &signupRepository{
		database:   database,
		collection: collection}

}


func (r *signupRepository) Create(c context.Context ,user domain.User) (domain.User , error) { 
	collection := r.database.Collection(r.collection)
	_, err := collection.InsertOne(c, user)

	if err != nil {
		return domain.User{} , err
	}

	return user , nil
}

func (s *signupRepository) FindUserByEmail(ctx context.Context, email string) (domain.User, error) {
	collection := s.database.Collection(s.collection)
	var user domain.User
	err := collection.FindOne(ctx, bson.M{"email" : email}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}


// the function that sets the otp code to the db 
func (r *signupRepository) SetOTP(c context.Context , email string , otp string) (error) { 
	collection := r.database.Collection(r.collection)
	// data we put in
	
	data := bson.M{"$set": bson.M{"otp": otp}}  // Corrected the key
	_ , err := collection.UpdateOne(c, bson.M{"email": email}, data)
	if err != nil {
		return err
	}

	return nil
}


// verify OTP code

func (r *signupRepository) VerifyOTP(c context.Context , email string , otp string) error { 
	collection := r.database.Collection(r.collection)
	// data we put in getting it and checking from the current otp 
	var user domain.User
	err := collection.FindOne(c , bson.M{otp: otp}).Decode(&user)

	if err != nil { 
		return err
	}
	return nil
}

// update user
func (r *signupRepository) VerifyUser(c context.Context ,  user domain.User) (domain.User , error) { 
	collection := r.database.Collection(r.collection)
	
	update := bson.M{"$set": user}
	_, err := collection.UpdateOne(c , bson.M{"email": user.Email} , update)
	if err != nil {
		return domain.User{} , err
}
	return user , nil
	
	
}


func (r *signupRepository) SetResetToken(c context.Context, email domain.ForgotPasswordRequest, token string , expiration time.Time) (domain.User, error) {
    // Reference to the collection
    collection := r.database.Collection(r.collection)

    // Filter to find the user by email
    filter := bson.M{"email": email.Email}

    // Update to set the reset_token
    update := bson.M{"$set": bson.M{"resetpasswordtoken": token , "resetpasswordexpires": expiration}}

    // Execute the update
    _ , err := collection.UpdateOne(c, filter, update)
    if err != nil {
        return domain.User{}, err // Return an empty User and the error
    }

    var updatedUser domain.User
    err = collection.FindOne(c, filter).Decode(&updatedUser)
    if err != nil {
        return domain.User{}, err
    }

    return updatedUser, nil 

}

func (r *signupRepository) FindUserByResetToken(c context.Context, token string) (domain.User, error) {
	// Reference to the collection
	collection := r.database.Collection(r.collection)

	// Filter to find the user by reset_token
	filter := bson.M{"resetpasswordtoken": token}

	// Find the user
	var user domain.User
	err := collection.FindOne(c, filter).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *signupRepository) UpdateUser(c context.Context, user domain.User) (domain.User ,error) {
    collection := r.database.Collection(r.collection)

	update := bson.M{"$set": user}
	_, err := collection.UpdateOne(c , bson.M{"email": user.Email} , update)
	if err != nil {
		return domain.User{} , err
}
	return user , nil

}

func (r *signupRepository) FindUserByUsername(c context.Context, username string) (domain.User , error) {
	collection := r.database.Collection(r.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

