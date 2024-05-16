package utils

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Errors     interface{} `json:"errors"`
	Method     string      `json:"method"`
	StatusCode int         `json:"statusCode"`
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

func ValidatorErrorResponse(ctx *gin.Context, StatusCode int, Method string, Error interface{}) {
	errResponse := ErrorResponse{
		StatusCode: StatusCode,
		Method:     Method,
		Errors:     Error,
	}

	ctx.JSON(StatusCode, errResponse)
	defer ctx.AbortWithStatus(StatusCode)
}
