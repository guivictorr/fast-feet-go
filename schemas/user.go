package schemas

import (
	"gorm.io/gorm"
)

type Role string

const (
	None    Role = ""
	Admin   Role = "admin"
	Courier Role = "courier"
)

func (r Role) IsValid() bool {
	switch r {
	case Admin, Courier:
		return true
	// Add more roles here...
	default:
		return false
	}
}

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Cpf      string `gorm:"unique;not null;size:11"`
	Password string `gorm:"not null" json:"-"`
	Role     Role   `gorm:"not null"`
	Packages []Package
}
