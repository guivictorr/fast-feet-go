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
	gorm.Model
	Name            string
	PickupDate      *time.Time
	DeliveryDate    *time.Time
	DeliveryPicture string
	Address         string
	Status          PackageStatus `gorm:"default:pending"`
	RecipientID     uint
}
