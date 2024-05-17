package userHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guivictorr/fast-feet-go/utils"
)

func (h *handler) DeleteUserHandler(ctx *gin.Context) {
	userId := ctx.Param("id")

	statusCode := h.service.DeleteUser(userId)

	switch statusCode {
	case http.StatusNoContent:
		utils.APIResponse(ctx, "User deleted successfully", statusCode, nil)
		return
	case http.StatusNotFound:
		utils.APIResponse(ctx, "User not found", statusCode, nil)
		return
	case http.StatusInternalServerError:
		utils.APIResponse(ctx, "Error deleting user", statusCode, nil)
		return
	default:
		utils.APIResponse(ctx, "Something went wrong", http.StatusBadRequest, nil)
	}
}
