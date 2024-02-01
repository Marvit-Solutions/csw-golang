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
	Province     string         `json:"Province" form:"Province"`       //Province
	RegencyCity  string         `json:"RegencyCity" form:"RegencyCity"` //RegencyCity/Kota
	SubDistrict  string         `json:"SubDistrict" form:"SubDistrict"` //Subdistrict
	UserDetailID string         `json:"UserDetailID" form:"UserDetailID"`
}
