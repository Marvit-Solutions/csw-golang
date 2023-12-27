package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type UserDetail struct {
	ID           string `gorm:"type:text;primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Nama         string         `json:"Nama" form:"Nama" validate:"required"`
	Telepon      string         `json:"Telepon" form:"Telepon" validate:"required,min=10,max=13,numeric"`
	FotoProfil   string         `json:"FotoProfil" form:"FotoProfil" gorm:"default: 'assets/img/account.png'"`
	TanggalLahir time.Time      `json:"TanggalLahir" form:"TanggalLahir" validate:"required"`
	Alamat       Address        `gorm:"foreignKey:UserDetailId"`
	UserId       string         `json:"UserId" form:"UserId"`
}
