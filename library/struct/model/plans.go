package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePlan = "plans"

// Plan mapped from table <plans>
type Plan struct {
	ID          string         `gorm:"column:id;primaryKey" json:"id"`
	ModuleID    string         `gorm:"column:module_id;not null" json:"module_id"`
	Name        string         `gorm:"column:name;not null" json:"name"`
	Price       float64        `gorm:"column:price;not null" json:"price"`
	GrupPejuang bool           `gorm:"column:grup_pejuang;not null" json:"grup_pejuang"`
	Exercise    int64          `gorm:"column:exercise;not null" json:"exercise"`
	Access      int64          `gorm:"column:access;not null" json:"access"`
	Module      bool           `gorm:"column:module;not null" json:"module"`
	TryOut      int64          `gorm:"column:try_out;not null" json:"try_out"`
	Zoom        bool           `gorm:"column:zoom;not null" json:"zoom"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

// TableName Plan's table name
func (*Plan) TableName() string {
	return TableNamePlan
}
