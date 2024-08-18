package usecases

import (
	"errors"
	"meleket/domain"
	"meleket/infrastructure"
	"meleket/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	userRepo    domain.UserRepositoryInterface
	// jwtSvc      infrastructure.JWTService
	// passwordSvc infrastructure.PasswordService
	// emailSvc	infrastructure.EmailService
}

// func NewUserUsecase(ur domain.UserRepositoryInterface,js infrastructure.JWTService) *UserUsecase {  //ps infrastructure.PasswordService, js infrastructure.JWTService)
// 	return &UserUsecase{
// 		userRepo:    ur,
// 		jwtSvc:      js,
// 		// passwordSvc: ps,
// 		// emailSvc: 	 es,
// 	}
// }

func NewUserUsecase(ur domain.UserRepositoryInterface) *UserUsecase {
	return &UserUsecase{userRepo: ur}
}

// Register registers a new user
func (u *UserUsecase) Register(user *domain.User) error {
	// Hash the user's password before storing it
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Save the user to the repository
	err = u.userRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}


func (u *UserUsecase) GetUserByUsername(username *string) (*domain.User, error) {
	return u.userRepo.GetUserByUsername(username)
}

func (u *UserUsecase) GetUserByEmail(email *string) (*domain.User, error) {
	return u.userRepo.GetUserByEmail(email)
}

// Login authenticates a user and returns JWT and refresh tokens if successful
func (u *UserUsecase) Login(authUser *domain.AuthUser) (string, string, error) {
	return "", "", nil
}

// DeleteRefreshToken deletes the refresh token for a user
func (u *UserUsecase) DeleteRefreshToken(userID primitive.ObjectID) error {
	// Implement the deletion of the refresh token if needed
	return nil
}

// ForgotPassword handles the forgot password logic
func (u *UserUsecase) ForgotPassword(email *string) error {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return errors.New("email not found")
	}

	otp := utils.GenerateOTP(6)
	err = infrastructure.SendOTPEmail(user.Email, otp)
	if err != nil {
		return err
	}

	return nil
}

// GetProfile retrieves a user's profile by ID
func (u *UserUsecase) GetProfile(objectID primitive.ObjectID) (*domain.Profile, error) {
	return nil, nil
}

// UpdateProfile updates a user's profile
func (u *UserUsecase) UpdateProfile(objectID primitive.ObjectID, profile *domain.Profile) (*domain.Profile, error) {
	updatedProfile, err := u.userRepo.UpdateProfile(objectID, profile)
	return updatedProfile, err
}

// GetAllUsers retrieves all users from the repository
func (u *UserUsecase) GetAllUsers() ([]*domain.User, error) {
	users, err := u.userRepo.GetAllUsers()
	return users, err
}

// DeleteUser deletes a user by ID
func (u *UserUsecase) DeleteUser(objectID primitive.ObjectID) error {
	return u.userRepo.DeleteUser(objectID)
}



// func (u *UserUsecase)GetUserByUsername(username *string) (*domain.User, error){
// 	return u.userRepo.GetUserByUsername(username)
// }

// func (u *UserUsecase) GetUserByEmail(email *string) (*domain.User, error){
// 	return u.userRepo.GetUserByEmail(email)
// }

// // Login authenticates a user and returns JWT and refresh tokens if successful
// func (u *UserUsecase) Login(authUser *domain.AuthUser) (string, string, error) {
// 	// // Retrieve the user by username
// 	// user, err := u.userRepo.GetUserByUsername(&authUser.Username)
// 	// if err != nil {
// 	// 	return "", "", errors.New("invalid username or password")
// 	// }

// 	// // Compare the provided password with the stored hashed password
// 	// if err := infrastructure.CheckPasswordHash(user.Password, authUser.Password); err != nil {
// 	// 	return "", "", errors.New("invalid username or password")
// 	// }

// 	// // Generate JWT and refresh tokens for the authenticated user
// 	// // token, err := u.jwtSvc.GenerateToken(user.ID, user.Role)
// 	// token, err := infrastructure.JWTService.GenerateToken(user.ID, user.Role)
// 	// if err != nil {
// 	// 	return "", "", err
// 	// }

// 	// refreshToken, err := infrastructure.JWTService.GenerateRefreshToken(user.ID, user.Role)
// 	// if err != nil {
// 	// 	return "", "", err
// 	// }

// 	// // Save the refresh token in the repository
// 	// err = u.userRepo.SaveToken(authUser.Username, refreshToken)
// 	// if err != nil {
// 	// 	return "", "", err
// 	// }

// 	// return token, refreshToken, nil
// 	return "", "", nil
// }

// // DeleteRefreshToken deletes the refresh token for a user
// func (u *UserUsecase) DeleteRefreshToken(userID primitive.ObjectID) error {
// 	// user, err := u.userRepo.GetUserByID(userID)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// return u.userRepo.DeleteToken(user.ID)
// 	return nil
// }

// // ForgotPassword handles the forgot password logic
// func (u *UserUsecase) ForgotPassword(email *string) error {
// 	user, err := u.userRepo.GetUserByEmail(email)
// 	if err != nil {
// 		return errors.New("email not found")
// 	}

// 	// Generate OTP and store it in the database
// 	otp := utils.GenerateOTP(6)
// 	// err = u.userRepo.StoreOTP(user.ID, otp)
// 	if err != nil {
// 		return err
// 	}

// 	// Send OTP via email
// 	// err = u.emailSvc.SendOTPEmail(user.Email, otp)
// 	err = infrastructure.SendOTPEmail(user.Email, otp)
// 	if err != nil {
// 		return err
// 	}


// 	return nil
// }


// // GetProfile retrieves a user's profile by ID
// func (u *UserUsecase) GetProfile(objectID primitive.ObjectID) (*domain.Profile, error) {
// 	// userProfile, err := u.userRepo.GetUserByID(objectID)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }
// 	// return userProfile, nil
// 	return nil, nil
// }

// // UpdateProfile updates a user's profile
// func (u *UserUsecase) UpdateProfile(objectID primitive.ObjectID, profile *domain.Profile) (*domain.Profile, error) {
// 	// user, err := u.userRepo.FindByID(objectID)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// user.Name = profile.Bio
// 	// user.avatar_url = profile.AvatarURL

// 	updatedprofile, err := u.userRepo.UpdateProfile(objectID, profile)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return updatedprofile, nil
// }

// // GetAllUsers retrieves all users from the repository
// func (u *UserUsecase) GetAllUsers() ([]*domain.User, error) {
// 	users, err := u.userRepo.GetAllUsers()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }

// // DeleteUser deletes a user by ID
// func (u *UserUsecase) DeleteUser(objectID primitive.ObjectID) error {
// 	return u.userRepo.DeleteUser(objectID)
// }