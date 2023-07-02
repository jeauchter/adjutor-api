package products

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Department struct {
	ID        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Active    int8      `gorm:"default:1;not null;index" json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy int32     `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy int32     `json:"updatedBy"`
}

func (handle *Department) AllDepartments(db *gorm.DB) (*[]Department, error) {
	var err error
	departments := []Department{}
	err = db.Debug().Model(&Department{}).Where("active = ?", 1).Limit(100).Find(&departments).Error
	if err != nil {
		return &[]Department{}, err
	}

	return &departments, nil
}

func (handle *Department) PrepareDepartment() {
	handle.ID = 0
	handle.Name = html.EscapeString(strings.TrimSpace(handle.Name))

	handle.CreatedAt = time.Now()
	handle.UpdatedAt = time.Now()
}

func (handle *Department) ValidateDepartment() error {

	if handle.Name == "" {
		return errors.New("name required")
	}
	return nil
}

func (handle *Department) CreateDepartment(db *gorm.DB) (*Department, error) {
	var err = db.Debug().Model(&Department{}).Create(&handle).Error
	if err != nil {
		return &Department{}, err
	}
	return handle, nil
}

func (handle *Department) DepartmentByName(db *gorm.DB, departmentName string) (*Department, error) {
	var err = db.Debug().Model(&Department{}).Where("name = ?", departmentName).Take(&handle).Error
	if err != nil {
		return &Department{}, err
	}
	return handle, nil
}

func (handle *Department) UpdateDepartment(db *gorm.DB, id uint32) (*Department, error) {
	var err = db.Debug().Model(&Department{}).Where("id = ?", id).Updates(Department{Name: handle.Name, Active: handle.Active}).Error
	if err != nil {
		return &Department{}, err
	}
	return handle, nil
}

func (handle *Department) DeleteDepartment(db *gorm.DB, id uint32) (int64, error) {

	db = db.Debug().Model(&Department{}).Where("id = ?", id).Take(&Department{}).Delete(&Department{})

	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("department not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (handle *Department) DepartmentById(db *gorm.DB, departmentId uint32) (*Department, error) {
	var err = db.Debug().Model(&Department{}).Where("id = ?", departmentId).Take(&handle).Error
	if err != nil {
		return &Department{}, err
	}
	return handle, nil
}
