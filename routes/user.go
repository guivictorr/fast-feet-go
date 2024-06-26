package routes

import (
	"github.com/gin-gonic/gin"
	userControllers "github.com/guivictorr/fast-feet-go/controllers/user-controllers"
	userHandlers "github.com/guivictorr/fast-feet-go/handlers/user-handlers"
	"github.com/guivictorr/fast-feet-go/middlewares"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, route *gin.RouterGroup) {
	userRepository := userControllers.NewUserController(db)
	userService := userControllers.NewUserService(userRepository)
	userHandlers := userHandlers.NewHandler(userService)

	route.Use(middlewares.Auth())

	route.GET("/", userHandlers.ListUsersHandler)
	route.GET("/:id", userHandlers.FindUserHandler)
	route.POST("/", userHandlers.CreateUserHandler)
	route.PUT("/:id", userHandlers.UpdateUserHandler)
	route.DELETE("/:id", userHandlers.DeleteUserHandler)
}
