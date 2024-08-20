package usecase

import (
	infrastructure "astu-backend-g1/Infrastructure"
	"astu-backend-g1/domain"
	"crypto/rand"
	"errors"
	"math/big"
	"time"
)

type userUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(u domain.UserRepository) (domain.UserUsecase, error) {
	return &userUsecase{userRepository: u}, nil
}

func (useCase *userUsecase) Get() ([]domain.User, error) {
	return useCase.userRepository.Get(domain.UserFilterOption{})
}
func (useCase *userUsecase)LoginUser(uname string,password string) (string, error) {
	user  ,err:= useCase.GetByUsername(uname)
	if err != nil {
		return "",err
	}
	if user.IsActive == false {
		return "", errors.New("Account not activated")
	}
	accesstoken, refreshToken, err := infrastructure.GenerateToken(&user, password)
	if err != nil {
		return "",  err
	}
	user.RefreshToken = refreshToken
	useCase.userRepository.Update(user.ID, user)

	return accesstoken, nil
}
// function for forget password
func (useCase *userUsecase) ForgetPassword(email string) (string, error) {
	user, err := useCase.GetByEmail(email)
	if err != nil {
		return "", err
	}
	if user.IsActive == false {
		return "", errors.New("Account not activated")
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	confirmationToken := make([]byte, 64)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < 64; i++ {
		num, _ := rand.Int(rand.Reader, charsetLength)
		confirmationToken[i] = charset[num.Int64()]
	}

	//adding the token to the user
	user.VerifyToken = string(confirmationToken)
	
	// Add 20 minutes to the current time
	expirationTime := time.Now().Add(20 * time.Minute)
	user.ExpirationDate = expirationTime
	useCase.userRepository.Update(user.ID, user)
	link := "http://localhost:8000/resetpassword/?email=" + user.Email + "&token=" + string(confirmationToken) + "/"
	err = infrastructure.SendEmail(user.Email, "Password Reset", "This is the password reset link: ",link)
	if err != nil {
		return "", err
	}


	return "Password reset token sent to your email", nil
}

// handle password reset
func (useCase *userUsecase) ResetPassword(email string, token string, password string) (string, error) {
	user, err := useCase.GetByEmail(email)
	if err != nil {
		return "", err
	}
		if user.IsActive == false {
		return "", errors.New("Account not activated")
	}
	currentTime := time.Now()
	if user.VerifyToken == token {
		// Check if token has expired
		if  currentTime.After(user.ExpirationDate){
			return "Token has expired", nil
		}
		user.Password, _ = infrastructure.PasswordHasher(password)
		useCase.userRepository.Update(user.ID, user)
		return "Password reset successful", nil
	}
	return "Invalid token", nil
}




func (useCase *userUsecase) GetByID(userID string) (domain.User, error) {
	filter := domain.UserFilter{UserId: userID}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	return users[0], err
}

func (useCase *userUsecase) GetByUsername(username string) (domain.User, error) {
	filter := domain.UserFilter{Username: username}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	return users[0], err
}
func (useCase *userUsecase) AccountVerification(uemail string, confirmationToken string) (string, error) {
	filter := domain.UserFilter{Email: uemail}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	if users[0].VerifyToken == confirmationToken {
		accesstoken,refreshToken, err := infrastructure.GenerateToken(&users[0], users[0].Password)
		if err != nil {
			return "", err
		}
		users[0].IsActive = true
		users[0].RefreshToken = refreshToken
		useCase.userRepository.Update(users[0].ID, users[0])
		
		return accesstoken, nil

	}
	return "", err
}

func (useCase *userUsecase) GetByEmail(email string) (domain.User, error) {
	filter := domain.UserFilter{Email: email}
	opts := domain.UserFilterOption{Filter: filter}
	users, err := useCase.userRepository.Get(opts)
	return users[0], err
}

func (useCase *userUsecase) Create(u *domain.User) (domain.User, error) {
	u.Password, _ = infrastructure.PasswordHasher(u.Password)
	u.IsActive = false
	
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	confirmationToken := make([]byte, 64)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < 64; i++ {
		num, _ := rand.Int(rand.Reader, charsetLength)
		confirmationToken[i] = charset[num.Int64()]
	}
	u.VerifyToken = string(confirmationToken)
	nUser, err := useCase.userRepository.Create(u)
	link := "http://localhost:8000/accountVerification/?email=" + u.Email + "&token=" + string(confirmationToken) + "/"
	err = infrastructure.SendEmail(u.Email, "Registration Confirmation", "This sign up Confirmation email to verify: ", link)
	
	if err != nil {
		return nUser, err
	}
	return nUser, err
}

func (useCase *userUsecase) Update(userId string, updateData domain.User) (domain.User, error) {
	return useCase.userRepository.Update(userId, updateData)
}

func (useCase *userUsecase) Delete(userId string) error {
	return useCase.userRepository.Delete(userId)
}
