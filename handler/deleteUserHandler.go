package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guivictorr/fast-feet-go/schemas"
)

func DeleteUserHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	user := schemas.User{}
	if err := db.Where("id = ?", userId).First(&user).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "this user doesn't exist")
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error trying to delete user")
		return
	}

	sendSuccess(ctx, nil, "delete user")
}
