package userHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guivictorr/fast-feet-go/utils"
)

func (h *handler) FindUserHandler(ctx *gin.Context) {
	userId := ctx.Param("id")

	userResult, statusCode := h.service.FindUser(userId)

	switch statusCode {
	case http.StatusOK:
		utils.APIResponse(ctx, "User found successfully", statusCode, http.MethodGet, userResult)
		return
	case http.StatusNotFound:
		utils.APIResponse(ctx, "User not found", statusCode, http.MethodGet, nil)
		return
	default:
		utils.APIResponse(ctx, "Something went wrong", http.StatusBadRequest, http.MethodGet, nil)
	}
}
