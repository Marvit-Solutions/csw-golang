package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        string         `gorm:"type:text;primaryKey"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Role      string         `json:"Role" form:"Role"`
	Users     []User
}
