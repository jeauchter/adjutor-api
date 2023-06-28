package products

import (
	"errors"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Audience struct {
	ID        uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Active    int8      `gorm:"default:1;not null;index" json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy int32     `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	UpdatedBy int32     `json:"updatedBy"`
}

func (handle *Audience) AllAudiences(db *gorm.DB) (*[]Audience, error) {
	var err error
	audiences := []Audience{}
	err = db.Debug().Model(&Audience{}).Where("active = ?", 1).Limit(100).Find(&audiences).Error
	if err != nil {
		return &[]Audience{}, err
	}

	return &audiences, nil
}

func (handle *Audience) PrepareAudience() {
	handle.ID = 0
	handle.Name = html.EscapeString(strings.TrimSpace(handle.Name))

	handle.CreatedAt = time.Now()
	handle.UpdatedAt = time.Now()
}

func (handle *Audience) ValidateAudience() error {

	if handle.Name == "" {
		return errors.New("name required")
	}
	return nil
}

func (handle *Audience) CreateAudience(db *gorm.DB) (*Audience, error) {
	var err = db.Debug().Model(&Audience{}).Create(&handle).Error
	if err != nil {
		return &Audience{}, err
	}
	return handle, nil
}

func (handle *Audience) AudienceByName(db *gorm.DB, audienceName string) (*Audience, error) {
	var err = db.Debug().Model(&Audience{}).Where("name = ?", audienceName).Take(&handle).Error
	if err != nil {
		return &Audience{}, err
	}
	return handle, nil
}

func (handle *Audience) UpdateAudience(db *gorm.DB, id uint32) (*Audience, error) {
	var err = db.Debug().Model(&Audience{}).Where("id = ?", id).Updates(Audience{Name: handle.Name, Active: handle.Active}).Error
	if err != nil {
		return &Audience{}, err
	}
	return handle, nil
}

func (handle *Audience) DeleteAudience(db *gorm.DB, id uint32) (int64, error) {

	db = db.Debug().Model(&Audience{}).Where("id = ?", id).Take(&Audience{}).Delete(&Audience{})

	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("audience not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}

func (handle *Audience) AudienceById(db *gorm.DB, audienceId uint32) (*Audience, error) {
	var err = db.Debug().Model(&Audience{}).Where("id = ?", audienceId).Take(&handle).Error
	if err != nil {
		return &Audience{}, err
	}
	return handle, nil
}
