package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Module struct {
	ID        string         `gorm:"type:text;primaryKey"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"Name" form:"Name"`
	SubModule []SubModule    `gorm:"foreignKey:ModuleID"`
}
