package userHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guivictorr/fast-feet-go/utils"
)

func (h *handler) ListUsersHandler(ctx *gin.Context) {
	userResult, statusCode := h.service.ListUsers()

	switch statusCode {
	case http.StatusOK:
		utils.APIResponse(ctx, "Success listing users", statusCode, http.MethodPost, userResult)
		return
	default:
		utils.APIResponse(ctx, "Something went wrong", http.StatusBadRequest, http.MethodPost, nil)
	}
}
