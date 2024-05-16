package utils

import (
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type Error struct {
	Results struct {
		Errors []map[string]map[string]string `json:"errors"`
	} `json:"results"`
}

// IMPROVE THIS VALIDATOR CAUSE THE ERROR OBJECT IS RETURNING EMPTY
func GoValidator(s interface{}) error {
	validator := gpc.GoValidator()
	err := validator.Struct(s)
	// res, err := gpc.Validator(s)
	// if err != nil {
	// 	panic(err)
	// }

	return err
}
