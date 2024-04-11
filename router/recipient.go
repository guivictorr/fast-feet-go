package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRecipientsRoutes(r *gin.RouterGroup) {
	r.GET("/recipients", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "list recipients",
		})
	})
	r.GET("/recipients/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "get delivery " + ctx.Param("id"),
		})
	})
	r.POST("/recipients", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "create recipients",
		})
	})
	r.PUT("/recipients/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "update recipient " + ctx.Param("id"),
		})
	})
	r.DELETE("/recipients/:id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "delete recipient " + ctx.Param("id"),
		})
	})
}
