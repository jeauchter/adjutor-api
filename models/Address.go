package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	ID           uint32    `gorm:"primaryKey;autoIncrement" json:"id"`
	Unit         string    `gorm:"size:255;index" json:"unit"`
	StreetNumber string    `gorm:"size:255;index" json:"streetNumber"`
	AddressOne   string    `gorm:"size:255;index" json:"addressOne"`
	AddressTwo   string    `gorm:"size:255;index" json:"addressTwo"`
	City         string    `gorm:"size:255;index" json:"city"`
	Region       string    `gorm:"size:255;index" json:"region"`
	PostalCode   string    `gorm:"size:255;index" json:"postalCode"`
	CountryID    int32     `gorm:"index" json:"countryId"`
	CreatedAt    time.Time `json:"createdAt"`
	CreatedBy    int32     `json:"createdBy"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UpdatedBy    int32     `json:"updatedBy"`
	Country      Country
}
