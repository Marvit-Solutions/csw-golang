package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type SubPaket struct {
	ID                string `gorm:"type:text;primaryKey"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	PaketID           uint           `json:"IdPaket" form:"IdPaket"`
	NamaSubPaket      string         `json:"NamaPaket" form:"NamaPaket"`
	DeskripsiSubPaket string         `json:"DeskripsiPaket" form:"DeskripsiPaket"`
	Harga             int            `json:"Harga" form:"Harga"`
}
