package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guivictorr/fast-feet-go/handler"
	"gorm.io/gorm"
)

func InitUsersRoutes(db *gorm.DB, route *gin.RouterGroup) {
	handler.InitializeHandler()
	route.GET("/", handler.ListUsersHandler)
	route.GET("/:id", handler.FindUserHandler)
	route.POST("/", handler.CreateUserHandler)
	route.PUT("/:id", handler.UpdateUserHandler)
	route.DELETE("/:id", handler.DeleteUserHandler)
}
