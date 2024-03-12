package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAddress = "addresses"

// Address mapped from table <addresses>
type Address struct {
	ID           string         `gorm:"column:id;primaryKey" json:"id"`
	UserDetailID string         `gorm:"column:user_detail_id;not null" json:"user_detail_id"`
	Province     string         `gorm:"column:province;not null" json:"province"`
	RegencyCity  string         `gorm:"column:regency_city;not null" json:"regency_city"`
	SubDistrict  string         `gorm:"column:sub_district;not null" json:"sub_district"`
	CreatedAt    time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

// TableName Address's table name
func (*Address) TableName() string {
	return TableNameAddress
}
