package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guivictorr/fast-feet-go/handler"
)

func initializeUsersRoutes(r *gin.RouterGroup) {
	r.GET("/users", handler.ListUsersHandler)
	r.GET("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "get user " + ctx.Param("id"),
		})
	})
	r.POST("/users", handler.CreateUserHandler)
	r.PUT("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "update user " + ctx.Param("id"),
		})
	})
	r.DELETE("/users/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "delete user " + ctx.Param("id"),
		})
	})
}
