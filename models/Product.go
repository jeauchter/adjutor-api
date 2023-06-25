package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID            uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string    `gorm:"size:255;not null;unique" json:"name"`
	ProductTypeID int32     `gorm:"not null;index" json:"productTypeId"`
	StyleID       int32     `gorm:"not null;index" json:"styleId"`
	TagID         int32     `gorm:"not null;index" json:"tagId"`
	ClassID       int32     `gorm:"not null;index" json:"classId"`
	VendorID      int32     `gorm:"not null;index" json:"vendorId"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     int32     `json:"createdBy"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UpdatedBy     int32     `json:"updatedBy"`
	ProductType   ProductType
	Style         Style
	Tag           Tag
	Class         Class
	Vendor        Vendor
}
