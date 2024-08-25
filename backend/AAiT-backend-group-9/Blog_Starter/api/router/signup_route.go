package router

import (
	"time"

	"Blog_Starter/api/controller"
	"Blog_Starter/config"
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"Blog_Starter/usecase"
	"Blog_Starter/utils/infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewSignupRouter sets up the signup routes.
func NewSignupRouter(env *config.Env, timeout time.Duration, db *mongo.Client, group *gin.RouterGroup) {
	database := db.Database(env.DBName) // Replace with your actual database name
	or := repository.NewOtpRepository(database, domain.CollectionOTP)
	ur := repository.NewUserRepository(database, domain.CollectionUser)
	tm := &infrastructure.NewTokenManager{}

	sc := controller.NewSignUpController(
		usecase.NewSignUpUsecase(ur, tm, env, timeout),
		usecase.NewOtpUsecase(or, timeout),
	)
	group.POST("/signup", sc.SignUp)
	group.POST("/verifyemail", sc.VerifyEmail)
	group.POST("/resendotp", sc.ResendOTP)
	group.POST("/continuewithgoogle", sc.FederatedSignup)
}
