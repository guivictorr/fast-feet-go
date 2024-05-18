package userControllers

import "github.com/guivictorr/fast-feet-go/models"

type UserInput struct {
	Name     string      `json:"name" validate:"required"`
	Cpf      string      `json:"cpf" validate:"required,lowercase,len=11"`
	Password string      `json:"password" validate:"required,gte=8"`
	Role     models.Role `json:"role" validate:"required,oneof=admin courier"`
}
type UpdateUserInput struct {
	Name     string      `json:"name" validate:"omitempty"`
	Cpf      string      `json:"cpf" validate:"omitempty,len=11"`
	Password string      `json:"password" validate:"omitempty,gte=8"`
	Role     models.Role `json:"role" validate:"omitempty,oneof=admin courier"`
}
