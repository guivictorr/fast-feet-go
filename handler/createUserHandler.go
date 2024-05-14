package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guivictorr/fast-feet-go/schemas"
	"golang.org/x/crypto/bcrypt"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type CreateUserRequest struct {
	Name     string       `json:"name"`
	Cpf      string       `json:"cpf"`
	Password string       `json:"password"`
	Role     schemas.Role `json:"role"`
}

type CreateUserResponse struct {
	Name string       `json:"name"`
	Cpf  string       `json:"cpf"`
	Role schemas.Role `json:"role"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Name == "" && r.Cpf == "" && r.Password == "" && r.Role == "" {
		return fmt.Errorf("empty request")
	}

	if r.Name == "" {
		return errParamIsRequired("Name", "string")
	}
	if r.Cpf == "" {
		return errParamIsRequired("CPF", "string")
	}
	if r.Password == "" {
		return errParamIsRequired("Password", "string")
	}
	if r.Role == "" {
		return errParamIsRequired("Role", "string")
	}
	return nil
}

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	pwd, err := generatePassword(request.Password)
	if err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusUnauthorized, "cpf or password incorrect")
		return
	}

	request.Password = pwd

	user := schemas.User{
		Name:     request.Name,
		Cpf:      request.Cpf,
		Password: request.Password,
		Role:     request.Role,
	}

	u := db.Where("Cpf = ?", request.Cpf).Find(&user)

	if u.RowsAffected == 1 {
		logger.Errorf("this user already exists")
		sendError(ctx, http.StatusConflict, "user already exists")
		return
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating user: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating user")
		return
	}

	userResponse := CreateUserResponse{
		Name: user.Name,
		Cpf:  user.Cpf,
		Role: user.Role,
	}
	sendSuccess(ctx, userResponse, "createUser")
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
