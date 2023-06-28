package products

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID            uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID     int32     `gorm:"not null;index" json:"productId"`
	ItemVariantID int32     `gorm:"not null;index" json:"itemVariantId"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     int32     `json:"createdBy"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UpdatedBy     int32     `json:"updatedBy"`
	ItemVariant   ItemVariant
	Product       Product
}
