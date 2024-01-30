package datastruct

import (
	"time"

	"gorm.io/gorm"
)

// MENYIMPAN PILIHAN JAWABAN DARI TIAP SOAL
type Answer struct {
	ID        string         `gorm:"type:text;primaryKey"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	SoalID    string         `json:"SoalID" form:"SoalID"`
	Jenis     string         `json:"Jenis" form:"Jenis"`
	Konten    string         `json:"Konten" form:"Konten"`
	Nilai     float64        `json:"Nilai" form:"Nilai"`
}
