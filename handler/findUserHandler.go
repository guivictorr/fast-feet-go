package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guivictorr/fast-feet-go/schemas"
)

func FindUserHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	users := schemas.User{}
	if err := db.Where("id = ?", userId).First(&users).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "this user doesn't exist")
		return
	}

	sendSuccess(ctx, users, "find user")
}
