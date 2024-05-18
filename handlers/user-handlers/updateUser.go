package userHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userControllers "github.com/guivictorr/fast-feet-go/controllers/user-controllers"
	"github.com/guivictorr/fast-feet-go/utils"
)

func (h *handler) UpdateUserHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	var input userControllers.UpdateUserInput
	ctx.ShouldBindJSON(&input)

	err := utils.GoValidator(&input)
	if err != nil {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, err)
		return
	}

	userResult, statusCode := h.service.UpdateUser(userId, &input)

	switch statusCode {
	case http.StatusOK:
		utils.APIResponse(ctx, "User updated successfully", statusCode, userResult)
		return
	case http.StatusNotFound:
		utils.APIResponse(ctx, "This doesn't user exists", statusCode, nil)
		return
	case http.StatusExpectationFailed:
		utils.APIResponse(ctx, "Unable to update an account", statusCode, nil)
		return
	default:
		utils.APIResponse(ctx, "Something went wrong", http.StatusBadRequest, nil)
	}
}
