package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Modull struct {
	ID          string         `gorm:"type:text;primaryKey"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Nama        string         `json:"Nama" form:"Nama"`
	Deskripsi   string         `json:"Deskripsi" form:"Deskripsi"`
	Materi      []Materi       `gorm:"foreignKey:MateriID"`
	LatihanSoal []LatihanSoal  `gorm:"foreignKey:LatihanSoalID"`
}
