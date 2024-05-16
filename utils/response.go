package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Method     string   `json:"method"`
	Errors     []string `json:"errors"`
	StatusCode int      `json:"statusCode"`
}
type Responses struct {
	Data       interface{} `json:"data"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
}

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Method string, Data interface{}) {
	jsonResponse := Responses{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}

	if StatusCode >= 400 {
		ctx.JSON(StatusCode, jsonResponse)
		defer ctx.AbortWithStatus(StatusCode)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func ValidatorErrorResponse(ctx *gin.Context, StatusCode int, Method string, Errors []validator.FieldError) {
	errorMessages := []string{}
	for _, err := range Errors {
		message := err.Error()
		errorMessages = append(errorMessages, message)
	}
	errResponse := ErrorResponse{
		StatusCode: StatusCode,
		Method:     Method,
		Errors:     errorMessages,
	}

	ctx.JSON(StatusCode, errResponse)
	defer ctx.AbortWithStatus(StatusCode)
}
