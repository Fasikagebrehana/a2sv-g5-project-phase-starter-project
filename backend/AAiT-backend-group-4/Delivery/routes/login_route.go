package routes

import (
	bootstrap "aait-backend-group4/Bootstrap"
	controllers "aait-backend-group4/Delivery/Controllers"
	infrastructure "aait-backend-group4/Infrastructure"
	repositories "aait-backend-group4/Repositories"
	usecases "aait-backend-group4/Usecases"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewLoginRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(db, env.UserCollection)
	userRepository := repositories.NewUserRepository(db, env.UserCollection)
	ts := infrastructure.NewTokenService(userRepository, env)
	lc := controllers.LoginController{
		LoginUsecase: usecases.NewLoginUsecase(ur, ts, timeout, env),
		Env:          env,
	}

	loc := controllers.NewLogoutController(usecases.NewLogoutUsecase(ts, env))

	group.POST("/user/login", lc.Login)
	group.GET("/user/logout", loc.Logout)
}
