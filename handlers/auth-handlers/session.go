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
		accessTokenData := map[string]interface{}{"id": sessionResult.ID, "cpf": sessionResult.Cpf, "role": sessionResult.Role}
		accessToken, errToken := utils.Sign(accessTokenData, "JWT_SECRET", 24*60*1)

		if errToken != nil {
			utils.APIResponse(ctx, "Generate accessToken failed", http.StatusBadRequest, nil)
			return
		}

		utils.APIResponse(ctx, "Login successfully", http.StatusOK, accessToken)
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
