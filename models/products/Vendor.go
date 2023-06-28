package products

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Vendor struct {
	ID        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Active    int8      `gorm:"default:1;not null;index" json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy int32     `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy int32     `json:"updatedBy"`
}

func (handle *Vendor) AllVendors(db *gorm.DB) (*[]Vendor, error) {
	var err error
	vendors := []Vendor{}
	err = db.Debug().Model(&Vendor{}).Where("active = ?", 1).Limit(100).Find(&vendors).Error
	if err != nil {
		return &[]Vendor{}, err
	}

	return &vendors, nil
}

func (handle *Vendor) PrepareVendor() {
	handle.ID = 0
	handle.Name = html.EscapeString(strings.TrimSpace(handle.Name))

	handle.CreatedAt = time.Now()
	handle.UpdatedAt = time.Now()
}

func (handle *Vendor) ValidateVendor() error {

	if handle.Name == "" {
		return errors.New("name required")
	}
	return nil
}

func (handle *Vendor) CreateVendor(db *gorm.DB) (*Vendor, error) {
	var err = db.Debug().Model(&Vendor{}).Create(&handle).Error
	if err != nil {
		return &Vendor{}, err
	}
	return handle, nil
}

func (handle *Vendor) VendorByName(db *gorm.DB, vendorName string) (*Vendor, error) {
	var err = db.Debug().Model(&Vendor{}).Where("name = ?", vendorName).Take(&handle).Error
	if err != nil {
		return &Vendor{}, err
	}
	return handle, nil
}

func (handle *Vendor) UpdateVendor(db *gorm.DB, id uint32) (*Vendor, error) {
	var err = db.Debug().Model(&Vendor{}).Where("id = ?", id).Updates(Vendor{Name: handle.Name, Active: handle.Active}).Error
	if err != nil {
		return &Vendor{}, err
	}
	return handle, nil
}

func (handle *Vendor) DeleteVendor(db *gorm.DB, id uint32) (int64, error) {

	db = db.Debug().Model(&Vendor{}).Where("id = ?", id).Take(&Vendor{}).Delete(&Vendor{})

	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("vendor not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (handle *Vendor) VendorById(db *gorm.DB, vendorId uint32) (*Vendor, error) {
	var err = db.Debug().Model(&Vendor{}).Where("id = ?", vendorId).Take(&handle).Error
	if err != nil {
		return &Vendor{}, err
	}
	return handle, nil
}
