package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializePackageRoutes(r *gin.RouterGroup) {
	r.GET("/packages", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "list packages",
		})
	})
	r.GET("/packages/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "get package " + ctx.Param("id"),
		})
	})
	r.POST("/packages", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "create packages",
		})
	})
	r.PUT("/packages/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "update package " + ctx.Param("id"),
		})
	})
	r.DELETE("/packages/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "delete package " + ctx.Param("id"),
		})
	})
}
