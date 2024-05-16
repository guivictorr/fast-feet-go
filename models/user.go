package models

import (
	"time"

	"github.com/guivictorr/fast-feet-go/utils"
	"gorm.io/gorm"
)

type UserEntity struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"column:name;not null"`
	Cpf       string `gorm:"column:cpf;unique;not null;size:11"`
	Password  string `gorm:"column:password;not null" json:"-"`
	Role      Role   `gorm:"column:role;not null"`
	ID        uint   `gorm:"primary_key"`
}

func (entity *UserEntity) BeforeCreate(db *gorm.DB) error {
	entity.Password = utils.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *UserEntity) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}

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
	default:
		return false
	}
}
