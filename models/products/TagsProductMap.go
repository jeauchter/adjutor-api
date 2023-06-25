package products

import (
	"time"
)

type TagProductMap struct {
	ID        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	TagID     int32     `gorm:"not null;index" json:"tagId"`
	ProductID int32     `gorm:"not null;index" json:"productId"`
	Active    int8      `gorm:"default:1;not null;index" json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy int32     `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy int32     `json:"updatedBy"`
	Product   Product
	Tag       Tag
}
