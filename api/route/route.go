package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naol86/go/authGo/config"
	"gorm.io/gorm"
)

func SetUpRoutes(env *config.Env, timeout time.Duration, db *gorm.DB, r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// public routes
	authRoute := r.Group("/auth")
	initAuthRoute(env, timeout, db, authRoute)

	// protected routes
}
