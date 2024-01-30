package datastruct

import (
	"time"

	"gorm.io/gorm"
)

// MENYIMPAN HASIL JAWABAN DARI TIAP USER
type StoredAnswer struct {
	ID        string         `gorm:"type:text;primaryKey"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    string         `json:"UserID" form:"UserID"`
	SoalID    string         `json:"SoalID" form:"SoalID"`
	Jenis     string         `json:"Jenis" form:"Jenis"`
	Konten    string         `json:"Konten" form:"Konten"`
	Dipilih   bool           `json:"Dipilih" form:"Dipilih"`
	Nilai     float64        `json:"Nilai" form:"Nilai"`
}
