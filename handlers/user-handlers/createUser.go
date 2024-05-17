package userHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userControllers "github.com/guivictorr/fast-feet-go/controllers/user-controllers"
	"github.com/guivictorr/fast-feet-go/utils"
)

func (h *handler) CreateUserHandler(ctx *gin.Context) {
	var input userControllers.UserInput
	ctx.ShouldBindJSON(&input)

	err := utils.GoValidator(&input)
	if err != nil {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, err)
		return
	}

	userResult, statusCode := h.service.CreateUser(&input)

	switch statusCode {
	case http.StatusCreated:
		utils.APIResponse(ctx, "User created successfully", statusCode, userResult)
		return
	case http.StatusConflict:
		utils.APIResponse(ctx, "This user already exists", statusCode, nil)
		return
	case http.StatusExpectationFailed:
		utils.APIResponse(ctx, "Unable to create an account", statusCode, nil)
		return
	default:
		utils.APIResponse(ctx, "Something went wrong", http.StatusBadRequest, nil)
	}
}
