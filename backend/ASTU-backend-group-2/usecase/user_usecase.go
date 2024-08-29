package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain/entities"
	mongopagination "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUsecase struct {
	userRepository entities.UserRepository
	contextTimeout time.Duration
	Env            *bootstrap.Env
}

func NewUserUsecase(userRepository entities.UserRepository, timeout time.Duration) entities.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (uu *userUsecase) CreateUser(c context.Context, user *entities.User) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.CreateUser(ctx, user)
}

func (uu *userUsecase) GetUserByEmail(c context.Context, email string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.GetUserByEmail(ctx, email)
}

func (uu *userUsecase) GetUserById(c context.Context, userId string) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.GetUserById(ctx, userId)
}
func (uu *userUsecase) GetUsers(c context.Context, userFilter entities.UserFilter) (*[]entities.User, mongopagination.PaginationData, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	filter := UserFilterOption(userFilter)
	users, meta, err := uu.userRepository.GetUsers(ctx, filter, userFilter)

	if err != nil {
		return nil, mongopagination.PaginationData{}, err
	}

	// map users to userout
	res := make([]entities.User, 0)

	for _, user := range *users {
		res = append(res, user)
	}

	return &res, meta, nil
}
func (uu *userUsecase) GetAllUsers(c context.Context) ([]entities.User, error) {
	return uu.userRepository.GetAllUsers(c)
}
func (uu *userUsecase) UpdateUser(c context.Context, userID string, updatedUser *entities.UserUpdate) (*entities.User, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	updatedUser.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	return uu.userRepository.UpdateUser(ctx, userID, updatedUser)
}

func (uu *userUsecase) ActivateUser(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	_, err := uu.userRepository.ActivateUser(ctx, userID)
	return err
}

func (uu *userUsecase) DeleteUser(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.DeleteUser(ctx, userID)
}

func (uu *userUsecase) IsUserActive(c context.Context, userID string) (bool, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.IsUserActive(ctx, userID)
}

func (uu *userUsecase) IsOwner(c context.Context) (bool, error) {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.IsOwner(ctx)
}

func (uu *userUsecase) ResetUserPassword(c context.Context, userID string, resetPassword *entities.ResetPasswordRequest) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.ResetUserPassword(ctx, userID, resetPassword)
}

func (uu *userUsecase) UpdateUserPassword(c context.Context, userID string, updatePassword *entities.UpdatePassword) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.UpdateUserPassword(ctx, userID, updatePassword)
}

func (uu *userUsecase) PromoteUserToAdmin(c context.Context, userID string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	return uu.userRepository.PromoteUserToAdmin(ctx, userID)
}

func (uu *userUsecase) DemoteAdminToUser(c context.Context, userID string) error {
	user, err := uu.GetUserById(c, userID)
	if err != nil {
		return err
	}
	if user.IsOwner {
		return errors.New("cannot demote owner")
	}
	return uu.userRepository.DemoteAdminToUser(c, userID)
}
func (uu *userUsecase) UpdateProfilePicture(c context.Context, userID string, filename string) error {
	ctx, cancel := context.WithTimeout(c, uu.contextTimeout)
	defer cancel()
	_, err := uu.GetUserById(c, userID)
	if err != nil {
		return err
	}
	return uu.userRepository.UpdateProfilePicture(ctx, userID, filename)
}

func UserFilterOption(filter entities.UserFilter) bson.M {

	query := bson.M{
		"$match": bson.M{},
	}
	semiquery := query["$match"].(bson.M)

	// Email filter
	if filter.Email != "" {
		semiquery["email"] = bson.M{"$regex": filter.Email, "$options": "i"}
	}

	// filter.Role
	if filter.Role != "" {
		semiquery["role"] = filter.Role
	}

	// Active filter
	if filter.Active != "" {
		semiquery["active"] = filter.Active == "true"
	}

	// Bio filter
	if filter.Bio != "" {
		semiquery["bio"] = bson.M{"$regex": filter.Bio, "$options": "i"} // case-insensitive search
	}

	// First name filter
	if filter.FirstName != "" {
		semiquery["first_name"] = bson.M{"$regex": filter.FirstName, "$options": "i"} // case-insensitive search
	}

	// Last name filter
	if filter.LastName != "" {
		semiquery["last_name"] = bson.M{"$regex": filter.LastName, "$options": "i"} // case-insensitive search
	}

	// Is owner filter
	if filter.IsOwner != "" {
		semiquery["is_owner"] = filter.IsOwner == "true"
	}

	// Is admin filter
	if filter.Role != "" {
		semiquery["role"] = filter.Role
	}

	// Date range filter
	if !filter.DateFrom.IsZero() && !filter.DateTo.IsZero() {
		semiquery["created_at"] = bson.M{
			"$gte": filter.DateFrom,
			"$lte": filter.DateTo,
		}
	} else if !filter.DateFrom.IsZero() {
		semiquery["created_at"] = bson.M{"$gte": filter.DateFrom}
	} else if !filter.DateTo.IsZero() {
		semiquery["created_at"] = bson.M{"$lte": filter.DateTo}
	}

	log.Println(query)
	return query

}

func (uu *userUsecase) SendReminderEmail(c context.Context) error {
	// ctx, cancel := context.WithTimeout(c, time.Second*10)
	// defer cancel()

	// emailTreshold := primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -3))
	// deleteTreshold := primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -10))

	// users, err := uu.userRepository.GetInactiveUsersForReactivation(ctx, emailTreshold, deleteTreshold)

	// if err != nil {
	// 	return err
	// }

	// for _, user := range *users {
	// 	// Send activation email
	// 	err := emailutil.SendVerificationEmail(user.Email, user.VerToken, uu.Env)
	// 	if err != nil {
	// 		log.Println("Error sending reminder email to user:", user.Email)
	// 		continue
	// 	}
	// }

	return nil
}

func (uu *userUsecase) DeleteInActiveUsers(c context.Context) error {
	ctx, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	deleteTreshold := primitive.NewDateTimeFromTime(time.Now().AddDate(0, 0, -10))

	err := uu.userRepository.DeleteInActiveUser(ctx, deleteTreshold)

	return err
}

// Schedule the process to send reminders and delete inactive users
func (uu *userUsecase) ScheduleDeleteAndReminderForInActiveUser(c context.Context) {
	ctx, cancel := context.WithTimeout(c, time.Second*10)
	defer cancel()

	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			// Send reminder emails to users who haven't activated within 3 days
			err := uu.SendReminderEmail(ctx)
			if err != nil {
				log.Println("Error sending reminder emails:", err)
			}

			// Delete users who have been inactive for more than 7 days
			err = uu.DeleteInActiveUsers(c)
			if err != nil {
				log.Println("Error deleting inactive users:", err)
			}
		}
	}()
}
