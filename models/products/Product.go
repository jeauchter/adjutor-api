package products

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID            uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string    `gorm:"size:255;not null;unique" json:"name"`
	ProductTypeID int32     `gorm:"not null;index" json:"productTypeId"`
	StyleID       int32     `gorm:"not null;index" json:"styleId"`
	ClassID       int32     `gorm:"not null;index" json:"classId"`
	VendorID      int32     `gorm:"not null;index" json:"vendorId"`
	Active        int8      `gorm:"default:1;not null;index" json:"active"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     int32     `json:"createdBy"`
	UpdatedAt     time.Time `json:"updatedAt"`
	UpdatedBy     int32     `json:"updatedBy"`
	ProductType   ProductType
	Style         Style
	Class         Class
	Vendor        Vendor
}

func (handle *Product) AllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Preload("Style").Preload("Class").Preload("ProductType").Preload("Vendor").Where("active = ?", 1).Limit(100).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}

	return &products, nil
}

func (handle *Product) PrepareProduct() {
	handle.ID = 0
	handle.Name = html.EscapeString(strings.TrimSpace(handle.Name))

	handle.CreatedAt = time.Now()
	handle.UpdatedAt = time.Now()
}

func (handle *Product) ValidateProduct() error {

	if handle.Name == "" {
		return errors.New("name required")
	}

	return nil
}

func (handle *Product) CreateProduct(db *gorm.DB) (*Product, error) {
	var err = db.Debug().Model(&Product{}).Create(&handle).Error
	if err != nil {
		return &Product{}, err
	}
	return handle, nil
}

func (handle *Product) ProductByName(db *gorm.DB, productName string) (*Product, error) {
	var err = db.Debug().Model(&Product{}).Where("name = ?", productName).Take(&handle).Error
	if err != nil {
		return &Product{}, err
	}
	return handle, nil
}

func (handle *Product) ProductById(db *gorm.DB, productId uint32) (*Product, error) {
	var err = db.Debug().Model(&Product{}).Preload("Style").Preload("Class").Preload("ProductType").Preload("Vendor").Where("id = ?", productId).Take(&handle).Error
	if err != nil {
		return &Product{}, err
	}
	return handle, nil
}

func (handle *Product) UpdateProduct(db *gorm.DB, id uint32) (*Product, error) {
	var err = db.Debug().Model(&Product{}).Where("id = ?", id).Updates(Product{Name: handle.Name, Active: handle.Active}).Error
	if err != nil {
		return &Product{}, err
	}
	return handle, nil
}

func (handle *Product) DeleteProduct(db *gorm.DB, id uint32) (int64, error) {

	db = db.Debug().Model(&Product{}).Where("id = ?", id).Take(&Product{}).Delete(&Product{})

	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("product not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
