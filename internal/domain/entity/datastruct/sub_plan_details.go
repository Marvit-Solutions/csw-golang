package datastruct

import (
	"time"

	"gorm.io/gorm"
)

type SubPlanDetails struct {
	ID          string         `gorm:"type:text;primaryKey"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	SubPlanID   string         `json:"SubPlanID" form:"SubPlanID"`
	GrupPejuang bool           `json:"GrupPejuang" form:"GrupPejuang"`
	Exercise    uint           `json:"Exercise" form:"Exercise"`
	Access      uint           `json:"Access" form:"Access"`
	Module      bool           `json:"Module" form:"Module"`
	TryOut      uint           `json:"TryOut" form:"TryOut"`
	Zoom        bool           `json:"Zoom" form:"Zoom"`
}
