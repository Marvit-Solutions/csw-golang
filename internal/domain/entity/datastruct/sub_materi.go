package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type SubMateri struct {
	ID        string         `gorm:"type:text;primaryKey"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	MateriID  string         `json:"MateriID" form:"MateriID"`
	Nama      string         `json:"Nama" form:"Nama"`
	Konten    string         `json:"Konten" form:"Konten"`
}
