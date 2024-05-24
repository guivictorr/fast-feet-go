package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitPackagesRoutes(db *gorm.DB, route *gin.RouterGroup) {
	route.GET("/packages", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "list packages",
		})
	})
	route.GET("/packages/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "get package " + ctx.Param("id"),
		})
	})
	route.POST("/packages", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "create packages",
		})
	})
	route.PUT("/packages/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "update package " + ctx.Param("id"),
		})
	})
	route.DELETE("/packages/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "delete package " + ctx.Param("id"),
		})
	})
}
