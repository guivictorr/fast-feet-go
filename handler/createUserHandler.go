package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guivictorr/fast-feet-go/schemas"
)

type CreateUserRequest struct {
	Name     string       `json:"name"`
	Cpf      string       `json:"cpf"`
	Password string       `json:"password"`
	Role     schemas.Role `json:"role"`
}

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	// TODO: VALIDATE BODY

	user := schemas.User{
		Name:     request.Name,
		Cpf:      request.Cpf,
		Password: request.Password,
		Role:     request.Role,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating user")
		return
	}

	sendSuccess(ctx, user, "createUser")
}
