package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naol86/go/authGo/api/controller"
	"github.com/naol86/go/authGo/config"
	"github.com/naol86/go/authGo/internal/repository"
	"github.com/naol86/go/authGo/internal/usecase"
	"gorm.io/gorm"
)

func initAuthRoute(env *config.Env, timeout time.Duration, db *gorm.DB, r *gin.RouterGroup) {
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository, timeout)
	userController := controller.UserController{
		UserUseCase: userUseCase,
		Env:         env,
	}

	r.POST("/signin", userController.Signin)
	// r.POST("/signup", userController.Signup)
}
