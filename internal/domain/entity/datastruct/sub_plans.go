package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type SubPlans struct {
	ID             string         `gorm:"type:text;primaryKey"`
	CreatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt      time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	ModuleID       string         `json:"ModuleID" form:"ModuleID"`
	Name           string         `json:"Name" form:"Name"`
	SubPlanDetails SubPlanDetails `gorm:"foreignKey:SubPlanID"`
	Price          float64        `json:"Price" form:"Price"`
}
