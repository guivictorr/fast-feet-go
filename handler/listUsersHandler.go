package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guivictorr/fast-feet-go/schemas"
)

func ListUsersHandler(ctx *gin.Context) {
	users := []schemas.User{}
	if err := db.Find(&users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error trying to list users")
		return
	}

	sendSuccess(ctx, users, "list users")
}
