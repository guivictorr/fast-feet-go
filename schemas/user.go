package schemas

import (
	"time"

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

type UserResponse struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Password  string    `json:"password"`
	Role      Role      `json:"role"`
	ID        uint      `json:"id"`
}
