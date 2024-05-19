package authHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authControllers "github.com/guivictorr/fast-feet-go/controllers/auth-controllers"
	"github.com/guivictorr/fast-feet-go/utils"
)

type handler struct {
	service authControllers.Service
}

func NewAuthHandler(service authControllers.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateSessionHandler(ctx *gin.Context) {
	var input authControllers.SessionInput

	ctx.ShouldBindJSON(&input)

	err := utils.GoValidator(&input)
	if err != nil {
		utils.ValidatorErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}

	sessionResult, statusCode := h.service.CreateSession(&input)

	switch statusCode {
	case http.StatusAccepted:
		utils.APIResponse(ctx, "Session created successfully", statusCode, sessionResult)
		return
	case http.StatusNotFound:
		utils.APIResponse(ctx, "This user doesn't exists", statusCode, nil)
		return
	case http.StatusUnauthorized:
		utils.APIResponse(ctx, "Cpf or Password incorrect", statusCode, nil)
		return
	default:
		utils.APIResponse(ctx, "Something went wrong", http.StatusBadRequest, nil)
	}
}
