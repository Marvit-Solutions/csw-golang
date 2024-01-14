package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type SubPaket struct {
	ID             string         `gorm:"type:text;primaryKey"`
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	PaketID        string         `json:"PaketID" form:"PaketID"`
	Nama           string         `json:"Nama" form:"Nama"`
	SubPaketDetail SubPaketDetail `gorm:"foreignKey:SubPaketID"`
	Harga          float64        `json:"Harga" form:"Harga"`
}
