package utils

import (
	"github.com/go-playground/validator/v10"
)

type Error struct {
	Results struct {
		Errors []map[string]map[string]string `json:"errors"`
	} `json:"results"`
}

func GoValidator(s interface{}) []validator.FieldError {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(s)
	if err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}
