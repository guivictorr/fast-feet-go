package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendSuccess(ctx *gin.Context, data interface{}, op string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

func sendError(ctx *gin.Context, status int, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(status, gin.H{
		"message":   message,
		"errorCode": status,
	})
}
