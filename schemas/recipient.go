package schemas

import (
	"gorm.io/gorm"
)

type Recipient struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Address  string `gorm:"not null"`
	Packages []Package
}
