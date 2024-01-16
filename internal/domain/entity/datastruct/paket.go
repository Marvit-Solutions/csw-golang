package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Paket struct {
	ID        string         `gorm:"type:text;primaryKey"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Nama      string         `json:"NamaPaket" form:"NamaPaket"`
	SubPaket  []SubPaket     `gorm:"foreignKey:PaketID"`
}
