package products

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ProductType struct {
	ID        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Active    int8      `gorm:"default:1;not null;index" json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy int32     `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy int32     `json:"updatedBy"`
}

func (handle *ProductType) AllProductTypes(db *gorm.DB) (*[]ProductType, error) {
	var err error
	producttypes := []ProductType{}
	err = db.Debug().Model(&ProductType{}).Where("active = ?", 1).Limit(100).Find(&producttypes).Error
	if err != nil {
		return &[]ProductType{}, err
	}

	return &producttypes, nil
}

func (handle *ProductType) PrepareProductType() {
	handle.ID = 0
	handle.Name = html.EscapeString(strings.TrimSpace(handle.Name))

	handle.CreatedAt = time.Now()
	handle.UpdatedAt = time.Now()
}

func (handle *ProductType) ValidateProductType() error {

	if handle.Name == "" {
		return errors.New("name required")
	}
	return nil
}

func (handle *ProductType) CreateProductType(db *gorm.DB) (*ProductType, error) {
	var err = db.Debug().Model(&ProductType{}).Create(&handle).Error
	if err != nil {
		return &ProductType{}, err
	}
	return handle, nil
}

func (handle *ProductType) ProductTypeByName(db *gorm.DB, producttypeName string) (*ProductType, error) {
	var err = db.Debug().Model(&ProductType{}).Where("name = ?", producttypeName).Take(&handle).Error
	if err != nil {
		return &ProductType{}, err
	}
	return handle, nil
}

func (handle *ProductType) UpdateProductType(db *gorm.DB, id uint32) (*ProductType, error) {
	var err = db.Debug().Model(&ProductType{}).Where("id = ?", id).Updates(ProductType{Name: handle.Name, Active: handle.Active}).Error
	if err != nil {
		return &ProductType{}, err
	}
	return handle, nil
}

func (handle *ProductType) DeleteProductType(db *gorm.DB, id uint32) (int64, error) {

	db = db.Debug().Model(&ProductType{}).Where("id = ?", id).Take(&ProductType{}).Delete(&ProductType{})

	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("producttype not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (handle *ProductType) ProductTypeById(db *gorm.DB, producttypeId uint32) (*ProductType, error) {
	var err = db.Debug().Model(&ProductType{}).Where("id = ?", producttypeId).Take(&handle).Error
	if err != nil {
		return &ProductType{}, err
	}
	return handle, nil
}
