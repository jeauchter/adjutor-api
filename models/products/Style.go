package products

import (
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
