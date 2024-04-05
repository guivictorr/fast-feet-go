package schemas

import (
	"time"

	"gorm.io/gorm"
)

type PackageStatus string

const (
	Pending   PackageStatus = "pending"
	InTransit PackageStatus = "in_transit"
	Delivered PackageStatus = "delivered"
	Returned  PackageStatus = "returned"
)

type Package struct {
	PickupDate      *time.Time
	DeliveryDate    *time.Time
	DeliveryPicture *string
	UserId          *uint
	gorm.Model
	Name        string        `gorm:"not null"`
	Address     string        `gorm:"not null"`
	Status      PackageStatus `gorm:"default:pending"`
	RecipientID uint          `gorm:"not null"`
}
