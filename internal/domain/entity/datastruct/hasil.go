package datastruct

import (
	"time"

	"gorm.io/gorm"
)

// UNTUK MENYIMPAN HASIL TES TIAP USER
type Hasil struct {
	ID            string         `gorm:"type:text;primaryKey"`
	CreatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	UserID        string         `json:"UserID" form:"UserID"`
	MateriID      string         `json:"MateriID" form:"MateriID"`
	LatihanSoalID string         `json:"LatihanSoalID" form:"LatihanSoalID"`
	Waktu         time.Duration  `json:"Waktu" form:"Waktu"`
	Hasil         float64        `json:"Hasil" form:"Hasil"`
}
