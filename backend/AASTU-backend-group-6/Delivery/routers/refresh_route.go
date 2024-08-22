package routers

import (
	infrastructure "blogs/Infrastructure"
	"blogs/mongo"
	"time"

	"blogs/Delivery/controllers"
	repositories "blogs/Repositories"
	usecase "blogs/Usecases"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(env *infrastructure.Config, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	aur := repositories.NewActiveUserRepository(db, env.ActiveUserCollection)
	ur := repositories.NewUserRepository(db, env.UserCollection)
	rtc := &controllers.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur,aur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
