package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Addresses struct {
	ID           string `gorm:"type:text;primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Province     string         `json:"Province" form:"Province"`       //Province
	RegencyCity  string         `json:"RegencyCity" form:"RegencyCity"` //RegencyCity/Kota
	Subdistrict  string         `json:"Subdistrict" form:"Subdistrict"` //Subdistrict
	UserDetailID string         `json:"UserDetailID" form:"UserDetailID"`
}
