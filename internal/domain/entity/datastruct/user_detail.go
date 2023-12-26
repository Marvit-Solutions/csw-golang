package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type UserDetail struct {
	ID        string `gorm:"type:text;primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"Name" form:"Name" validate:"required"`
	Phone     string         `json:"Phone" form:"Phone" validate:"required,min=10,max=13,numeric"`
	UserId    string         `json:"UserId" form:"UserId"`
}
