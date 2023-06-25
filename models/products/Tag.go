package products

import (
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Active    int8      `gorm:"default:1;not null;index" json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy int32     `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy int32     `json:"updatedBy"`
}

func (p *Tag) AllTags(db *gorm.DB) (*[]Tag, error) {
	var err error
	tags := []Tag{}
	err = db.Debug().Model(&Tag{}).Where("active = ?", 1).Limit(100).Find(&tags).Error
	if err != nil {
		return &[]Tag{}, err
	}

	return &tags, nil
}
