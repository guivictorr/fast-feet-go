package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello-world",
		})
	})
	r.Run()
}
