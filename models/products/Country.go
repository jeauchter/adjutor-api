package products

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Country struct {
	ID        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Active    int8      `gorm:"default:1;not null;index" json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy int32     `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy int32     `json:"updatedBy"`
}

func (handle *Country) AllCountrys(db *gorm.DB) (*[]Country, error) {
	var err error
	countrys := []Country{}
	err = db.Debug().Model(&Country{}).Where("active = ?", 1).Limit(100).Find(&countrys).Error
	if err != nil {
		return &[]Country{}, err
	}

	return &countrys, nil
}

func (handle *Country) PrepareCountry() {
	handle.ID = 0
	handle.Name = html.EscapeString(strings.TrimSpace(handle.Name))

	handle.CreatedAt = time.Now()
	handle.UpdatedAt = time.Now()
}

func (handle *Country) ValidateCountry() error {

	if handle.Name == "" {
		return errors.New("name required")
	}
	return nil
}

func (handle *Country) CreateCountry(db *gorm.DB) (*Country, error) {
	var err = db.Debug().Model(&Country{}).Create(&handle).Error
	if err != nil {
		return &Country{}, err
	}
	return handle, nil
}

func (handle *Country) CountryByName(db *gorm.DB, countryName string) (*Country, error) {
	var err = db.Debug().Model(&Country{}).Where("name = ?", countryName).Take(&handle).Error
	if err != nil {
		return &Country{}, err
	}
	return handle, nil
}

func (handle *Country) UpdateCountry(db *gorm.DB, id uint32) (*Country, error) {
	var err = db.Debug().Model(&Country{}).Where("id = ?", id).Updates(Country{Name: handle.Name, Active: handle.Active}).Error
	if err != nil {
		return &Country{}, err
	}
	return handle, nil
}

func (handle *Country) DeleteCountry(db *gorm.DB, id uint32) (int64, error) {

	db = db.Debug().Model(&Country{}).Where("id = ?", id).Take(&Country{}).Delete(&Country{})

	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("country not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
