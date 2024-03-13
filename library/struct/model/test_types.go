package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameTestType = "test_types"

// TestType mapped from table <test_types>
type TestType struct {
	ID        string         `gorm:"column:id;primaryKey" json:"id"`
	Name      string         `gorm:"column:name;not null" json:"name"`
	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

// TableName TestType's table name
func (*TestType) TableName() string {
	return TableNameTestType
}
