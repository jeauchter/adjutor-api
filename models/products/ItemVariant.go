package products

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ItemVariant struct {
	ID        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Active    int8      `gorm:"default:1;not null;index" json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy int32     `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy int32     `json:"updatedBy"`
}

func (handle *ItemVariant) AllItemVariants(db *gorm.DB) (*[]ItemVariant, error) {
	var err error
	itemvariants := []ItemVariant{}
	err = db.Debug().Model(&ItemVariant{}).Where("active = ?", 1).Limit(100).Find(&itemvariants).Error
	if err != nil {
		return &[]ItemVariant{}, err
	}

	return &itemvariants, nil
}

func (handle *ItemVariant) PrepareItemVariant() {
	handle.ID = 0
	handle.Name = html.EscapeString(strings.TrimSpace(handle.Name))

	handle.CreatedAt = time.Now()
	handle.UpdatedAt = time.Now()
}

func (handle *ItemVariant) ValidateItemVariant() error {

	if handle.Name == "" {
		return errors.New("name required")
	}
	return nil
}

func (handle *ItemVariant) CreateItemVariant(db *gorm.DB) (*ItemVariant, error) {
	var err = db.Debug().Model(&ItemVariant{}).Create(&handle).Error
	if err != nil {
		return &ItemVariant{}, err
	}
	return handle, nil
}

func (handle *ItemVariant) ItemVariantByName(db *gorm.DB, itemvariantName string) (*ItemVariant, error) {
	var err = db.Debug().Model(&ItemVariant{}).Where("name = ?", itemvariantName).Take(&handle).Error
	if err != nil {
		return &ItemVariant{}, err
	}
	return handle, nil
}

func (handle *ItemVariant) UpdateItemVariant(db *gorm.DB, id uint32) (*ItemVariant, error) {
	var err = db.Debug().Model(&ItemVariant{}).Where("id = ?", id).Updates(ItemVariant{Name: handle.Name, Active: handle.Active}).Error
	if err != nil {
		return &ItemVariant{}, err
	}
	return handle, nil
}

func (handle *ItemVariant) DeleteItemVariant(db *gorm.DB, id uint32) (int64, error) {

	db = db.Debug().Model(&ItemVariant{}).Where("id = ?", id).Take(&ItemVariant{}).Delete(&ItemVariant{})

	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("itemvariant not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (handle *ItemVariant) ItemVariantById(db *gorm.DB, itemvariantId uint32) (*ItemVariant, error) {
	var err = db.Debug().Model(&ItemVariant{}).Where("id = ?", itemvariantId).Take(&handle).Error
	if err != nil {
		return &ItemVariant{}, err
	}
	return handle, nil
}
