package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	ID              uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID       int32     `gorm:"not null;index" json:"productId"`
	ItemAttributeID int32     `gorm:"not null;index" json:"itemAttributeId"`
	CreatedAt       time.Time `json:"createdAt"`
	CreatedBy       int32     `json:"createdBy"`
	UpdatedAt       time.Time `json:"updatedAt"`
	UpdatedBy       int32     `json:"updatedBy"`
	ItemAttribute   ItemAttribute
	Product         Product
}
