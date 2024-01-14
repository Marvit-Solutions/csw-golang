package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type SubPaketDetail struct {
	ID          string         `gorm:"type:text;primaryKey"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	SubPaketID  string         `json:"SubPaketID" form:"SubPaketID"`
	GrupPejuang bool           `json:"GrupPejuang" form:"GrupPejuang"`
	SoalLatihan uint           `json:"SoalLatihan" form:"SoalLatihan"`
	Akses       uint           `json:"Akses" form:"Akses"`
	Modul       bool           `json:"Modul" form:"Modul"`
	TryOut      uint           `json:"TryOut" form:"TryOut"`
	Zoom        bool           `json:"Zoom" form:"Zoom"`
}
