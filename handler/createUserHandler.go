package handler

import (
	"fmt"
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
		// TODO: CREATE A SENDERROR FUNC
		ctx.Header("Content-type", "application/json")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message":   "error creating user on database",
			"errorCode": http.StatusInternalServerError,
		})
		return
	}

	// TODO: CREATE A SENDSUCCESS FUNC
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", "createUser"),
		"data":    user,
	})
}
