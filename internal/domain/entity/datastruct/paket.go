package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type Plans struct {
	ID        string         `gorm:"type:text;primaryKey"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"PlanName" form:"PlanName"`
	SubPlan   []SubPlans     `gorm:"foreignKey:PlanID"`
}
