package products

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Tag struct {
	ID        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Active    int8      `gorm:"default:1;not null;index" json:"active"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:DATETIME;not null;default:CURRENT_TIMESTAMP"`
	CreatedBy int32     `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;not null;default:CURRENT_TIMESTAMP"`
	UpdatedBy int32     `json:"updatedBy"`
}

func (t *Tag) AllTags(db *gorm.DB) (*[]Tag, error) {
	var err error
	tags := []Tag{}
	err = db.Debug().Model(&Tag{}).Where("active = ?", 1).Limit(100).Find(&tags).Error
	if err != nil {
		return &[]Tag{}, err
	}

	return &tags, nil
}

func (handle *Tag) PrepareTag() {
	handle.ID = 0
	handle.Name = html.EscapeString(strings.TrimSpace(handle.Name))

	handle.CreatedAt = time.Now()
	handle.UpdatedAt = time.Now()
}

func (handle *Tag) ValidateTag() error {

	if handle.Name == "" {
		return errors.New("name required")
	}
	return nil
}

func (handle *Tag) CreateTag(db *gorm.DB) (*Tag, error) {
	var err = db.Debug().Model(&Tag{}).Create(&handle).Error
	if err != nil {
		return &Tag{}, err
	}
	return handle, nil
}

func (handle *Tag) TagByName(db *gorm.DB, tagName string) (*Tag, error) {
	var err = db.Debug().Model(&Tag{}).Where("name = ?", tagName).Take(&handle).Error
	if err != nil {
		return &Tag{}, err
	}
	return handle, nil
}

func (handle *Tag) TagById(db *gorm.DB, tagId uint32) (*Tag, error) {
	var err = db.Debug().Model(&Tag{}).Where("id = ?", tagId).Take(&handle).Error
	if err != nil {
		return &Tag{}, err
	}
	return handle, nil
}

func (handle *Tag) UpdateTag(db *gorm.DB, id uint32) (*Tag, error) {
	var err = db.Debug().Model(&Tag{}).Where("id = ?", id).Updates(Tag{Name: handle.Name, Active: handle.Active}).Error
	if err != nil {
		return &Tag{}, err
	}
	return handle, nil
}

func (handle *Tag) DeleteTag(db *gorm.DB, id uint32) (int64, error) {

	db = db.Debug().Model(&Tag{}).Where("id = ?", id).Take(&Tag{}).Delete(&Tag{})

	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("tag not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
