package routes

import (
	"github.com/gin-gonic/gin"
	authControllers "github.com/guivictorr/fast-feet-go/controllers/auth-controllers"
	authHandlers "github.com/guivictorr/fast-feet-go/handlers/auth-handlers"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, route *gin.RouterGroup) {
	authRepository := authControllers.NewAuthRepository(db)
	authService := authControllers.NewAuthService(authRepository)
	authHandlers := authHandlers.NewAuthHandler(authService)

	route.POST("/session", authHandlers.CreateSessionHandler)
}
