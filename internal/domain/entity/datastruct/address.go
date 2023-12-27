package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID           string `gorm:"type:text;primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Provinsi     string         `json:"Provinsi" form:"Provinsi"`
	Kabupaten    string         `json:"Kabupaten" form:"Kabupaten"`
	Kecamatan    string         `json:"Kecamatan" form:"Kecamatan"`
	UserDetailId string         `json:"UserDetailId" form:"UserDetailId"`
}
