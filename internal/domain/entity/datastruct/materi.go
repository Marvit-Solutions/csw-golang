package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Materi struct {
	ID        string         `gorm:"type:text;primaryKey"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ModulID   string         `json:"ModullID" form:"ModulID"`
	Nama      string         `json:"Nama" form:"Nama"`
	SubMateri []SubMateri    `gorm:"foreignKey:SubMateriID"`
}
