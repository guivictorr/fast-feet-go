package schemas

import (
	"gorm.io/gorm"
)

type Role string

const (
	Admin       Role = "admin"
	DeliveryMan Role = "deliveryman"
)

type User struct {
	gorm.Model
	Name     string
	Cpf      string
	Password string
	Role     Role
}
