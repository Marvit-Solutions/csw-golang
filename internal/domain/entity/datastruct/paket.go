package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Paket struct {
	ID             string `gorm:"type:text;primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	NamaPaket      string         `json:"NamaPaket" form:"NamaPaket"`
	DeskripsiPaket string         `json:"DeskripsiPaket" form:"DeskripsiPaket"`
}
