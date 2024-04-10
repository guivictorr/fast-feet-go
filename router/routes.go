package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/packages", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "list-packages",
			})
		})
	}
}
