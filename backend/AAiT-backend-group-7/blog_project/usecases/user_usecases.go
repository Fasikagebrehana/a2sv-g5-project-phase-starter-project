package usecases

import (
	"blog_project/domain"
	"blog_project/infrastructure"
	"context"
	"errors"
	"os"
	"regexp"
	"sync/atomic"
	"time"
)

type UserUsecase struct {
	UserRepo domain.IUserRepository
}

func NewUserUsecase(userRepo domain.IUserRepository) domain.IUserUsecase {
	return &UserUsecase{
		UserRepo: userRepo,
	}
}

func (u *UserUsecase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	return u.UserRepo.GetAllUsers(ctx)
}

func (u *UserUsecase) GetUserByID(ctx context.Context, id int) (domain.User, error) {
	user, err := u.UserRepo.GetUserByID(ctx, id)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}

	return user, nil
}

func (u *UserUsecase) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	existingUser, _ := u.UserRepo.SearchByEmail(ctx, user.Email)
	// if err != nil {
	// 	return domain.User{}, errors.New(err.Error())
	// }
	if existingUser.ID != 0 {
		return domain.User{}, errors.New("email already in use")
	}

	existingUser, _ = u.UserRepo.SearchByUsername(ctx, user.Username)
	// if err != nil {
	// 	return domain.User{}, errors.New(err.Error())
	// }
	if existingUser.ID != 0 {
		return domain.User{}, errors.New("username already in use")
	}

	if !isValidEmail(user.Email) {
		return domain.User{}, errors.New("invalid email")

	}

	if !isValidPassword(user.Password) {
		return domain.User{}, errors.New("invalid password, must contain at least one uppercase letter, one lowercase letter, one number, one special character, and minimum length of 8")
	}

	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}

	user.Password = hashedPassword

	user.ID = generateUniqueID()

	return u.UserRepo.CreateUser(ctx, user)
}

func (u *UserUsecase) UpdateUser(ctx context.Context, id int, user domain.User) (domain.User, error) {
	return u.UserRepo.UpdateUser(ctx, id, user)
}

func (u *UserUsecase) DeleteUser(ctx context.Context, id int) error {
	return u.UserRepo.DeleteUser(ctx, id)
}

func (u *UserUsecase) AddBlog(ctx context.Context, userID int, blog domain.Blog) (domain.User, error) {

	return u.UserRepo.AddBlog(ctx, userID, blog)
}

func (u *UserUsecase) DeleteBlog(ctx context.Context, userID int, blogID int) (domain.User, error) {
	user, err := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		return domain.User{}, errors.New(err.Error())
	}

	for i, blog := range user.Blogs {
		if blog == blogID {
			user.Blogs = append(user.Blogs[:i], user.Blogs[i+1:]...)
			break
		}
	}

	return u.UserRepo.UpdateUser(ctx, userID, user)
}

func (u *UserUsecase) Login(ctx context.Context, username, password string) (string, string, error) {
	user, err := u.UserRepo.SearchByUsername(ctx, username)
	if err != nil || user.ID == 0 {
		return "", "", errors.New("invalid credentials")
	}

	err = infrastructure.ComparePassword(user.Password, password)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	token, err := infrastructure.GenerateJWTAccessToken(&user, os.Getenv("jwt_secret"), 1)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := infrastructure.GenerateJWTRefreshToken(&user, os.Getenv("jwt_secret"), 5)

	if err != nil {
		return "", "", err
	}

	err = u.UserRepo.StoreRefreshToken(ctx, user.ID, refreshToken)
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, nil
}

func (u *UserUsecase) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	validatedToken, err := infrastructure.IsAuthorized(refreshToken, os.Getenv("jwt_secret"))
	if err != nil {
		return "", errors.New("invalid token")
	}

	// Convert the id from float64 to int
	userID, ok := validatedToken["id"].(float64)
	if !ok {
		return "", errors.New("invalid token ID type")
	}

	// Convert float64 to int
	user, err := u.UserRepo.GetUserByID(ctx, int(userID))
	if err != nil {
		return "", errors.New("user not found")
	}

	newToken, _ := infrastructure.GenerateJWTAccessToken(&user, os.Getenv("jwt_secret"), 1)

	return newToken, nil
}

func (u *UserUsecase) ForgetPassword(ctx context.Context, email string) error {
	user, err := u.UserRepo.SearchByEmail(ctx, email)
	if err != nil || user.ID == 0 {
		return errors.New("user not found")
	}

	// Assume infrastructure is implemented to send password reset emails
	// return infrastructure.SendResetLink(user.Email)
	return nil
}

func (u *UserUsecase) ResetPassword(ctx context.Context, username, password string) error {
	user, err := u.UserRepo.SearchByUsername(ctx, username)
	if err != nil || user.ID == 0 {
		return errors.New("user not found")
	}

	// Assume infrastructure is implemented to hash passwords
	// hashedPassword, err := infrastructure.HashPassword(password)
	// if err != nil {
	// 	return err
	// }

	// user.Password = hashedPassword
	_, err = u.UserRepo.UpdateUser(ctx, user.ID, user)
	return err
}

func (u *UserUsecase) PromoteUser(ctx context.Context, userID int) (domain.User, error) {
	user, err := u.UserRepo.GetUserByID(ctx, userID)

	if err != nil {
		return domain.User{}, nil
	}

	user.Role = "admin"

	u.UpdateUser(ctx, user.ID, user)

	return user, nil
}

func (u *UserUsecase) DemoteUser(ctx context.Context, userID int) (domain.User, error) {
	user, err := u.UserRepo.GetUserByID(ctx, userID)

	if err != nil {
		return domain.User{}, nil
	}

	user.Role = "user"

	u.UpdateUser(ctx, user.ID, user)

	return user, nil
}

// Email validation function
func isValidEmail(email string) bool {
	// Regex pattern for valid email format
	const emailRegex = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// Password strength validation function
func isValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`\d`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[\W_]`).MatchString(password)

	return hasUpper && hasLower && hasNumber && hasSpecial
}

var counter int32

func generateUniqueID() int {
	// Use a larger portion of the timestamp
	timestamp := int(time.Now().UnixNano() / 1e6 % 1e6) // Last 6 digits

	// Combine with counter
	uniqueID := timestamp*1000 + int(atomic.AddInt32(&counter, 1)%1000)

	// Ensure uniqueID fits within a 32-bit integer
	if uniqueID > 2147483647 { // Max int32 value
		uniqueID = uniqueID % 1000000
	}

	return uniqueID
}
