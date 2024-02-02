package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type SubPlans struct {
	ID          string         `gorm:"type:text;primaryKey"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	ModuleID    string         `json:"ModuleID" form:"ModuleID"`
	Name        string         `json:"Name" form:"Name"`
	Price       float64        `json:"Price" form:"Price"`
	GrupPejuang bool           `json:"GrupPejuang" form:"GrupPejuang"`
	Exercise    uint           `json:"Exercise" form:"Exercise"`
	Access      uint           `json:"Access" form:"Access"`
	Module      bool           `json:"Module" form:"Module"`
	TryOut      uint           `json:"TryOut" form:"TryOut"`
	Zoom        bool           `json:"Zoom" form:"Zoom"`
}
