package router

import (
	"github.com/gin-gonic/gin"
	"github.com/guivictorr/fast-feet-go/handler"
)

func initializeUsersRoutes(r *gin.RouterGroup) {
	r.GET("/users", handler.ListUsersHandler)
	r.GET("/users/:id", handler.FindUserHandler)
	r.POST("/users", handler.CreateUserHandler)
	r.PUT("/users/:id", handler.UpdateUserHandler)
	r.DELETE("/users/:id", handler.DeleteUserHandler)
}
