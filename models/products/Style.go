package products

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Style struct {
	gorm.Model
	ID           uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"size:255;not null;unique" json:"name"`
	DepartmentID int32     `gorm:"not null;index" json:"departmentId"`
	AudienceID   int32     `gorm:"not null;index" json:"audienceId"`
	CreatedAt    time.Time `json:"createdAt"`
	CreatedBy    int32     `json:"createdBy"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UpdatedBy    int32     `json:"updatedBy"`
	Department   Department
	Audience     Audience
}

func (handle *Style) AllStyles(db *gorm.DB) (*[]Style, error) {
	var err error
	styles := []Style{}
	err = db.Debug().Model(&Style{}).Preload("Department").Preload("Audience").Find(&styles).Error
	if err != nil {
		return &[]Style{}, err
	}

	return &styles, nil
}

func (handle *Style) PrepareStyle() {
	handle.ID = 0
	handle.Name = html.EscapeString(strings.TrimSpace(handle.Name))

	handle.CreatedAt = time.Now()
	handle.UpdatedAt = time.Now()
}

func (handle *Style) ValidateStyle() error {

	if handle.Name == "" {
		return errors.New("name required")
	}
	return nil
}

func (handle *Style) CreateStyle(db *gorm.DB) (*Style, error) {
	var err = db.Debug().Model(&Style{}).Preload("Department").Preload("Audience").Create(&handle).Error
	if err != nil {
		return &Style{}, err
	}
	return handle, nil
}

func (handle *Style) StyleByName(db *gorm.DB, styleName string) (*Style, error) {
	var err = db.Debug().Model(&Style{}).Where("name = ?", styleName).Take(&handle).Error
	if err != nil {
		return &Style{}, err
	}
	return handle, nil
}

func (handle *Style) UpdateStyle(db *gorm.DB, id uint32) (*Style, error) {
	var err = db.Debug().Model(&Style{}).Where("id = ?", id).Updates(Style{Name: handle.Name, DepartmentID: handle.DepartmentID, AudienceID: handle.AudienceID}).Error
	if err != nil {
		return &Style{}, err
	}
	return handle, nil
}

func (handle *Style) DeleteStyle(db *gorm.DB, id uint32) (int64, error) {

	db = db.Debug().Model(&Style{}).Where("id = ?", id).Take(&Style{}).Delete(&Style{})

	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("style not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (handle *Style) StyleById(db *gorm.DB, styleId uint32) (*Style, error) {
	var err = db.Debug().Model(&Style{}).Preload("Department").Preload("Audience").Where("id = ?", styleId).Take(&handle).Error
	if err != nil {
		return &Style{}, err
	}
	return handle, nil
}
