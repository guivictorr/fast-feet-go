package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guivictorr/fast-feet-go/schemas"
)

type UpdateUserRequest struct {
	Name string       `json:"name"`
	Cpf  string       `json:"cpf"`
	Role schemas.Role `json:"role"`
}

type UpdateUserResponse struct {
	Name string       `json:"name"`
	Cpf  string       `json:"cpf"`
	Role schemas.Role `json:"role"`
}

func (r *UpdateUserRequest) Validate() error {
	if r.Name == "" && r.Cpf == "" && r.Role == "" {
		return fmt.Errorf("empty request")
	}
	if len(r.Cpf) > 0 && len(r.Cpf) != 11 {
		return fmt.Errorf("cpf must have 11 characters")
	}
	if len(r.Role) > 0 && !r.Role.IsValid() {
		return fmt.Errorf("this role is invalid: %v", r.Role)
	}
	return nil
}

func UpdateUserHandler(ctx *gin.Context) {
	userId := ctx.Param("id")
	request := UpdateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Name: request.Name,
		Cpf:  request.Cpf,
		Role: request.Role,
	}

	u := db.Where("id = ?", userId).Find(&user)

	if u.RowsAffected == 0 {
		logger.Errorf("this user doesn't exist")
		sendError(ctx, http.StatusConflict, "this user doesn't exist")
		return
	}

	if err := db.Model(&user).Updates(request).Error; err != nil {
		logger.Errorf("error updating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	sendSuccess(ctx, user, "update user")
}
