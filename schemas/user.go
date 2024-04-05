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
	Name     string `gorm:"not null"`
	Cpf      string `gorm:"unique;not null;size:11"`
	Password string `gorm:"not null"`
	Role     Role   `gorm:"not null"`
	Packages []Package
}
