package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Modul struct {
	ID          string         `gorm:"type:text;primaryKey"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Nama        string         `json:"Nama" form:"Nama"`
	Deskripsi   string         `json:"Deskripsi" form:"Deskripsi"`
	Materi      []Materi       `gorm:"foreignKey:ModulID"`
	// LatihanSoal []LatihanSoal  `gorm:"foreignKey:MateriID"`
}
