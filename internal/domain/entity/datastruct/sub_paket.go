package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type SubPlan struct {
	ID            string         `gorm:"type:text;primaryKey"`
	CreatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	PlanID        string         `json:"PlanID" form:"PlanID"`
	Name          string         `json:"Name" form:"Name"`
	SubPlanDetail SubPlanDetail  `gorm:"foreignKey:SubPlanID"`
	Price         float64        `json:"Price" form:"Price"`
}
