package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Role      string         `json:"Role" form:"Role"`
	Users     []User
}
